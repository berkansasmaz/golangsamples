package main

import "fmt"

func main() {
	//Kanalın bir hafızası olsa, fonksiyonun birisi kanala bir mesaj bıraksa ve gitse, sonra başka bir fonksiyon kanaldan bu mesajı alıp işlese olur mu? Evet, buda mümkün. Kanalların bu tarz kullanımına tamponlu kanal ismi veriliyor.
	//Varsayılan olarak kanallar tamponsuzdur.Tamponsuz kanal, sadece canlı iletşimde kullanılabilir. Tamponlu kanal kullanmak için, kanal oluşturulurken tampon miktarı belirtilmelidir.
	//Tamponlu kanal kullanabilmek için, kanal oluşturulurken kanalın tampon kapasitesi(tutabileceği mesaj sayısı) belirtilmelidir.
	//Kanal tamponu, first In first out(İlk giren ilk çıkar) mantığı ile çalışır.

	tamponlukanal:=make(chan string,2)
	tamponlukanal <- "Birinci Mesaj"
	tamponlukanal <- "İkinci Mesaj"

	//Kanaldaki mesajları alalım
	fmt.Println(<-tamponlukanal)
	fmt.Println(<-tamponlukanal)
}
