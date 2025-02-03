package generate

import (
	"fmt"
	"sync"
)

type generator struct {
	goFile *file
}

func Generate() {
	g := generator{
		goFile: newFile("./api.go"),
	}
	defer g.goFile.finish()

	api := &api{
		variablesLock: new(sync.Mutex),
	}
	fmt.Println()
	api.parseTranslationUnits()
	api.enrich1()
	// api.enrich2()
	api.generate(g)
	// fmt.Println()
	// api.printStats()
	// fmt.Println()
}
