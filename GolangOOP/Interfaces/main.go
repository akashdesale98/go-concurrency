package main

import "fmt"

type Language interface {
	getDevelopedBy()
}

type JavaScript struct {
	DevelopedBy string
}

func (javaScript JavaScript) getDevelopedBy() {
	fmt.Println(javaScript.DevelopedBy)
}

type Python struct {
	DevelopedBy string
}

func (python Python) getDevelopedBy() {
	fmt.Println(python.DevelopedBy)
}

type AllLanguages struct {
	L1 JavaScript
	L2 Python
}

func (Language AllLanguages) getDevelopedBy() {
	Language.L1.getDevelopedBy()
	Language.L2.getDevelopedBy()
}

func main() {
	Lang := &AllLanguages{
		L1: JavaScript{
			DevelopedBy: "Brendan Eich",
		},
		L2: Python{
			DevelopedBy: "Guido van Rossum",
		},
	}

	var ILanguage Language

	ILanguage = Lang

	ILanguage.getDevelopedBy()
}
