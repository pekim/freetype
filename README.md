# freetype

## development

### source files

Most source files are named after one of the sections in the FreeType
[API Reference](https://freetype.org/freetype2/docs/reference/index.html).
The types and functions in those files are maintained in the same order as they appear in the section's documentation.

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
