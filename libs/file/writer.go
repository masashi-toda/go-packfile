package file

import (
	"io"
	"log"
	"strings"
)

type Writer interface {
	Write(value []byte) Writer
	WriteString(value string) Writer
	WriteStrings(values []string, sep string) Writer
	WriteNewLine() Writer
}

type writer struct {
	internal io.Writer
}

func (w *writer) Write(value []byte) Writer {
	_, err := w.internal.Write(value)
	if err != nil {
		log.Panicf("failed to write data [%s]", err.Error())
	}
	return w
}

func (w *writer) WriteString(value string) Writer {
	_, err := w.internal.Write([]byte(value))
	if err != nil {
		log.Panicf("failed to write data [%s]", err.Error())
	}
	return w
}

func (w *writer) WriteStrings(values []string, sep string) Writer {
	_, err := w.internal.Write([]byte(strings.Join(values, sep)))
	if err != nil {
		log.Panicf("failed to write data [%s]", err.Error())
	}
	return w
}

func (w *writer) WriteNewLine() Writer {
	_, err := w.internal.Write([]byte("\n"))
	if err != nil {
		log.Panicf("failed to write data [%s]", err.Error())
	}
	return w
}
