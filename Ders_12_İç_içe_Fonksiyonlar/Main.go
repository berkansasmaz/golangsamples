package main

import "fmt"

func main() {
	//main fonksiyonu yada bir başka fonksiyonun içinde fonksiyon tanımlaması farklı bir şekilde tanımlanır ve tanımlana fonksiyon hangi fonksiyon içindeyse onun değişkenlerini kullanabilir.

	//Yaş Kontrolü
	var yaş int

	yaşkontrol := func() {//fonksiyon yerel değişken tanımlanıyor
		switch {
		case yaş<18:
			fmt.Println("Sisteme giriş için yaşınız müsait değil. ")
		default:
			fmt.Println("Yaş onaylandı. ")
		}
	}
	//kullanıcaya yaşını sor
	fmt.Println("Yaşınız: ")
	fmt.Scanln(&yaş)

	//fonksiyonu çağıralım
	yaşkontrol()
}
