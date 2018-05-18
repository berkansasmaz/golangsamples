package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	//Önceki uygulamalarda dosyadan okuma yapabildiğimiz gibi, dosyaya veri yazmak için de hazır Go fonksiyonları mevcuttur.
	dosya,hata:=os.Create("test.txt")

	if hata !=nil {
		fmt.Println("Dosya oluşturulamadı")
		return
	}

	defer dosya.Close()
	dosya.WriteString("Go, Google' dan bir ile açık kaynak topluluğundan çok sayıda katılımcı tarafından geliştirilen açık kaynaklı bir projedir.")

	baytkesit,hata2:=ioutil.ReadFile("test.txt")
	if hata2 != nil {
		fmt.Println("Dosya okunamadı")
	}
	fmt.Println(string(baytkesit))
}
