# freetype

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pekim/freetype-go)](https://pkg.go.dev/github.com/pekim/freetype-go)

This is a Go library that implements bindings for the [FreeType](https://freetype.org/) library.

## Requirements

### Build

#### C compiler

`cgo` is used. So a C compiler, such as clang or gcc, is required.

#### freetype headers

The FreeType headers must be available, so the freetype development package should be installed.
For example `dnf install freetype-devel` or `apt install libfreetype6-dev`.

### Runtime

The FreeType library must be available, as it will be dynamically loaded at runtime.

## Examples

Simple examples can be found in the `example` directory.

## Completeness

Most types and functions in the [Core API](https://freetype.org/freetype2/docs/reference/index.html#core-api) are implemented.
That should suffice for many glyph rasterization needs.

## Development

### source files

Most source files are named after one of the sections in the FreeType
[API Reference](https://freetype.org/freetype2/docs/reference/index.html).
The types and functions in those files are maintained in the same order as they appear in their section's documentation.

### type aliases

When defining Go types that correspond to C types, type aliases are used. This is done to reduce the need to convert the Go type to the C type when passing an argument to a C function.

### pre-commit hook

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` if not already installed
  - https://golangci-lint.run/usage/install/#local-installation
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
