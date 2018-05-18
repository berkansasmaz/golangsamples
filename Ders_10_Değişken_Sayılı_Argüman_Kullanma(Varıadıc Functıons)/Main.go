package main

import "fmt"

func main() {
	//Argüman topla(a,b) burada ki parantez içerisinde olan a ve b ye argüman diyoruz. Eğer bu argümanların sayısı belli değilse Variadic Functıons kullanılır.

	//Topla fonksiyonunu değişken argümanlı uygulaması
	fmt.Println(topla(-5, 4))
	fmt.Println(topla(1,9,2,3))
}

func topla(toplanacaksayılar ...int) (int) {

	toplam:=0
	for _,sayılar:=range toplanacaksayılar{//Argüman sayılarını bilmediğimiz için for i:=0;i<5;i++ gibi bir yazım kullanamayız. Toplanacaksayılar,range fonsiyonu kullanarak for döngüsüne sokuldu.
		toplam+=sayılar
	}

	return toplam
}