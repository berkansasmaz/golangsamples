package main

import "fmt"

func main() {
	//Fonsiyon; Veri girişi ve veri çıkışı yapabilen, belirli bir amaca yönelik olarak tanımlanmış işlemleri yapan kod bloğudur.Fonksiyonlar cevap olarak veri dödürebilir.Örneğim;Bir fonksiyona bana rastgele sayı üretir misin ? diyebiliriz . Eğer fonksiyon bu amaçla yazılmışsa cevap olarak rasgele bir sayı döndürecektir.Fonksiyon bir değer döndürmeden de çalışabilir.Bir fonksiyon hiç bir veri almadan ve geriye değer döndürmeden de çalışabilir.Fonksiyonlar birden fazla veride geriye döndürebilir.Fonksiyonlar ihtiyaç oldukça tekrar terkar çağrılabilir.Hata bulma düzeltme, geliştirme projeyi genişletme ve sadeleştirme gibi işlemler da kolay yapılır.Her go programında en az bir fonksiyon olmak zorundadır.Önceki uygulamaların tamamında func main(){...} şeklinde tanımlanan ana fonksiyonu kullandık.

	//Fonksiyonların genelleştirilmiş yazım biçimi:
	a:=5
	b:=4
	topla(a,b)
}
//topla isimli fonsiyon aşağıda tanımlanıyor
func topla(m,n int)  {
	fmt.Println("Sayıların toplamı=",m+n)
}


