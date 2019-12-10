package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestParseInterface(t *testing.T) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "./main.go", nil, parser.AllErrors)
	if err != nil {
		t.Error(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			if t.Name.IsExported() {

				switch x := t.Type.(type) {
				case *ast.InterfaceType:
					fmt.Println("Interface:", t.Name.Name)
					printMedhods(x.Methods.List)
				case *ast.StructType:
					fmt.Println("Struct:", t.Name.Name)

					//onj := node.Scope.Lookup(t.Name.Name)

				}
			}
		}
		return true
	})
}

func printMedhods(list []*ast.Field) {
	for _, method := range list {
		switch z := method.Type.(type) {
		case *ast.FuncType:
			fmt.Println("\tMethod:", method.Names[0].Name)
			printFunctionNames(z.Params.List)
		}
	}
}

func printFunctionNames(list []*ast.Field) {

	for _, sig := range list {
		for _, namesig := range sig.Names {
			fmt.Println("\t\t", namesig.Name, sig.Type)
		}
	}

}

func Test2(t *testing.T) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "./main.go", nil, parser.AllErrors)

	if err != nil {
		t.Error(err)
	}

	for k, v := range f.Scope.Objects {
		fmt.Println(k, v)
	}

}
