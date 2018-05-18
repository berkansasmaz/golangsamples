package main

import (
	"fmt"
	"strings"//if !strings.Constain ifadesinde strings kütüphanesini kullan ve onun içinde ki constain fonksiyonunu kullan dedik.
)

func main() {
	//Bir metin içinde başka bir metni aramak için kullanılır. Cevap mantıksaldır (true\false).
	adresler:=[]string{"ali@example.com", "veli.example.com", "ayse@com"}
	for _,adres:=range adresler{//kesit döngüye sokuldu
		//İf !strings......... ifadesinde adres değişkeni içerisinde "@" arıyor. if sözcüğünden spnra gelen ! işareti, if koşulunun (@ işareti içeren) tersini belirtmektedir.
		if !strings.ContainsAny(adres,"@"){//döngüye girenlerde @ varmı yokmu diye bakıldı.Yoksa hata mesajı veriyor.
			fmt.Println("Hatalı e-posta adresi: ", adres)
		}
	}
}
