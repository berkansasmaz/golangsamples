package main

import "fmt"

func main() {
	//İşaretçi bir değişkenin adresini tutan ve o değişkeni kullanıcağımız zaman değişkeni direk kullanmak yerine adresini verip kullanma kaynakları verimli kullanmak için daha yararlı.Kısacası veriler yerine verilerin adresini gönderiyoruz.

	//Örnek bir sayının karesini alıcaz.
	sayi:=5
	kare(&sayi)//Herhangi bir değişkenin önüne & simgesi konulursa, o değikenin adresini temsil eder.
	
}

func kare(x *int){
	fmt.Println(*x * *x )//Herhangi bir değişkenin önüne * simgesi konulursa işaretçinin gösterdiği adresteki veriyi temsil eder

}
