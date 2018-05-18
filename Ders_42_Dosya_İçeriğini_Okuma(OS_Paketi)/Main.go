package main

import (
	"os"
	"fmt"
	"reflect"
)

func main() {
	//Program ile aynı klasörde bulunan deneme.txt isimli dosyanın içeriğiniz yazdırıcaz.

	dosya, hata:= os.Open("deneme.txt")
	if hata != nil {
		//dosya açılmazsa fonksiyondan çık
		return
	}

	defer dosya.Close()
	//dosya bilgisi al
	dosyabilgisi,_:=dosya.Stat()

	//Dosya boyutu kadar değişken oluştur
	baytkesit:= make([]byte, dosyabilgisi.Size())// Bu kesitin boyutu(eleman sayısı) dosyabilgisi.size() fonksiyonu (yani dosya boyutu kadar) ile belirlenmektedir.

	//Dosya içeriğini oku
	okunanbayt,_:=dosya.Read(baytkesit)//dosya içeriği okunup baytkesit değişkenine aktarılmaktadır. Read() fonksiyonu dosya içeriğini ASCII kodları ile okur. baytkesit değişkeni içeriğini ekrana yazdırrsak, ASCII kodlarını görürüz.

	//okunan bayt verisini metin biçimine dönüştür
	okunanmetin:=string(baytkesit)// Burası, dosyametnini anlaşılır şekilde ekrana yazdırmak için, baytkesit ifadesini metin türüne çevirdk.

	fmt.Println(dosya.Name(), "isimli dosyadan okunan", okunanbayt,"bayt şu şekildedir\n", okunanmetin)

	//Değişken türleri
	fmt.Println("baytkesit\t" , reflect.TypeOf(baytkesit))// []uint8 veri türü, yukarıdaki programon kodlarında değişken tanımlarken kullandığımız []byte veri türü ile aynıdır.
	fmt.Println("okunanbayt\t",reflect.TypeOf(okunanbayt))
	fmt.Println("okunanmetin\t",reflect.TypeOf(okunanmetin))


	//Go geliştiricileri, uint8 türüne göbek adı olarak byte ismini vermişler. ASCII kodlama sisteminde her karakter 1 bayt ile temsil edilir. Bu sayede 8 bit ile ifade edilebilecek olan karakter sayısı 2^8=256 olmaktadır. Bu nedenle byte kullanmak istediğimiz her yerde, işaretsiz 8 bitlik tamsayı (uint8) kullanabiliyoruz. İkiside tamamen aynıdır.
}
