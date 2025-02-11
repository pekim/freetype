# freetype

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pekim/freetype-go)](https://pkg.go.dev/github.com/pekim/freetype-go)

This is a cgo-free library that implements the [FreeType](https://freetype.org/) library.
It is a relatively thin wrapper around [libfreetype](https://pkg.go.dev/modernc.org/libfreetype),
providing a slightly more Go friendly API.

## API

### functions

For the most part there is a one-to-one mapping between C API function and Go functions.

### structs

Non-private struct fields are exported, with a few exceptions.

- Where there are a pair of fields with a count and a pointer to an array, they will not be exported.
  Instead a method that returns a slice is exported.
- Fields that are a pointer to a zero-terminated string are not exported.
  An exported method with the same name returns a Go string.

### types

Most C types defined by FreeType have a corresponding Go type.

### C macros

Most FreeType C macros do not need to be exposed.
One exception is the [Font Testing Macros](https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html).
They are exposed as functions, such as
[HasHorizontal](https://pkg.go.dev/pekim/freetype-go#Face.HasHorizontal).

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

### pre-commit hook

There are configuration files for linting and other checks.
To use a pre-commit hook for the checks

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` if not already installed
  - https://golangci-lint.run/usage/install/#local-installation
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install a git pre-commit hook in this repo's workspace
  - `pre-commit install`
