package main

import "fmt"

func main() {
	//Kanal Senkronizasyonu yaparken close fonksiyonunda da faydalanabiliriz.
	kontrol:=make(chan string)
	go func() {
		println("gorutin' den gelen mesaj")//çıktıda kırmızı yazılmasının sebebi fmt nin olmamasıdır.
		close(kontrol)//Eğer bu olmasayı <-kontrol ifadesinde hata vericekti.
	}()
	fmt.Println("main' den gelen mesaj")
	<-kontrol//program bu haliyle bu satırdan mesaj bekleyecek.ancak mesaj yerine kanalın kapatılmasıda bu satırın sorunsuz olarak tamamlanmasını sağlıyor.
}
