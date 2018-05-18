package main

import (
	"os"
	"fmt"
)

func main() {
	//Dosya sistemindeki klasörlere erişmek için, os paketinden faydalanabiliriz. Aşağıdaki örnekte klasör içeriği listelemek için bir fonksiyon yazılmıştır. listele ismindeki bu fonksiyona argüman olarak, listelenecek olan klasörün yolu verilmektedir. Fonksiyon geriye dosya_sayisi ve klasör_boyutu isimli iki sayısal değişken döndürmektedir. Klasör içerindeki dosyaların listesini almak için, dosyalar, _ := klasör.Readdir(-1) ifadesi kullanılmaktadır. Readdir() fonksiyonu, listelenecek olan dosya sayısını argüman olarak almaktadır. Eğer bu sayı negatif olarak verilirse, klasördeki tüm dosyaları listelemektedir. Bu şekilde, dosyalar değişkenine aktarılan dosya listesi, range fonsiyonları ile isim boyutları ekrana yazdırılmaktadır.

	var dosya_sayısı int
	var klasör_boyutu int64
	dosya_sayısı,klasör_boyutu=listele("")
	fmt.Printf("Toplam %d dosya, %dKB yer kaplıyor. %n",dosya_sayısı,klasör_boyutu/1024)
}

func listele(yol string) (int, int64) {
	klasör,_ := os.Open(yol)
	defer klasör.Close()
	//Readdir(x) x<0 ise "tüm dosyalar" demek
	dosyalar, _ := klasör.Readdir(-1)
	fmt.Println("Boyut\tDosya adı\n","-------\t------")
	//Dosya boyut toplamı için int64 kullanalım:
	var toplam int64
	for _ , dosya:=range dosyalar{
		fmt.Print(dosya.Size(),"\t")
		fmt.Println(dosya.Name())
		toplam=toplam+dosya.Size()
	}
	//Dosya sayısı ve toplam boyutu döndürelim
	return len(dosyalar),toplam
}

