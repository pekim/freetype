package generate

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/go-clang/clang-v15/clang"
)

type translationUnit struct {
	headerFile string
	clang.TranslationUnit
	sourceFiles sourceFiles
}

var clangResourceDir = sync.OnceValue[string](func() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	fatalOnError(err)

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		fatal("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		fatalf("expected clang resource dir to start with '/', but it %s", resDir)
	}

	return resDir
})

func newTranslationUnit(headerFile string) translationUnit {
	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-I", "./internal/include/",
		"-x", "c-header",
	}

	transUnit := translationUnit{
		headerFile:  headerFile,
		sourceFiles: newSourceFiles(),
	}
	index := clang.NewIndex(0, 1)
	errCode := index.ParseTranslationUnit2(headerFile, parseArgs, nil,
		clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord,
		&transUnit.TranslationUnit,
	)
	if errCode != clang.Error_Success {
		fatal(errCode)
	}
	return transUnit
}

func (tu translationUnit) enrichApi(_api *api) {
	tu.TranslationUnitCursor().Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		file, line, col, _ := cursor.Location().FileLocation()
		if !strings.Contains(file.Name(), "freetype/") {
			// Skip declarations from non-freetype files.
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {
		case clang.Cursor_MacroDefinition:
			if strings.HasPrefix(cursor.Spelling(), "FT_") {
				fmt.Println(file.Name(), line, col, cursor.Kind(), cursor.Spelling())
				line := tu.sourceFiles.line(file.Name(), int(line))
				fmt.Println("  ", line)
			}

			// case clang.Cursor_MacroExpansion:
			// 	fmt.Println("!!!!", cursor.Spelling())
			// fmt.Println("  ", cursor.IsMacroFunctionLike())
			// cursor.Visit(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
			// 	fmt.Println("  ", cursor.Kind(), cursor.Spelling())
			// 	return clang.ChildVisit_Continue
			// })
		}

		// 	switch cursor.Kind() {
		// 	case clang.Cursor_Namespace:
		// 		namespace := cursor.Spelling()
		// 		cursor.Visit(func(cursor, parent clang.Cursor) (status clang.ChildVisitResult) {
		// 			qualifiedName := fmt.Sprintf("%s::%s", namespace, cursor.Spelling())

		// 			switch cursor.Kind() {
		// 			case clang.Cursor_FunctionDecl:
		// 				if function, ok := api.findFunction(qualifiedName); ok {
		// 					function.enrich1(nil, cursor)
		// 				}

		// 			case clang.Cursor_EnumDecl:
		// 				if !cursorHasChildren(cursor) {
		// 					// Skip forward declarations.
		// 					break
		// 				}
		// 				if enum, ok := api.findEnum(qualifiedName); ok {
		// 					enum.enrich1(nil, cursor)
		// 				}
		// 			}

		// 			return clang.ChildVisit_Continue
		// 		})

		// 	case clang.Cursor_ClassDecl:
		// 		if !cursorHasChildren(cursor) {
		// 			// Skip forward declarations.
		// 			break
		// 		}
		// 		if record, ok := api.findRecord(cursor.Spelling()); ok {
		// 			record.enrich1(cursor, nil)
		// 		}

		// 	case clang.Cursor_ClassTemplate:
		// 		if record, ok := api.findRecord(cursor.Spelling()); ok {
		// 			record.enrich1(cursor, nil)
		// 		}

		// 	case clang.Cursor_EnumDecl:
		// 		if !cursorHasChildren(cursor) {
		// 			// Skip forward declarations.
		// 			break
		// 		}

		// 		if enum, ok := api.findEnum(cursor.Spelling()); ok {
		// 			enum.enrich1(nil, cursor)
		// 		}

		// 	case clang.Cursor_StructDecl:
		// 		if !cursorHasChildren(cursor) {
		// 			// Skip forward declarations.
		// 			break
		// 		}
		// 		if record, ok := api.findRecord(cursor.Spelling()); ok {
		// 			record.enrich1(cursor, nil)
		// 		}

		// 	case clang.Cursor_FunctionDecl:
		// 		if function, ok := api.findFunction(cursor.Spelling()); ok {
		// 			function.enrich1(nil, cursor)
		// 		}

		// 	case clang.Cursor_TypedefDecl:
		// 		if typedef, ok := api.findTypedef(cursor.Spelling()); ok {
		// 			typedef.enrich1(cursor)
		// 		}

		// 	case clang.Cursor_VarDecl:
		// 		api.variablesLock.Lock() // multiple translation units may append in different Go routines
		// 		api.Variables = append(api.Variables, newVariable(cursor))
		// 		api.variablesLock.Unlock()
		// 	}

		return clang.ChildVisit_Continue
	})
}

// func cursorHasChildren(cursor clang.Cursor) bool {
// 	childCount := 0
// 	cursor.Visit(func(_, _ clang.Cursor) (status clang.ChildVisitResult) {
// 		childCount++
// 		return clang.ChildVisit_Break
// 	})

// 	return childCount > 0
// }

// type speller interface {
// 	Spelling() string
// }

// const dumpIndent = "  "

// func init() {
// 	// Prevent 'unused' warning if dumpCursor is not called.
// 	// Avoids having to comment out the function when it's not used.
// 	dumpCursor(clang.Cursor{}, "dummy")
// }

// func dumpCursor(cursor clang.Cursor, prefix string) {
// 	if prefix == "dummy" {
// 		return
// 	}

// 	fmt.Printf("%s%s : %s : %s\n", prefix,
// 		cursor.Kind().Spelling(),
// 		cursor.Spelling(),
// 		cursor.Type().Kind())

// 	dumpSpelling(prefix, "Type", cursor.Type())
// 	if cursor.AccessSpecifier() != clang.AccessSpecifier_Invalid {
// 		dumpSpelling(prefix, "AccessSpecifier", cursor.AccessSpecifier())
// 	}
// 	dumpCursorValue(prefix, "DisplayName", cursor.DisplayName())
// 	if cursor.CXXRecord_IsAbstract() {
// 		fmt.Println(prefix, "Abstract")
// 	}

// 	cursor.Visit(func(cursor, _ clang.Cursor) (status clang.ChildVisitResult) {
// 		dumpCursor(cursor, prefix+dumpIndent)
// 		return clang.ChildVisit_Continue
// 	})
// }

// func dumpCursorValue(prefix string, name string, value string) {
// 	if len(value) > 0 {
// 		fmt.Printf("%s%s = %s\n", prefix+dumpIndent+dumpIndent, name, value)
// 	}
// }

// func dumpSpelling(prefix string, name string, speller speller) {
// 	value := speller.Spelling()
// 	dumpCursorValue(prefix, name, value)
// }
