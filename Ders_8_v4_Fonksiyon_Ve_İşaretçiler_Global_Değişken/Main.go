package main

import "fmt"

//Gloabal değişkenler fonksiyon dışında tanımlanırlar
var a,b int=1,2//bütün fonksiyonlarda kullanmak için global değişken tanımlarız publicc
//global değişkenin dezavantajı garbage collector bunları temizlemez bellekte gereksiz yer kaplar ama fonksiyon içinde tanımlarsak fonksiyonun işi bittiğinde bellek alanı geri iade edilir.
//Özel durum olmadıkça global değiken kullanma

func main() {

	fmt.Println(a,b)
	değiştir()
	fmt.Println(a,b)

}

func değiştir()  {

	gecici:=a//gecici isminde değişken oluşturduk a ya atadık yani a 1 di gecici de 1 oldu
	a=b//a yı b ye eşitledik yani b 2 idi a da 2 oldu
	b=gecici//b yi gecici ye eşitledik yani gecici 1 idi b de bir oldu


}
