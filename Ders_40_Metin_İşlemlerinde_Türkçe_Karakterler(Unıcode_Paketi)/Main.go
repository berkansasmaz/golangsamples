package main

import (
	"unicode"
	"fmt"
	"strings"
)

func main() {
	//Go dili doğuştan UNICODE(universal code - evrensel karakter kodlamısı) destekli olarak gelmektedir. Önceki uygulamalarda Türkçe karakterleri hem metinlerde hemde değişkenlerde kullanmıştık. Ancak ı ve i harflerimizdeki problem Go dilini aşan bir durumdur. Çünkü tüm dünyada standart olarak kullanılan ASCII tablasuna göre i harfinin büyüğü I harfidir. Neyse ki Go bunun için de çözüm sunmaktadır. Unicode paketinde her dile uygun karakter seti tanımlanmıştır.

	tr:=unicode.TurkishCase // Bu satırda tr değişkenine Türkçe karakter seti kullanacağımız bilgisini atıyoruz.
	metin:="Ben aşığım dalgaların SESİNE"
	fmt.Println(strings.ToLowerSpecial(tr,metin))//ToLower yerine ToLowerSpecial fonksiyonu kullanarak argüman olarak tr değişkenini de fonksiyona gönderiyoruz.
	fmt.Println(strings.ToUpperSpecial(tr,metin))//Benzer şekilde, büyük harfe çevirmek içinde ToUpperSpecial fonksiyonu kullanılabilmektedir.

	//Yukarıdaki program, tüm Türkçe karakterlerde sorunsuz bir şekilde büyük\küçük dönüşümü yapılabilmektedir.
}
