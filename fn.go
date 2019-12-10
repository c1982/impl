package main

var Flist []Fitem = []Fitem{}

type Fitem struct {
	Name    string
	Parent  string
	Params  map[string]interface{}
	Returns map[string]interface{}
}
