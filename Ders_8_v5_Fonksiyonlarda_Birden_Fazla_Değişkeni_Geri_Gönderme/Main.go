package main

import "fmt"

func main() {
	a :=1 //renema ile bir kodu değiştiriken tüm satırlada değitirmesini sağlayabiliriz
	b :=2
	fmt.Println(a, b)
	değiştir(a, b)
	fmt.Println(değiştir(a,b ))
}

func değiştir(x, y int)(int,int)  {
	return y,x
}
//Bir fonksiyon ,cevap olarak brden fazla değer dödürebilir.Bu duruda ,fonsiyonun tanımlandığı satıda func değiştir(x,y int)(int,int) ikinci parantez deki int int kavramı geriye döndürülecek olan değerlerin ikisinin de int olduğudur.
