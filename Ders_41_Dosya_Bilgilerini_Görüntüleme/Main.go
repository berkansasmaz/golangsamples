package main

import (
	"os"
	"fmt"

)

func main() {
	//Bilgisayarda kayıtlı olan bir dosya hakkında bilgi almak için, dosyanın open() fonksiyonu ile açılması gerekir. Aşağıda basit bir uygulama verilmiştir. Bu program, çalışılan klasördeki deneme.txt dosyasının bilgilerini ekrana dökmektedir. Çalışıln kolasörde deneme.txt dosyası yoksa, program hata verecektir.

	dosya,_ :=os.Open("deneme.txt")//Bu ifade ile dosya açılmaktadır.Eşitliğin sol tarafındaki dosya ifadesi açılan dosyanın dosya isimli değişken tarafından ifade ediliceğini gösteriyor. Alt çizgi(_) ifadesi ise os.open() fonksiyonundan geriye dönecek olan iki değerden ikincisini önemsemediğimiz anlamına gelmektedir.
	//Çıkarken (defer) dosyayı açık unutmayalım.
	defer dosya.Close()
	//Dosya bilgilerini alalım
	dosyabilgisi, _ :=dosya.Stat()
	fmt.Println("Dosya ismi", dosyabilgisi.Name())
	fmt.Println("Klasör mü", dosyabilgisi.IsDir())
	fmt.Println("Yetkileri", dosyabilgisi.Mode())
	fmt.Println("Değişim tarihi", dosyabilgisi.ModTime())
	fmt.Println("Boyut", dosyabilgisi.Size(),"bayt")

	//program çalıştırıldığında deneme.txt dosyası bulunmazsa, biraz kızmakta ve şöyle bir mesaj vermektedir.
	//panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentation violation code:0*1 addr=0*38 pc:0x4010b7]

}
