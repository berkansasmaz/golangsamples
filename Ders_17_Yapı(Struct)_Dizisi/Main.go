package main

import "fmt"

func main() {
	//Yapı değişkenlerini kullanırken,her bir değişkenin normal olarak bir ismi olması gerekiyor.Bazen değişkene isim vermek yerine, değikeni dizi olarak tanımlamak ve ismi de yapının bir parçası haline getirmek daha kullanışlı olabilir.
	//Örnekte; kuşlar adında bir yapı tanımlanıyor. Sonra kuşlar türünde, papağan isimli birdizi tanmlanıyor.
	var papağan [3]kuşlar
	papağan[0]=kuşlar{yaş:1,renk:"beyaz",isim:"Çaki",tür:"sultan"}
	papağan[1].yaş=3
	papağan[1].renk="mavi"
	papağan[1].tür="Yeşil_Papağan"
	papağan[1].isim="inci"

	//Ekrana yazdıralım
	for _, i :=range papağan  {
		fmt.Println(i)
	}

}

type kuşlar struct {
	tür, renk, isim string
	yaş int
}
