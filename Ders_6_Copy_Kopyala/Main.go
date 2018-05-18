package main

import "fmt"

func main() {

	// Copy fonksiyonuna argüman olarak iki tane kesit verilir. Sağdaki kesitin(kaynak içeriği),soldaki kesite (hedef)kopyalanır.Burada copy fonksiyonun avantajı ; iki kesitin de eleman sayılarını kontrol etmesi ve en az olan eleman sayısı kadar elemanı kopyalamasıdır
	kaynak:=[]int{1,2,9,9}
	hedef:=make([]int,3)//2 kapasiteli kesit
	fmt.Println(kaynak,hedef)
	sayi:=copy(hedef,kaynak)//sağdaki kesit kaynak soldaki kesite hedefe kapasitesi kadar depolanır
	fmt.Println("Kopyalanan eleman sayısı", sayi)//copy fonksiyonun sayı isimli değişkene aktarıp yazdırdık istersek sadece copy(hedef,kaynak) yapabilirdik
	fmt.Println(kaynak,hedef)

}
