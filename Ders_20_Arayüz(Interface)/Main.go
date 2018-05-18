package main

import "fmt"

func main() {
	//Değişkenler (nesneler gibi düşünebilirsiniz)
	karabaş:=köpek{}
	karabaş.tür="Kangal"
	karabaş.yaş=3

	var ceptelefonu telefon
	var poyraz rüzgar 

	//Arayüz üzerinden metotlara erişim
	sescikar(karabaş)
	sescikar(ceptelefonu)
	sescikar(poyraz)
	yağmurgetir(poyraz)

	//Arayüz kullanmadan
	fmt.Println("==============")
	fmt.Println(karabaş.sescikarma())
	fmt.Println(poyraz.yağmurgetir())
}

//tür tanımları
type köpek struct {
	yaş int
	tür string
}
type telefon string
type rüzgar string

//Türlere ait metotlar

func (x köpek) sescikarma() string  {
	return  "havhavhav!"
}

func (x telefon) sescikarma() string {
	return "didit didit!"
}

func (x rüzgar) sescikarma() string  {
	return "fiuuvv !"
}

func (x rüzgar) yağmurgetir() string {
	return "Bazı insanlar yağmuru hiseder, diğerleri ıslanır."
}

//Arayüz tanımları
type varlik interface {
	sescikarma() string
}

type meterolojik interface {
	yağmurgetir() string
}

//Arayüz-metot aracılığı yapan fonksiyonlar
func sescikar(x varlik)  {
	fmt.Println(x.sescikarma())
}

func yağmurgetir(x meterolojik) {
	fmt.Println(x.yağmurgetir())
}






