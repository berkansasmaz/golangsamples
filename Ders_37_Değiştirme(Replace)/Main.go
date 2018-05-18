package main

import (
	"fmt"
	"strings"
)

func main() {
	//Metin içerisindeki bir ifadeyi başka bir iade ile değiştirmek için Replace fonksiyonu kullanılır. Bu fonksiyon 4 tane argüman almaktadır. İlk argüman, işlenecek olan metindir. İkinci ve üçüncü argüman, birbiri ile değiştirilecek olan ifadelerdir. Son argüman ise metinde bulunan ifadelerden kaç tanesinin değiştirileceğidir. Son argüman olarak negatif bir sayı verilmesi durumunda, metinde bulunan tüm ifadeler değiştirilecektir.

	metin:="kedileri çok severim, sokaktaki kedilere yiyecek veririm"
	fmt.Println(strings.Replace(metin,"kedi","kedi ve köpek",1))
	fmt.Println(strings.Replace(metin,"kedi","köpek",-1))// -1 ile tüm kedi olan yerlerin köpek olarak yer değiştiridik

}
