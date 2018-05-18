package main

import "fmt"

func main() {
	var sayi int=54
	fmt.Println("Sayı: " ,sayi)
	fmt.Println("Sayını adresi: ", &sayi)// adrestki veriyi & simgesi temsil eder bellek adresini yazdırdık
	adres:=&sayi//adress isminde bir değişken oluşturduk
	fmt.Println("=================")
	fmt.Println("Adres: ",adres)
	fmt.Println("Adres' teki veri: ", *adres)//  *  işaretçinin gösterdiği adresteki veriyi temsil eder. yani adres içeriğini yazdırır

}
