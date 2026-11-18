package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRenderLargeDiagram(t *testing.T) {
	source, err := os.ReadFile("testdata/architecture-1024.d2")
	if err != nil {
		t.Fatal(err)
	}
	response := postDiagram(newHandler(renderD2, renderTimeout), string(source), "text/plain; charset=utf-8")
	if response.Code != http.StatusOK || response.Header().Get("Content-Type") != "image/svg+xml" {
		t.Fatalf("status = %d, content type = %q", response.Code, response.Header().Get("Content-Type"))
	}

	var root struct {
		XMLName xml.Name
		ViewBox string `xml:"viewBox,attr"`
	}
	if err := xml.Unmarshal(response.Body.Bytes(), &root); err != nil {
		t.Fatalf("invalid SVG: %v", err)
	}
	if root.XMLName.Local != "svg" || root.ViewBox == "" {
		t.Fatalf("unexpected SVG root: %+v", root)
	}
	for _, label := range []string{"Ingest 001", "UI 128"} {
		if !strings.Contains(response.Body.String(), label) {
			t.Errorf("missing label %q", label)
		}
	}
}

func TestRenderRequestErrors(t *testing.T) {
	handler := newHandler(renderD2, renderTimeout)
	cases := []struct {
		name, source, contentType, code string
		status                          int
	}{
		{"unsupported media type", "a -> b", "application/json", "unsupported_media_type", 415},
		{"unsupported charset", "a -> b", "text/plain; charset=utf-16", "unsupported_media_type", 415},
		{"empty input", "  \n", "text/plain", "invalid_source", 400},
		{"invalid UTF-8", "\xff", "text/plain", "invalid_source", 400},
		{"oversized source", strings.Repeat("a", maxSourceBytes+1), "text/plain", "source_too_large", 413},
		{"invalid diagram", "a: {", "text/plain", "invalid_diagram", 422},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertError(t, postDiagram(handler, tc.source, tc.contentType), tc.status, tc.code)
		})
	}
}

func TestRendererDoesNotOpenHostImports(t *testing.T) {
	input := filepath.Join(t.TempDir(), "included.d2")
	if err := os.WriteFile(input, []byte("label: local test fixture\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	handler := newHandler(renderD2, renderTimeout)
	response := postDiagram(handler, "...@"+input, "text/plain")
	assertError(t, response, http.StatusUnprocessableEntity, "invalid_diagram")
	if strings.Contains(response.Body.String(), input) {
		t.Fatal("the response exposed a host path")
	}
}

func TestRenderTimeoutKeepsSlotOccupiedUntilRenderStops(t *testing.T) {
	started := make(chan struct{})
	release := make(chan struct{})
	defer close(release)

	handler := newHandler(func(context.Context, string) ([]byte, error) {
		close(started)
		<-release
		return nil, nil
	}, 40*time.Millisecond)

	response := make(chan *httptest.ResponseRecorder, 1)
	go func() {
		response <- postDiagram(handler, "a -> b", "text/plain")
	}()
	<-started
	assertError(t, <-response, http.StatusGatewayTimeout, "render_timeout")
	assertError(t, postDiagram(handler, "c -> d", "text/plain"), http.StatusServiceUnavailable, "renderer_busy")
}

func TestRenderedImageLimit(t *testing.T) {
	handler := newHandler(func(context.Context, string) ([]byte, error) {
		return make([]byte, maxSVGBytes+1), nil
	}, renderTimeout)
	assertError(t, postDiagram(handler, "a", "text/plain"), http.StatusUnprocessableEntity, "svg_too_large")
}

func postDiagram(handler http.Handler, body, contentType string) *httptest.ResponseRecorder {
	request := httptest.NewRequest(http.MethodPost, "/render/d2", strings.NewReader(body))
	request.Header.Set("Content-Type", contentType)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	return response
}

func assertError(t *testing.T, response *httptest.ResponseRecorder, wantStatus int, wantCode string) {
	t.Helper()
	if response.Code != wantStatus || response.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected HTTP %d with JSON, got %d %s", wantStatus, response.Code, response.Header().Get("Content-Type"))
	}

	var body apiError
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.Code != wantCode || body.Message == "" {
		t.Fatalf("expected %s with a message, got %+v", wantCode, body)
	}
}
