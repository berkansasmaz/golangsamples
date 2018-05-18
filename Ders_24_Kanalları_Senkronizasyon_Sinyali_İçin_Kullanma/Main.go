package main

import (
	"fmt"
	"time"
)

func main() {
	//Bir gorutinin tamamlanmasını bekleyip sonra başka bir iş yaptırmak için, bir fonksiyonu kanalın başına koyup gorutinin kanala mesaj göndermesini bekleyebiliriz.
	kontrol_yolu:=make(chan string,1)
	go formatla(kontrol_yolu)
	<-kontrol_yolu//Bu satır, kanalın ucundan ses bekleyen kısımdır. Kanalın diğer ucundan ses gelince sonraki satıra geçip devam eder.//printle yazarsak temizlik tamamlandı da yazdırılır program calıştığında

}

func formatla(kontrol_yolu chan string) {
	fmt.Println("Sabit Disk Biçimlendiriliyor...")
	time.Sleep(time.Second*2)
	kontrol_yolu <- " Temizlik Tamamlandı."
}
