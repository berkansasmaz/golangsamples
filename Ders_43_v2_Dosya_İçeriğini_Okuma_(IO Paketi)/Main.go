package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	//hata denemi yapılan hali

	baytkesit, hata:= ioutil.ReadFile("deneme.txt")
	if hata !=nil{ // deneme.txt açılmazsa, belirtilen mesajı ekrana yazdıracaktır.
		fmt.Println("Dosya açılamadı")
		return
	}
	fmt.Println(string(baytkesit))

}
