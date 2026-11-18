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

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
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
			writeError(w, http.StatusUnsupportedMediaType, "unsupported_media_type", "Send UTF-8 text/plain.")
			return
		}

		body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxSourceBytes))
		if err != nil {
			var tooLarge *http.MaxBytesError
			if errors.As(err, &tooLarge) {
				writeError(w, http.StatusRequestEntityTooLarge, "source_too_large", "The D2 source exceeds 128 KiB.")
				return
			}
			writeError(w, http.StatusBadRequest, "invalid_source", "The D2 source could not be read.")
			return
		}

		source := string(body)
		if !utf8.Valid(body) || strings.TrimSpace(source) == "" {
			writeError(w, http.StatusBadRequest, "invalid_source", "Send non-empty UTF-8 D2 source.")
			return
		}

		select {
		case slot <- struct{}{}:
		default:
			writeError(w, http.StatusServiceUnavailable, "renderer_busy", "A render is already running.")
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
			writeError(w, http.StatusGatewayTimeout, "render_timeout", "The render deadline was exceeded.")
		case result := <-done:
			switch {
			case ctx.Err() != nil:
				writeError(w, http.StatusGatewayTimeout, "render_timeout", "The render deadline was exceeded.")
			case errors.Is(result.err, errInvalidDiagram):
				writeError(w, http.StatusUnprocessableEntity, "invalid_diagram", "The D2 source could not be compiled. File imports are disabled.")
			case result.err != nil:
				writeError(w, http.StatusInternalServerError, "render_failed", "The D2 renderer failed.")
			case len(result.svg) > maxSVGBytes:
				writeError(w, http.StatusUnprocessableEntity, "svg_too_large", "The rendered SVG exceeds 2 MiB.")
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

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(apiError{Code: code, Message: message})
}
