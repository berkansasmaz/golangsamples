package main

import "fmt"

func main() {
	//Gorutin kullanımı birden fazla yerden ver beklendiği durumlarda çok işe yaramaktadır.Örneğin birkaç web sitesinden veri çekmek gerektiğine veya farklı veritabanı sunucularından veriler almak gerektiğinde ya da birden fazla sensörden aynı anda veri çekilmek istendiğinde, işlerin birbirini beklememesi için gorutin kullanılabilir.
 	//GORUTIN(GOROUTINE)
	//Fonksiyon çağrılırken başına go eklememiz yeterlidir.
	//Gorutinlerin çalışması sırasında yönetim go dilindedi. Hangi sürecin hangi aşamada ne iş yapacağını tam olarak bilemeyiz. Bu nedenle sıralı işlemesi gereken süreçleri gorutin olarak çalıştırmamalıyız.
//Kısacası her çalıştırmada farklı sonuç ne nerden belli değil büyük projelerde avantajı bellek kulanımı dezavantaj karışık.
	go f1()
	go f2()
	//Programdan çıkmadan Enter' a basılmasını bekleyelim
	var EnterTuşu string
	fmt.Println("Çıkmak için Enter tuşuna basın")
	fmt.Scanln(&EnterTuşu)
}

func f1()  {
	var toplam int
	for i:=0;i<10 ;i++  {
		toplam+=i
	}
	fmt.Println(toplam)
}

func f2() {
	harfler:="abcdefghij"
	for _,harf:=range harfler {
		fmt.Println(string(harf))
	}
}