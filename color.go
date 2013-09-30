package logger

import (
	"github.com/wsxiaoys/terminal"
	"io"
)

type ColorWriter struct {
	Color  string
	Writer *terminal.TerminalWriter
}

var _ io.Writer = ColorWriter{}

func NewColorWriter(color string) ColorWriter {
	return ColorWriter{
		Color:  color,
		Writer: terminal.Stdout,
	}
}

func (w ColorWriter) Write(p []byte) (n int, err error) {
	return w.Writer.Color(w.Color).Write(p)
}
