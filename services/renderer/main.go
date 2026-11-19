package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	maxSourceBytes = 128 * 1024
	maxSVGBytes    = 2 * 1024 * 1024
	renderTimeout  = 3 * time.Second
)

type renderFunc func(context.Context, string) ([]byte, error)

type renderResult struct {
	svg []byte
	err error
}

type errorCode string

const (
	codeInvalidSource        errorCode = "invalid_source"
	codeSourceTooLarge       errorCode = "source_too_large"
	codeUnsupportedMediaType errorCode = "unsupported_media_type"
	codeInvalidDiagram       errorCode = "invalid_diagram"
	codeSVGTooLarge          errorCode = "svg_too_large"
	codeRenderFailed         errorCode = "render_failed"
	codeRendererBusy         errorCode = "renderer_busy"
	codeRenderTimeout        errorCode = "render_timeout"
)

type apiError struct {
	Code       errorCode         `json:"code"`
	Message    string            `json:"message"`
	Diagnostic *sourceDiagnostic `json:"diagnostic,omitempty"`
}

func main() {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	server := &http.Server{
		Addr:              addr,
		Handler:           newHandler(renderD2, renderTimeout),
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       8 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    16 * 1024,
	}
	log.Printf("D2 renderer listening on %s", addr)
	log.Fatal(server.ListenAndServe())
}

func newHandler(render renderFunc, timeout time.Duration) http.Handler {
	slot := make(chan struct{}, 1)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /render/d2", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil || mediaType != "text/plain" || !isUTF8(params["charset"]) {
			writeError(w, http.StatusUnsupportedMediaType, codeUnsupportedMediaType, "Send UTF-8 text/plain.")
			return
		}

		body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxSourceBytes))
		if err != nil {
			var tooLarge *http.MaxBytesError
			if errors.As(err, &tooLarge) {
				writeError(w, http.StatusRequestEntityTooLarge, codeSourceTooLarge, "The D2 source exceeds 128 KiB.")
				return
			}
			writeError(w, http.StatusBadRequest, codeInvalidSource, "The D2 source could not be read.")
			return
		}

		source := string(body)
		if !utf8.Valid(body) || strings.TrimSpace(source) == "" {
			writeError(w, http.StatusBadRequest, codeInvalidSource, "Send non-empty UTF-8 D2 source.")
			return
		}

		select {
		case slot <- struct{}{}:
		default:
			writeError(w, http.StatusServiceUnavailable, codeRendererBusy, "A render is already running.")
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		done := make(chan renderResult, 1)
		go func() {
			// Keep the slot until the engine exits, including after an HTTP timeout.
			var result renderResult
			defer func() {
				if recover() != nil {
					log.Print("D2 renderer panicked")
					result = renderResult{err: errRendererFailed}
				}
				<-slot
				done <- result
			}()
			result.svg, result.err = render(ctx, source)
		}()

		select {
		case <-ctx.Done():
			writeError(w, http.StatusGatewayTimeout, codeRenderTimeout, "The render deadline was exceeded.")
		case result := <-done:
			var invalid *invalidDiagram
			switch {
			case ctx.Err() != nil:
				writeError(w, http.StatusGatewayTimeout, codeRenderTimeout, "The render deadline was exceeded.")
			case errors.As(result.err, &invalid):
				writeErrorResponse(w, http.StatusUnprocessableEntity, apiError{
					Code: codeInvalidDiagram, Message: invalid.message, Diagnostic: invalid.diagnostic,
				})
			case result.err != nil:
				writeError(w, http.StatusInternalServerError, codeRenderFailed, "The D2 renderer failed.")
			case len(result.svg) > maxSVGBytes:
				writeError(w, http.StatusUnprocessableEntity, codeSVGTooLarge, "The rendered SVG exceeds 2 MiB.")
			default:
				w.Header().Set("Content-Type", "image/svg+xml")
				w.Header().Set("Content-Security-Policy", "sandbox; default-src 'none'; style-src 'unsafe-inline'; font-src data:; img-src data:")
				_, _ = w.Write(result.svg)
			}
		}
	})
	return mux
}

func isUTF8(charset string) bool {
	return charset == "" || strings.EqualFold(charset, "utf-8")
}

func writeError(w http.ResponseWriter, status int, code errorCode, message string) {
	writeErrorResponse(w, status, apiError{Code: code, Message: message})
}

func writeErrorResponse(w http.ResponseWriter, status int, body apiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
