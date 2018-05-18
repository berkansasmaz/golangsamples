package main

import (
	"fmt"
	"time"
)

func main() {
	sesör_veri_yolu:=make(chan string)//sensör_veri_yolu ismin de kanal oluşturduk ve diğer fonksiyonlara bu kanal bilgisini göndererek onları çağırdık.
	fmt.Println("Çıkmak için Enter Tuşuna Basınız.")
	//Eş zamanlı çalıştırıldığı için altaki mi bu mu önce çalışacağı önemsizdir.
	go Sensörden_kanala(sesör_veri_yolu)
	go kanaldan_ekrana(sesör_veri_yolu)
	//Program çalışırken, sensör_veri_yolu isimli kanalın bir ucunda sensörden_kanala fonksiyonu sürekli olarak zaman bilgisini kanala koymaya hazır olarak durmaktadır. Kanalın diğer ucunda çalışan kanaldan_ekrana fonksiyonu da 2 saniye aralıklarla kanaldan veri çekmektedir. Bu şekilde 2 saniyede bir sensör değeri (sistem zamanı) ekrana yazdırılmaktadır.
	bekle()

}

func bekle() {
	var enter_tusu string
	fmt.Scanln(&enter_tusu)
	fmt.Println("Program Tamamlandı")
}

func Sensörden_kanala(k chan string) {
	for {
		k<-"Zaman: " + time.Now().String()
	}
}

func kanaldan_ekrana(k chan string) {
	for  {
		mesaj:= <- k
		fmt.Println(mesaj)
		time.Sleep(time.Second*2)//artış miktarı denilebilir.
	}
}