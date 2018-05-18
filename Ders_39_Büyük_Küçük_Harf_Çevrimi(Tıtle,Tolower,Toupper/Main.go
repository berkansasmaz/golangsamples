package main

import (
	"fmt"
	"strings"
)
//Türkçe alfabeye özgü karakterlerde, i ve ı harflerin büyük\küçük durumunu ingilizce alfabesine göre farklı olmasından dolayı sorun çıkmaktadır. Bunun dışındaki harflerde sorun olmamaktadır.

//İkis satırda Title fonksiyonu kullanımıştır. Bu fonksiyon, her sözcüğün ilk harfini büyük harfle çevirmekte, diğer harflerrini değişmemektedir.

//ikinci satırda ToLower fonksiyonu kullanılmıştır. Çıktıda görüleceği üzere, tüm harfler küçük harfe dönüştürülmüştür.

//Üçüncü satırda ToUpper fonksiyonu kullanılmıştır. Tüm harfler büyük harfe döndürülmüştür.

func main() {
	metin:= " BeN aşığIM dalgalarıN sesinE"
	fmt.Println(strings.Title(metin))
	fmt.Println(strings.ToLower(metin))
	fmt.Println(strings.ToUpper(metin))
	fmt.Println(strings.Title(strings.ToLower(metin)))
}
