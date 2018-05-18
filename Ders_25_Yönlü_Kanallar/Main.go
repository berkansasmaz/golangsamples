package main

import "fmt"

func main() {
	//İstenirse kanal kullanımında yön belirtilebilir. Kanalı tek yönlü veri aktarımı yapacak şekilde kullanmak için, kanalı kullanacak olan fonksiyona özel bir bildirimde bulunumalıdır. Aşağıda, yönlü kanal kullanımı için bildirim kısmı verilmiştir.
	ses:=make(chan string,1)
	mikrofon(ses,"Cihaz kontrol!")
	hoparlor(ses,3)

}

func hoparlor(S <-chan string, ses_seviyesi int) {
	fmt.Printf("Anons Ses Seviyesi: %d\n", ses_seviyesi)
	fmt.Printf("Anons Metni: %s\n", <-S)
}

func mikrofon(S chan<- string, mesaj string) {
	S<-mesaj
}