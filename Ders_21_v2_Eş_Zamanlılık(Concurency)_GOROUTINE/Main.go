package main

import (
	"fmt"
	"time"
)

func main() {
	//Bu sefer aynı fonksiyonu hem sıralı olarak, hemde eşzamanlı çalışacak şekilde çağırıp bakalım. Önceki program da Enter' a basılmasını bekledğimiz yerde bu sefer time paketi(kütüphanesi) altında Sleep(uyuma) fonksiyonunu kullanıyoruz.
	//kısacası eş zamanlı aynı anda bir şey okutulmak istenildiğinde kullan ama sıralamayı goruntine yapacağı için şifre kırmada antivirüste bu goruntine kullanamazsın.
	//Harfyaz fonksiyonu 5 kere sıralı olarak çalıştıralım.
	fmt.Println("---sıralı çalışan---")
	for i := 0; i < 5; i++ {
		Harfyaz(i)//fonksiyonu çağırır.
	}

	//Harfyaz fonksiyonunu 5 kere eşamanlı olarak çalıştıralım.
	fmt.Println("---eşzamanlı çalışan---")
	for i := 0; i < 5; i++ {
		go Harfyaz(i)//goruntin olarak çağır
	}
	// 1 saniye bekle, fonksiyonlar tamamlansın.
	time.Sleep(time.Millisecond*100)
}

func Harfyaz(i int)  {
	harfler:="abcçd"
	for _,harf:=range harfler {
		fmt.Printf("%d%s ",i ,string(harf))
	}
	fmt.Println()
}