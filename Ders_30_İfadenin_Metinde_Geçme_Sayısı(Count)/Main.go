package main

import (
	"strings"
	"fmt"
)

func main() {
	//Bir metindeki sözcükleri saymanın en kolay yolu boşlukları saymaktır.Count fonksiyonu kullanarak metindeki sözcük sayısı bulunmaktadır:

	istiklalmarsi :=`
		  Korkma, sönmez bu şafaklarda yüzen al sancak;
	    Sönmeden yurdumun üstünde tüten en son ocak.
	    O benim milletimin yıldızıdır, parlayacak;
	    O benimdir, o benim milletimindir ancak. `

	boşluksayısı:=strings.Count(istiklalmarsi," ")
	satırsayısı:=strings.Count(istiklalmarsi,"\n")
	fmt.Println("Metindeki sözcük sayısı= ", boşluksayısı+satırsayısı-1)
}
//Uzun metinleri birden fazla satırda yazmak için, tırnak işareti(") yerine ters tırnak işareti (`) işareti kullanılabilir.
