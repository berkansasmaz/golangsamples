package main

import (
	"time"
	"fmt"
)

func main() {
//Timer kullanımı, aslında beklemekten daha farklı özelliklere de sahiptir. Örneğin bir timer yeniden başlatılabilir(reset) veya durdurulabilir(stop).Aşağıda sayacın tamamlanmadan durdurulmasına örnek bir uygulama verilmiştir.
	sayac:=time.NewTimer(time.Second*3)
	fmt.Println("3 saniye geri sayım başlatıldı...")
	go func() {
		<-sayac.C
		fmt.Println("sayaç tamamlandı")
	}()

	//istenildiğinde timer durdurulabilir:
	time.Sleep(time.Second*1)
	iptal:=sayac.Stop()
	fmt.Println("Sayaç durduruldu mu?",iptal)
	//Bu programda biraz değişiklik yapalım ve time.sleep satırını time.sleep(time.second*4) şeklinde düzenleyelim. Bu sefer sayaç tamamlanacak kadar süre bulmakta ve ekran çıktısı değişmektedir. Sayaç kendi kendine tamamlandığı için, sayaç.stop() işlevi false değerini döndürmektedir.
}
