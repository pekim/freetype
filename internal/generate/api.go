package generate

import (
	"fmt"
	"sync"
	"time"
)

// _ "embed"
// "muzzammil.xyz/jsonc"

// //go:embed api.jsonc
// var apiJson []byte

type api struct {
	// Records       []record   `json:"records"`
	// Enums         []enum     `json:"enums"`
	// Functions     []callable `json:"functions"`
	// Typedefs      []typedef  `json:"typedefs"`
	// Variables     []variable
	tu            translationUnit
	variablesLock *sync.Mutex
}

func (api *api) parseTranslationUnits() {
	// err := json.Unmarshal(jsonc.ToJSON(apiJson), &api)
	// fatalOnError(err)

	fmt.Print("parse header file ")
	start := time.Now()
	api.tu = newTranslationUnit("internal/generate/freetype.h")
	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
}

func (api *api) enrich1() {
	fmt.Print("enrich 1")
	start := time.Now()
	api.tu.enrichApi(api)
	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
}

// func (api api) enrich2() {
// 	fmt.Print("enrich 2")
// 	start := time.Now()
// 	for i := range api.Enums {
// 		enum := &api.Enums[i]
// 		enum.enrich2(api)
// 	}
// 	for i := range api.Typedefs {
// 		record := &api.Typedefs[i]
// 		record.enrich2(api)
// 	}
// 	for i := range api.Functions {
// 		function := &api.Functions[i]
// 		function.enrich2(nil, api)
// 	}
// 	for i := range api.Records {
// 		record := &api.Records[i]
// 		record.enrich2(api)
// 	}
// 	for i := range api.Variables {
// 		variable := &api.Variables[i]
// 		variable.enrich2(api)
// 	}
// 	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
// }

// func (api api) findRecord(name string) (*record, bool) {
// 	for i, record := range api.Records {
// 		if record.CppName == name {
// 			return &api.Records[i], true
// 		}
// 	}
// 	return nil, false
// }

// func (api api) findEnum(name string) (*enum, bool) {
// 	for i, enum := range api.Enums {
// 		if enum.CppName == name {
// 			return &api.Enums[i], true
// 		}
// 	}
// 	return nil, false
// }

// func (api api) findFunction(name string) (*callable, bool) {
// 	for i, function := range api.Functions {
// 		if function.CppName == name {
// 			return &api.Functions[i], true
// 		}
// 	}
// 	return nil, false
// }

// func (api api) findTypedef(name string) (*typedef, bool) {
// 	for i, typedef := range api.Typedefs {
// 		if typedef.CppName == name {
// 			return &api.Typedefs[i], true
// 		}
// 	}
// 	return nil, false
// }

// func (api api) generate(g generator) {
// 	fmt.Print("generate")
// 	start := time.Now()

// 	for _, record := range api.Records {
// 		record.generateCStruct(g)
// 		for _, record := range record.Records {
// 			record.generateCStruct(g)
// 		}
// 	}

// 	for _, record := range api.Records {
// 		record.generate(g)
// 	}

// 	for _, enum := range api.Enums {
// 		enum.generate(g)
// 	}

// 	for _, function := range api.Functions {
// 		function.generate(g)
// 	}

// 	for _, typedef := range api.Typedefs {
// 		typedef.generate(g)
// 	}

// 	g.headerFile.writeln()
// 	slices.SortFunc(api.Variables, func(a, b variable) int {
// 		return strings.Compare(a.cppName, b.cppName)
// 	})
// 	for _, variable := range api.Variables {
// 		variable.generate(g)
// 	}

// 	fmt.Printf(" %dms\n", time.Since(start).Milliseconds())
// }

// func (api api) printStats() {
// 	classCount := 0
// 	structCount := 0
// 	classEnumCount := 0
// 	classMethodCount := 0
// 	classRecordCount := 0
// 	structEnumCount := 0
// 	structMethodCount := 0
// 	structRecordCount := 0
// 	for _, record := range api.Records {
// 		if record.isClass {
// 			classCount++
// 			classEnumCount += len(record.Enums)
// 			classMethodCount += len(record.Methods)
// 			classRecordCount += len(record.Records)
// 		} else {
// 			structCount++
// 			structEnumCount += len(record.Enums)
// 			structMethodCount += len(record.Methods)
// 			structRecordCount += len(record.Records)
// 		}
// 	}

// 	fmt.Printf("%4d classes (with %d enums, %d methods, %d records)\n",
// 		classCount, classEnumCount, classMethodCount, classRecordCount)
// 	fmt.Printf("%4d enums\n", len(api.Enums))
// 	fmt.Printf("%4d functions\n", len(api.Functions))
// 	fmt.Printf("%4d structs (with %d enums, %d methods, %d records)\n",
// 		structCount, structEnumCount, structMethodCount, structRecordCount)
// 	fmt.Printf("%4d typedefs\n", len(api.Typedefs))
// }

// func (api api) goNameForCppName(cppName string) (string, bool) {
// 	if enum, ok := api.findEnum(cppName); ok {
// 		return enum.goName, true
// 	}
// 	if record, ok := api.findRecord(cppName); ok {
// 		return record.goName, true
// 	}
// 	if typedef, ok := api.findTypedef(cppName); ok {
// 		return typedef.goName, true
// 	}

// 	return "", false
// }
