module github.com/pekim/freetype-go

go 1.23.4

// Address deprecation of clang_getDiagnosticCategoryName.
replace github.com/go-clang/clang-v15 => github.com/pekim/clang-v15 v0.0.0-20240830114552-c0d27ccce9ec

require (
	github.com/go-clang/clang-v15 v0.0.0-20230222085438-ee3102fa0c71
	golang.org/x/sync v0.10.0
)
