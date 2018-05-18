package main

import "fmt"

func main() {
	//Bir dizinin içindeki sayıların geriye değer döndürerek ortalamasının bulunması
	aday:=[]int{40,100,15,70,45}
	fmt.Println(ort(aday))
}

func ort(elemanlar []int) int {//kesit değişkenini elemanlar ismi ile kullandık.Daha sonra range fonksiyonu kullanarak her bir puan değerini for döngüsüne dahil ederek topladık
	var toplam int
	for _,puan:=range elemanlar {
		toplam+=puan
	}
	return toplam/5  //puan:=range .... burada puan yerine toplam kullanmamızın sebebi toplamı geri değer olarak döndürmek isteğimizdir
}