package main

// type Imain interface {
// 	Do(ekle string) bool
// 	It() string
// 	Make(code int, text string)
// 	VoidFunc()
// 	Kod() (action float64)
// }

type DeneStruct struct {
	name string
}

func (ds *DeneStruct) DenemeFunction(funcparametresi int, ikinci bool) (geridonus string) {
	return ds.name
}

func main() {
}
