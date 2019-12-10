package main

type Imain interface {
	Do(ekle string) bool
	It() string
	Make(code int, text string)
	VoidFunc()
}

type Imptest struct {
	name string
}

func (i *Imptest) It() string {
	return i.name
}

func main() {

	a := Imptest{}
	a.It()
}
