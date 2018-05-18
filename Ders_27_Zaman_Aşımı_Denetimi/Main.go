package main

import (
	"time"
	"fmt"
)

func main() {
//Eş zamanlı çalışma sırasında, işletimi uzun süren süreçlerin tamamlanması için azami süre belirlenebilir. Süre bitiminde, zaman aşımı oldu tarzında bir mesaj verilebilir. Go dilinde bununla ilgili güzel bir yöntem var. Zaman aşımı denetimi için bir kanal ve önceki başlıkta açıklanan select ifadesi birlikte kullanılmaktadır.

	//Temel mantık; tamponlu iki tane kanal kullanmak ve kanallardan birisini süre denetimi için kullanmak şeklindedir. Diğer kanal ise belirli sürede işi yapması istenen fonksiyon için kullanılmaktadır.
	gerisayım:=make(chan bool,1)
	işcevabi:=make(chan bool,1)

	//kaç saniyelik iş yapsın
	go işyap(işcevabi,1)
	//işin birmesini kaç saniye bekleyelim
	go süretut(gerisayım,2)

	select {
	case <-işcevabi:
		fmt.Println("süre bitmeden iş yetişti")
	case <- gerisayım:
		fmt.Println("Zaman aşımı oldu")
	}

}

func işyap(k chan bool, saniye time.Duration) {
	//iş yapıyorum
	time.Sleep(saniye*time.Second)
	k<-true
}

func süretut(k chan bool, saniye time.Duration)  {
	//bekliyorum
	time.Sleep(saniye*time.Second)
	k<-true
}