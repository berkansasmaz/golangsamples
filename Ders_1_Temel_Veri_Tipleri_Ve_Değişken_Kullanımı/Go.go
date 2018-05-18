package main

import (
	"fmt"

	"reflect"
)

func main() {
	//Integer Tamsayı tanımlama
	var sayi1,sayi2 int
	sayi1=3
	sayi2=sayi1+5
	fmt.Println("Sayıların Toplam:" ,sayi1+sayi2)

	//Float Ondaklık sayı tanımlama
	var bolum float64
	bolum=32.0/5.0//8 bölü 25 diyemeyiz float olarak tanımladığımız için .0 gibi ondalık sayı kullanmalıyız
	fmt.Println("Bölümün Sonucu:",bolum)

	//String Karakter Tanımlama
	var ad,soyad string
	ad="berkan"  // bir tab boşluk için \t bir alt satır için \n
	soyad="\tşaşmaz"
	fmt.Println(ad+soyad)

	//Boolean Mantıksal ifade tanımlama

	//   &&=and(ve)
	//   ||=or(veya)
	//    !=NOT(değil)

	var isim,parola,eposta bool
	isim=true
	parola=true
	eposta=false
	fmt.Println(isim&&parola)
	fmt.Println(eposta||isim)
	fmt.Println(!isim)
	fmt.Println("Manisa"=="Turgutlu")

	//Constant Sabitleri tanımlama

	    //Değeri değiştirilemeyen değikenlere sabit denmektedir program süresince değeri değişmeyecek veriler için kullanışlıdır

	const Çanakkale int =17
	const Ülke string = "TR"
	fmt.Println(Çanakkale,Ülke)
	//Değişken tanımlanırken aynı anda değer ataması da yapılırsa ,bu değer başlangıç değeri ismi verilir

	//Değişken Tanımlama Yöntemleri

	//1. Yöntem
	var isim2,parola2,eposta2  bool
	isim2=true
	parola2=true
	eposta2=false
	fmt.Println(isim2,parola2,eposta2)


	//2. Yöntem
	const Ülke2 string = "TR"

	//3. Yöntem
	var a=3
	var b=5
	var c="go"

	fmt.Println("Toplam:",a+b)
	fmt.Println(reflect.TypeOf(a),reflect.TypeOf(c))

	/*
	var sözcüğünü kaldırıp = yerine := kullanıla bilir
	*/

	Adım:="berkan"
	soyadım:=" Şaşmaz"
	Adım="Burak"//bir sonraki kullanımda yine = kullanabiliriz
	fmt.Println(Adım+soyadım)

	//değişkenleri sadece func mainde kullanılmasını istemiyorsak func mainden önce değişkenleri tanımlamamız lazım

	//Çok fazla değişken tanımlaması şu şekilde yapılabilir
	const
	(
		çanakkale =17
		bilecik=11
		bursa=16

	)
	fmt.Println(Çanakkale,bilecik,bursa)


}
