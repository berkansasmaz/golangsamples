package main

import (

	"fmt"
	"encoding/json" //sözlük verilerimisi farklı bir biçimde görmek için kullandık
)

func main() {
	//Dizide tutulan veriler sayısal bir indeks ile erişilebilen verilerdi. Harita ise sırasız ve indeksi olmayan verlerde çalışmak için tasarlanmıştır.Harita yapısında kullanılan ikililere ,anahtar(key) ve değer(value) ismi verilir.Map isim olarak sözlük veri yapısı ,ilişkisel diziler,tablo(hash table) gibi farklı kullanımları mevcut

	//Örnek bir harita tanımlama
//	var il_nufus map[string]int  //ilnufus isimli harita değişkeni tanımladık
//	il_nufus=make(map[string]int)	//il_nufus adındaki haritayı oluşturduk.


	//Kısaca Örnek Harita oluşturma
//	il_nufus:=make(map[string]int)
//	[string] kısmı anahtarların veri türünü belirtiyor
//	int kısmı değerler bölümünün veri türünü belirtiyor

	il_nufus:=make(map[string]int)
	il_nufus["Bilecik"]=210000
	il_nufus["İstanbul"]=15000000
	il_nufus["Sakarya"]=1000000
	fmt.Println("Bilecik:",il_nufus["Bilecik"])
	fmt.Println(il_nufus)//Haritalar dizierin aksine sıralı değildir

	fmt.Println("==================================================")

	//Basit Bir Bilişim Kısaltmaları Sözlüğü
	sozluk:=make(map[string]string)
	sozluk["IP"]="Internet Protocol"
	sozluk["LAMP"]="Linux Apache MySQL PHP"
	sozluk["GNU"]="GNU is Not Unix"
	sozluk["WWW"]="Word Wide Web"
	sozluk["IMAP"]="Internet Message Access Protocol"
	tum_elemanlar, _ :=json.MarshalIndent(sozluk,""," ")//Json ismindeki veri sunma biçimi,günümüzde farklı platformlar arasında veri aktarımı konusunda en popüler yöntemlerden birisidir.
	fmt.Println(string(tum_elemanlar))

	fmt.Println("==================================================")

	//Aynı uygulama(sözlük) lakin blok olarak veri girdik ayrı ayrı yazmadık ve haritada tüm verileri okumak için range fonksiyonunu kullandık
	Nufus:=map[string]int {//blok halinde kullanımda make gerek yok çünkü kesit değil dizi sayısı belli

		"Bilecik":210000,
		"İstanbul":15000000,
		"Sakarya":1000000,

	}

	for i ,değer:=range Nufus{
		fmt.Printf("İl : %s\t Nüfüs : %d\n",i,değer)

	}

	fmt.Println("==================================================")

	//Bir Harita oluşturup içinden bir kaydı(anahtar ve değeri)silicez.Sonra bu kayıt yazdırmak istendğinde bulunamadı mesajının verilmesini sağlayacağız
	sozluk2:=make(map[string]string)
	sozluk2["IP"]="Internet Protocol"
	sozluk2["GNU"]="GNU is Not Unix"
	sozluk2["WWW"]="Word Wide Web"
	delete(sozluk2,"GNU")//silme işlemi delete fonksiyonu ile gerçekleşir
	if karşılık, sonuc := sozluk2["GNU"]; sonuc {//bu kısımda iki işlem yapılıyor birincisi karşılık,sonuc:=sozluk2["GNU"] da atama işlemi yapılıyor sözlükten bir kısaltma istenerek bunun karşılık bir değişkenine atanması  isteniyor.Ancak burada sonuç adında bir değişken daha verilmiş.Buradan anlıyoruz ki, bir harita değerini başka değişkene atarken, iki farklı değer dönüyor.Bir tanesi anahtara ait değer,d,ğeride (true/false şeklinde) işlemin başarılı olup olmadığı bilgisi.GNU varsa sonuc değişkeninde true değeri vardır.Atama başarısız ise sonuç değişkeni false dır.
		fmt.Println(karşılık)
	} else {
		fmt.Println("Aranan Sözcük Bulunamadı")

	}

	fmt.Println("==================================================")


	//Bir Şirket Yıl Sonunda envanter sayımı yapıyor.Bilgisayar türleri ve sayılarını harita değişkenine kaydetmişler.Toplam PC sayısı
	toplam:=0
	bilgisayar:=map[string]int{
		"Dizüstü":27,
		"Masaüstü":123,
		"Tablet":8,
		"Sunucu":3,

	}
	for _,değer:=range bilgisayar {

	toplam=değer

	}
	fmt.Printf("Toplam %d bilgisayar var.\n",toplam)









}
