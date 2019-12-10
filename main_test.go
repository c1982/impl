package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFindStruct(t *testing.T) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "./main.go", nil, parser.AllErrors)
	if err != nil {
		t.Error(err)
	}

	for _, d := range node.Decls {

		switch t := d.(type) {
		case *ast.GenDecl:
			for _, spec := range t.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					fmt.Println("Struct:", s.Name.Name)
				}
			}
		}
	}
}

func TestFindFunctionAndSignature(t *testing.T) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "./main.go", nil, parser.AllErrors)
	if err != nil {
		t.Error(err)
	}

	//ast.Print(fset, node)

	fi := Fitem{
		Params:  make(map[string]interface{}),
		Returns: make(map[string]interface{}),
	}

	for _, d := range node.Decls {

		switch t := d.(type) {
		case *ast.FuncDecl:

			if t.Name.Name == "main" {
				continue
			}

			fi.Name = t.Name.Name

			fmt.Println("Func:", t.Name.Name)

			if t.Recv != nil {
				for _, r := range t.Recv.List {
					if st, ok := r.Type.(*ast.StarExpr); ok {

						switch stx := st.X.(type) {
						case *ast.Ident:
							fmt.Println("\tStruct:", stx.Name)
							fi.Parent = stx.Name
						}
					}
				}
			}

			for _, p := range t.Type.Params.List {
				for _, pname := range p.Names {
					fmt.Println("\tparameter:", pname.Name, "type:", p.Type)
					fi.Params[pname.Name] = p.Type
				}
			}

			if t.Type.Results != nil {
				for _, sult := range t.Type.Results.List {
					for _, sname := range sult.Names {
						fmt.Println("\treturn:", sname.Name, "Type:", sult.Type)
						fi.Returns[sname.Name] = sult.Type
					}
				}
			}

		}
	}

	if fi.Name != "DenemeFunction" {
		t.Error("Value:", fi.Name, "Expected: DenemeFunction")
	}

	if fi.Parent != "DeneStruct" {
		t.Error("Value:", fi.Parent, "Expected: DeneStruct")
	}

	vl, ok := fi.Params["funcparametresi"]

	if !ok {
		t.Error("key not found: funcparametresi")
	}

}

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
				// switch x := t.Type.(type) {
				// case *ast.InterfaceType:
				// 	fmt.Println("Interface:", t.Name.Name)l
				// 	printMedhods(x.Methods.List)
				// case *ast.StructType:
				// 	fmt.Println("Struct:", t.Name.Name)
				// }
			}
		case *ast.FuncDecl:

			fmt.Println("Func:", t.Name)

			if t.Recv != nil {

				for _, rsv := range t.Recv.List {
					if len(rsv.Names) > 0 {
						fmt.Println("\tRecv:", rsv.Names[0], "Type", rsv.Type)
					} else {
						fmt.Println("\tRecv:", rsv.Type)
					}
				}

			}

			for _, sig := range t.Type.Params.List {
				fmt.Println("\tName:", sig.Names[0], "Type:", sig.Type)
			}
		}

		return true
	})
}

func TestInterface2(t *testing.T) {

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

					for _, meths := range x.Methods.List {

						for _, name := range meths.Names {

							fmt.Println("\tMethod Name:", name.String())

							switch ft := meths.Type.(type) {
							case *ast.FuncType:

								for _, p := range ft.Params.List {
									fmt.Printf("\t\tParams Type: %v\n", p.Type)
								}

								if ft.Results == nil {
									continue
								}

								for _, r := range ft.Results.List {
									switch t := r.Type.(type) {
									case *ast.StarExpr:
										if ident, ok := t.X.(*ast.Ident); ok {
											fmt.Println("\t\tReturn Name:", ident.Name)
										}
									case *ast.Ident:
										fmt.Println("\t\tReturn Type:", t.Name)
									}
								}

							default:
								fmt.Println("\t\tReturn df Type:", ft)
							}
						}
					}
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
