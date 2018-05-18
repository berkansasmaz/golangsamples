package main

import "fmt"

func main() {
	var karabaş köpek
	karabaş =köpek{tür:"kangal",ses:"hav hav!"}
	//sescıkar(karabaş)
	karabaş.sescıkarma()
	fmt.Println(karabaş.sescıkarma())
}

type köpek struct {
	tür string
	ses string
}

func (k köpek) sescıkarma() string {
	return "havhav"
}

type varlık interface {
	sescıkarma()
}

func sescıkar(x varlık) {
	x.sescıkarma()
}

