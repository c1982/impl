package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"testing"
)

func TestParseInterface(t *testing.T) {

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "./main.go", nil, parser.AllErrors)

	if err != nil {
		t.Error(err)
	}

	//pksg, err := parser.ParseDir(fset, "./")
	//ast.Print(fset, f)

	//pksg[""].Files

	im := map[string]*ast.FuncType{}

	ast.Inspect(f, func(n ast.Node) bool {

		switch x := n.(type) {
		case *ast.TypeSpec:

			if x.Name.IsExported() {

				switch d := x.Type.(type) {
				case *ast.StructType:
					//Type'ı buradan yakala.

				case *ast.InterfaceType:

					//fmt.Println(x.Name)
					list := d.Methods

					for i := 0; i < len(list.List); i++ {

						k := list.List[i]

						for z := 0; z < len(k.Names); z++ {
							//m := k.Names[z]
							//fmt.Println(m.Obj.Name)
						}

						switch ftt := k.Type.(type) {
						case *ast.FuncType:

							im[k.Names[0].Name] = ftt

							if ftt.Params.NumFields() > 0 {
								//fmt.Println("Aldığı parametreler")

								for ll := 0; ll < len(ftt.Params.List); ll++ {
									//ppp := ftt.Params.List[ll]
									//fmt.Print(ppp.Type, ",")
								}
							}

							if ftt.Results.NumFields() > 0 {
								//fmt.Println("Geri dönüş tipi")

								for oo := 0; oo < len(ftt.Results.List); oo++ {
									//rrr := ftt.Results.List[oo]
									//fmt.Print(rrr.Type, ",")
								}
							}

						}

						//fmt.Println("")
					}
				}
			}
		}

		return true
	})

	for k, v := range im {

		fmt.Println(k)

		for i := 0; i < len(v.Params.List); i++ {
			fmt.Println(v.Params.List[i].Names, v.Params.List[i].Type)
		}

		if v.Results.NumFields() > 0 {
			for p := 0; p < len(v.Results.List); p++ {
				fmt.Println(v.Results.List[p].Names)
			}
		}
	}

	structName := "Imptest"
	structDeclare := ""

	//struct'ı oku
	//interface methodlarını ekle.
	//dosyaya yaz.

	fileData := []byte(structDeclare)
	ioutil.WriteFile(structName+".go", fileData, 666)

}
