package main

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"log/slog"
	"strings"
	"sync/atomic"

	"github.com/d2lang/d2/d2graph"
	"github.com/d2lang/d2/d2layouts/d2dagrelayout"
	"github.com/d2lang/d2/d2lib"
	"github.com/d2lang/d2/d2parser"
	"github.com/d2lang/d2/d2renderers/d2svg"
	d2log "github.com/d2lang/d2/lib/log"
	"github.com/d2lang/d2/lib/textmeasure"
)

var (
	errRendererFailed = errors.New("renderer failed")
	diagramLogger     = slog.New(slog.NewTextHandler(io.Discard, nil))
)

const diagramPath = "diagram.d2"

type sourceDiagnostic struct {
	Line    int    `json:"line"`
	Column  int    `json:"column"`
	Message string `json:"message"`
}

type invalidDiagram struct {
	message    string
	diagnostic *sourceDiagnostic
}

func (e *invalidDiagram) Error() string { return e.message }

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
	imports := &noImports{}
	diagram, _, err := d2lib.Compile(d2log.With(ctx, diagramLogger), source, &d2lib.CompileOptions{
		// D2 0.9.0 falls back to the host filesystem when FS is nil.
		FS:        imports,
		InputPath: diagramPath,
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
		return nil, diagramCompileError(err, imports.blocked.Load())
	}

	return d2svg.Render(diagram, renderOpts)
}

func diagramCompileError(err error, importsBlocked bool) *invalidDiagram {
	if importsBlocked {
		return &invalidDiagram{message: "File imports are disabled in D2 previews."}
	}
	result := &invalidDiagram{message: "The D2 source could not be compiled."}
	var compileError *d2parser.ParseError
	if errors.As(err, &compileError) && len(compileError.Errors) > 0 {
		first := compileError.Errors[0]
		if first.Range.Path == diagramPath {
			result.diagnostic = &sourceDiagnostic{
				Line:    first.Range.Start.Line + 1,
				Column:  first.Range.Start.Column + 1,
				Message: strings.TrimPrefix(first.Message, first.Range.String()+": "),
			}
		}
	}
	return result
}

type noImports struct {
	blocked atomic.Bool
}

func (imports *noImports) Open(string) (fs.File, error) {
	imports.blocked.Store(true)
	return nil, fs.ErrPermission
}
