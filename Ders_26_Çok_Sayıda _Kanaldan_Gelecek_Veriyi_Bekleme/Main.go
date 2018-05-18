package main

import (
	"time"
	"fmt"
)

func main() {
	//Çok sayıda gorutin çalıştırıp, hepsinin tamamlanmasını beklememiz gerektiğini düşünelim. Gorutinler sonlanmadan önce kanal üzerinden gönderecekleri veriyi de işlememiz gereksin. Kanallar ile birlikte kullanılan select ifadesi bu tür ihtiyaçları kolaylıkla karşılamaktadır. Select kullanımı, şekil ve mantık olarak switch-case yapısına benzer. Eş zamanlı olarak çalışan gorutinlerin kanallara göndereceği veriyi her kanal için bekler. Gelen veriyi bir koşula göre değerlendirip başka işler yapabilmemizi sağlar.

	//Bu uygulamada biri uzun diğeri kısa sürede tamamlanan iki farklı gorutin bulunmaktadır. Bu iki gorutin, sonlanmadan önce kendine ait kanala mesaj bırakmaktadır. Önce uzun_işlem başlatılıyor, hemen arkasından kisa_işlem başlatılıyor. Sonra da select-case yapısı sayesinde, kanallardan gelen mesajlar bekleniyor. Select bloğunun for döngüsü içinde olduğuna ve döngü sayısının, mesajı beklenen kanal sayına eşit olduğuna dikkat edelim.

	zaman_yaz()
	k1:=make(chan string)
	k2:=make(chan string)
	go uzun_işlem(k1)
	go kısa_işlem(k2)

	for i := 1; i <= 2; i++ {//döngü sayısı kanallardan gelicek mesaj sayına eşit olmalı
		select {
		case mesaj1:= <-k1://uzun mesaj önce case edildiği halde kısa mesaj daha önce yazdırıldı
			fmt.Println(mesaj1,"iş tamamlandı")
		case mesaj2 := <-k2:
			fmt.Println(mesaj2,"İş tamamlandı")
		}
	}
	zaman_yaz()
}

func uzun_işlem(k chan string) {
	time.Sleep(time.Second*5)
	k <- "uzun"
}

func kısa_işlem(k chan string) {
	time.Sleep(time.Second*2)
	k <- "kısa"
}

func zaman_yaz()  {
	fmt.Println(time.Now())
}