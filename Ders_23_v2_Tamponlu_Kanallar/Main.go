package main

import "fmt"

func main() {
	//Kanaldan gelen veriler üzerinde dizi veya kesit işlemlerinde olduğu gibi range fonksiyonu kullanılabilir.
	//Başta kanalın kapasitesini 5 olarak belirleyip sonra 3 tane mesaj tanımladğımız da sorun olmaz.
	sayılar:=make(chan string,5)
	sayılar <- "1"
	sayılar <- "2"
	sayılar <- "3"

	//kanalı kapatalım
	close(sayılar)// kanaldan gelen tüm verilerin işlenmesindei verilerin tamamlandığı veya gelmeye devam edeceği bilinmemektedir. Bu nedenle veriler işlenmeden önce kanalın kapatılması gerekmektedir.
	for sayı:=range sayılar{
		fmt.Println(sayı)
	}
}
