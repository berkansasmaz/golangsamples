package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	//Önceki uygulamada, os paketi içinde tanımlanmış olan File veri türüne ait olan Read() metodunu kullanmıştık. Bu sefer io\ioutil (io paketinin alt paketi olan ioutil paketi) paketi ile işlem yapıcaz. Bu pakatte gelen fonksiyonlar, dosyadan veri okuma işlemini oladukça kolaylaştırmaktadır. Önceki uygulamada  dosyadan veri okumak için; önce dosyayı açmış, dosya boyutunu almış, buna göre baytlar halinde okumuştuk. Bu sefer, tüm dosya içeriğini tek seferde bir değişkene aktarabiliyoruz.

	//dosya içeriğini, baytkesit değişkenine aktar:
	baytkesit, _:= ioutil.ReadFile("deneme.txt")

	//ASCII bayt kodlarını metne dönüştürüp ekrana yaz:
	fmt.Println(string(baytkesit))


}
