package main

import (
	"os"
	"fmt"
)

func main() {
	dosya, hata:=os.Open("deneme.txt")
	if hata!=nil {//dosya bulunamazsa  //dosya açma işlemi sırasında hata isimli değişkene bir değer atanıp atanmadığını kontrol ediyor. Eğer dosya başarılı bir şekilde açıldıysa, hata değişkenine bir değer atanmayacağı için nill(boş) olarak kalmaktadır.
		println("Dosya açılmadı. Gelen hata mesajı:\t" , hata)
		return
	}

	//çıkarken (dosyayı açık unutmayalım) :
	defer dosya.Close()
	fmt.Println("Dosya başarılı bir şekilde açıldı")
	//bundan sonra dosya ile ilgili işlemleri devam ettirebiliriz.
}
