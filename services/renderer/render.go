package main

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"log/slog"

	"github.com/d2lang/d2/d2graph"
	"github.com/d2lang/d2/d2layouts/d2dagrelayout"
	"github.com/d2lang/d2/d2lib"
	"github.com/d2lang/d2/d2renderers/d2svg"
	d2log "github.com/d2lang/d2/lib/log"
	"github.com/d2lang/d2/lib/textmeasure"
)

var (
	errInvalidDiagram = errors.New("invalid diagram")
	errRendererFailed = errors.New("renderer failed")
	diagramLogger     = slog.New(slog.NewTextHandler(io.Discard, nil))
)

func renderD2(ctx context.Context, source string) ([]byte, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	ruler, err := textmeasure.NewRuler()
	if err != nil {
		return nil, err
	}

	pad := int64(36)
	layout := "dagre"
	renderOpts := &d2svg.RenderOpts{Pad: &pad}
	diagram, _, err := d2lib.Compile(d2log.With(ctx, diagramLogger), source, &d2lib.CompileOptions{
		// D2 0.9.0 falls back to the host filesystem when FS is nil.
		FS:        noImports{},
		InputPath: "diagram.d2",
		Ruler:     ruler,
		Layout:    &layout,
		LayoutResolver: func(string) (d2graph.LayoutGraph, error) {
			return d2dagrelayout.DefaultLayout, nil
		},
	}, renderOpts)
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if err != nil {
		return nil, errInvalidDiagram
	}

	return d2svg.Render(diagram, renderOpts)
}

type noImports struct{}

func (noImports) Open(string) (fs.File, error) {
	return nil, fs.ErrPermission
}
