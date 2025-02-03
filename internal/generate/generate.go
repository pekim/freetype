package generate

import (
	"fmt"
	"sync"
)

var headerFiles = []string{
	"internal/generate/freetype.h",
}

type generator struct {
	// goFile     *fileGo
	// headerFile *fileHeader
	// cFile    *fileCpp
}

func Generate() {
	g := generator{}
	fmt.Println(g)

	// g.goFile = newFileGo()
	// defer g.goFile.finish()

	// g.headerFile = newFileHeader()
	// defer g.headerFile.finish()

	// g.cppFile = newFileCpp()
	// defer g.cppFile.finish()

	api := &api{
		variablesLock: new(sync.Mutex),
	}
	fmt.Println()
	api.parseTranslationUnits()
	api.enrich1()
	// api.enrich2()
	// api.generate(g)
	// fmt.Println()
	// api.printStats()
	// fmt.Println()
}
