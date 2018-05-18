package main

import "fmt"

func main() {
	//Yapı(Struct) programlamada kaba tabirle, birden fazla değişkenin bir arada paketlenmiş hali gibi düşünebiliriz.
	//Örnek, farklı türden ulaşım araçlarını kaydetmek için istediğimiz değişkenleri kullanıcaz
	oto1:=vasıta{2010,25000,"bordo","4*4"}//oto1 değişkeni oluşturulurken içine bazı değerler girilmiş
	var tosbaa vasıta//tosbaa değişkeni oluşturulduktan sonra bazı alanlarına veriler girilmiş.
	mx:=vasıta{tür:"motosiklet",yıl:2010,fiyat:15000}//mx değişkeni oluşturulurken de alanların isimleri belirtilerek bazı alnkara veri girilmiş.
	//3 tane yapı değişkeni oluşturuldu
 	//oto1,tosbaa,mx diye 3 değişken oluşturduk ve bu değişkenlerin türü vasıta olarak belirlendi.Bu demek oluyor ki değişkenlerin  aşağıda belirttiğimiz 4 tane alanı(özelliği) olabilir.

	tosbaa.yıl=1980
	tosbaa.fiyat=10000
	tosbaa.tür="3kapı"

	fmt.Println(oto1)
	fmt.Println(tosbaa)
	fmt.Println(mx)
	fmt.Println("Toplam Fiyat=", oto1.fiyat + tosbaa.fiyat + mx.fiyat)
}
//vasita isminde bir tür oluşturalım.Bu tür içinde 2 tamsayı, 2 metin alanı olsun.
type vasıta struct {//bir veri türü oluşturduk.
	yıl, fiyat int//yıl,fiyat,renk,tür bunlara alan diyoruz 4 farklı veri tutula biliyor
	renk, tür string
}
