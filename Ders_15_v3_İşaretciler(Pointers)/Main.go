package main

import "fmt"

func main() {
	//Fonsiyona işaretçi gönderildiğinde; değişkenin kopyası değil, adresi gönderilir. Fonksiyoların bu işaretçiyi kullanarak yapılan değişiklik, orjinal veri üzerinde yapılır.
	metin:="Türkiye"
	fmt.Println(metin)// metin = Türkiye

	fonk1(&metin)//metin' in adresini fonksiyona bildirdik yada new keyword de kullanılabilir
	fmt.Println(metin)//metin="Türkiye Cumhuriyeti
}

func fonk1(ülke *string)  {
	//ülke = "Türkiye
	*ülke = *ülke + " Cumhuriyeti"

}