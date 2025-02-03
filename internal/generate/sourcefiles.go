package generate

import (
	"os"
	"strings"
)

type sourceFiles struct {
	files map[string][]string
}

func newSourceFiles() sourceFiles {
	return sourceFiles{
		files: make(map[string][]string),
	}
}

func (sf *sourceFiles) line(filename string, lineNumber int) string {
	lines, haveFile := sf.files[filename]
	if !haveFile {
		content, err := os.ReadFile(filename)
		fatalOnError(err)
		lines = strings.Split(string(content), "\n")
		sf.files[filename] = lines
	}

	return lines[lineNumber]
}
