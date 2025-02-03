package generate

import (
	"fmt"
	"go/format"
	"os"
	"strings"
)

type file struct {
	filename string
	content  strings.Builder
}

func newFile(filename string) *file {
	f := &file{
		filename: filename,
	}

	f.writelnf(`
		// This is a generated file. DO NOT EDIT.

		package freetype

		// #cgo pkg-config: freetype2
		//
		// #include <ft2build.h>
		// #include FT_FREETYPE_H
		import "C"
	`)

	return f
}

func (f *file) close() {
	unformatted := []byte(f.content.String())
	var formatted []byte

	formatted, err := format.Source(unformatted)
	if err != nil {
		fmt.Printf("failed to format %s: %s\n", f.filename, err.Error())
		formatted = unformatted
	}

	err = os.WriteFile(f.filename, formatted, 0644)
	fatalOnError(err)
}

func (f *file) write(content string) {
	f.content.WriteString(content)
}

func (f *file) writeln(args ...any) {
	line := fmt.Sprint(args...)
	f.write(line)
	f.write("\n")
}

func (f *file) writef(format string, args ...any) {
	f.write(fmt.Sprintf(format, args...))
}

func (f *file) writelnf(format string, args ...any) {
	f.writef(format, args...)
	f.write("\n")
}

func (f *file) finish() {
	f.writeln(`
	`)

	f.close()
}
