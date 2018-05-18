package main

import "fmt"

func main() {
	a:=5
	b:=2
	sonuc:=topla(a,b)
	fmt.Println("Toplam (fomksiyonun cevabı):",sonuc)

}

func topla(a,b int)int{//a ve b parantez içinde int olarak belitilmiştir ayrıca dönen cevabın da parantez dışındaki int yani geri gönderdiğimiz değerin tipide int dir
	return a+b//fonksiyon sadece toplama işlemi yapıp, sonucu cevap olarak geri döndürmektedir.
}
//Değişken isimleri hem main de hemde topla fonksiyonun da aynı olmasıdır ancak bunlar farklı değişkenlerdir.Fonksiyonlar ana program beleğinden ayrı bellekte çalışır.Fonksiyon içinde tanımlanan bir değişken ,sadece o fonksiyon içinde geçerlidir.
