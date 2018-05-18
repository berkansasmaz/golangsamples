package main

import (
	"fmt"
	"strings"
)

func main() {
	// Metnin her bir karakterinin bir kurala göre değiştirilmesi için Map fonksiyonu kullanılabilir. Bir metni şifreleme ihtiyacı varsa, bu fonsiyondan faydalanılabilir.
	//Bu bölümde Sezar şifreleme algoritmasının benzeri olan ROT13, ASCII tablosunda tanımlı olan 26 harf üzerinde kullanılmak üzere tasarlanmıştır. Tablodaki her harfi, kendisinden 13 sonraki harf ile değitirerek şifreleme sağlar. 13 sayısı rastgele seçilmemiştir, şifrelenen bir harfe yeniden ROT13 uygulandığında orjinal metni vermektedir.

	rot13:= func(karakter rune) rune{
		return karakter+13
	}
			fmt.Println(strings.Map(rot13,"klmn"))

	//Programda kullanılan rune veri türü, int32' nin diğer adıdır.
}
