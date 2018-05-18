package main

import "fmt"

func main() {
	// Go dilinde işaretçi oluşturmak için, adres:=&sayı yazılacağı gibi, dahili olarak gelen new fonksiyonu da kullanılabilir.
	adres:=new(int)
	*adres=2016
	fmt.Println("Adres: ", adres)
	fmt.Println("Adresteki veri", *adres)
}
