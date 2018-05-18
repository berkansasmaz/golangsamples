package main

import "fmt"

func main() {
//C# taki derslerden aşina olduğum bir konu kalıtım miras alma(İnheritance) gibi düşünebilirsin
//Gömülü türler, yapı şeklinde tanımlanan bir türe ait elemanları(alanları ve metotları) başka bir yapı türü içinde kullanabilmeyi sağlar.

//Bir türü başkka bir tür içine gömdüğümüzde, dıştaki türün tüm metotları içteki türe de katılmış olacaktır. Ancak bu metotlar çalıştırıldığında, metodun alıcısı dıştaki değil, içteki tür olacaktır.

	tren:=taşıt{}
	tren.Adı="Haydarpaşa-Adapazarı Ekspresi"
	tren.ilerle() // taşıt' ın ilerle() metotunu çağırıyoruz.

	kayık:=deniz_taşıtı{}
	kayık.Adı="Ertuğrul" //taşıt altında tanımlanmış olan "Adı"alanı
	kayık.yelken=false //yelkeni yok
	fmt.Println(kayık)
	fmt.Println("// ----metotlar çağrılıyor----")
	kayık.taşıt.ilerle() // taşıt' a ait olan metotu çağrıyoruz.
	kayık.ilerle() // aynı metodu doğrudan çalıştıralım.
	fmt.Println(kayık.taşıt.Adı)
	fmt.Println(kayık.Adı)


}

type taşıt struct {//taşıt isminde bit yapı(tür) tanımlanıyor
	Adı string
}

func (t taşıt) ilerle()  {//taşıt türüne bağlı ve ilerle() isminde bir metot oluşturduk
	fmt.Println(t.Adı, "İleri doğru gidiyor...")

}

type deniz_taşıtı struct {//deniz taşıtı isminde yeni bir tür tanımladık.
	taşıt // ***Gömülü türü tanımlama***
	yelken bool //deniz_taşıtı' nda, yelken isimli bir mantıksal alan olsun
}
