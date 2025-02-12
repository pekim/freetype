# freetype

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pekim/freetype)](https://pkg.go.dev/github.com/pekim/freetype)

This is a cgo-free library that implements the much of the [FreeType](https://freetype.org/) C API.
It is a relatively thin wrapper around the excellent [libfreetype](https://pkg.go.dev/modernc.org/libfreetype),
providing a slightly more Go friendly API.

A degree of familiarity with FreeType's API will be required.
The [Freetype Tutorial](https://freetype.org/freetype2/docs/tutorial/index.html) is a good starting point.

## API

The API attempts to balance Go conventions with FreeType conventions.
And while it violates both from time to time, hopefully it doesn't stray too far from either.

### functions

For the most part there is a one-to-one mapping between C API function and Go functions.

### structs

Non-private struct fields are exported, with a few exceptions.

- Where there are a pair of fields with a count and a pointer to an array, they will not be exported.
  Instead a method that returns a slice is exported.
- Fields that are a pointer to a zero-terminated string are not exported.
  Instead there will be an exported method with the same name, and it will return a Go string.

### types

Most C types defined by FreeType have a corresponding Go type.

### C macros

Most FreeType C macros do not need to be exposed.
One exception is the [Font Testing Macros](https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html).
They are exposed as functions, such as
[HasHorizontal](https://pkg.go.dev/pekim/freetype#Face.HasHorizontal).

## Examples

Simple examples can be found in the `example` directory.

## Completeness

Most types and functions in the [Core API](https://freetype.org/freetype2/docs/reference/index.html#core-api) are implemented.
That should suffice for many glyph rasterization needs.

## Alternatives

While this library has its benefits there are alternatives that should be considered.

- [modernc.org/libfreetype](https://pkg.go.dev/modernc.org/libfreetype) -
  This is the library that this library wraps.
  It is the result of using ccgo to compile the FreeType C source code to Go.
  The main challenge with using it is that its API is a little tricky to use.
  Pointers are represented with `uintptr`. An instance of [libc.TLS](https://pkg.go.dev/modernc.org/libc#TLS) must be used,
  and passed to all functions. It does not export many constants that are needed.
- [github.com/golang/freetype](https://pkg.go.dev/github.com/golang/freetype) -
  This is a Go port of the C FreeType library.
  It has some limitations, a number of open bugs, and has not been updated in 8 years.
  However it's a good library, and will likely suffice for many use cases.
- [FreeType](https://freetype.org/) -
  It's perfectly possible to use cgo to use the FreeType library itself.
  This requires the library to be installed, for building apps the headers available,
  and a knowledge of C and cgo.

## Development

### source files

Many source files are named after one of the sections in the FreeType
[API Reference](https://freetype.org/freetype2/docs/reference/index.html).
The types and functions in those files are maintained in the same order as they appear in their section's documentation.

### pre-commit hook

There are configuration files for linting and other checks.
To use a git pre-commit hook for the checks

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` if not already installed
  - https://golangci-lint.run/usage/install/#local-installation
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install a git pre-commit hook in this repo's workspace
  - `pre-commit install`
