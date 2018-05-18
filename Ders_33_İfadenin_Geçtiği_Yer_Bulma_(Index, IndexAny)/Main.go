package main

import (
	"fmt"
	"strings"
)

func main() {
	//Index fonksiyonu, ifadenin metnin içinde ilk geçtiği pozisyonun sayısını (0'dan başlayarak) verir. IndexAny fonksiyonu ise ifadenin içindeki her karakteri ayrı ayrı metin içinde arayarak ilk eşleşmenin denk geldiği pozisyonun sayısını verir.

	metin:= "Korkma, sönmez bu şafaklarda yüzen al sancak;"
	fmt.Println(strings.Index(metin,"k"))//büyük\küçük duyarlı olarak "k" aramaktadır bulduğu sura numarasını ekrana yazmaktadır.
	fmt.Println(strings.IndexAny(metin,"?.,;!:"))//metni içinde noktalama işaretlerinin her birini arıyor ve ilk eşleşme bulduğu yeri belirtmektedir.
}

