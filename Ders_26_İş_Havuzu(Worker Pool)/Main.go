package main

import (
	"fmt"
	"time"
)

func main() {
	//Yapılacak çok sayıda iş varsa, bir işçi havuzu oluşturmak ve işleri bu işçilere atamak işe yarayabilir. Örneğin bir web sunucusuna gelen çok sayıda web erişim isteği, işçi havuzu ile yönetilebilir. Eş zamanlı çalışan süreçler (gorutinler) kanal kullanarak haberleştiği için, İşçi havuzları da kanallar kullanarak haberlemektedir. Bir işçi, iş kuyruğundan ilk işi alıp bu işi yaparken diğer bir işçi de kuyruktan sıradaki işi alır. Bu şekilde tüm işçiler kuyruktaki işler bitene kadar çalışırlar. Tüm işçiler aynı anda çalıştıklarından, işçi sayısı kadar eş zamanlı süreç çalıştırılmaktadır. Go dilinde bu yapıyı gerçeklemek için, gorutin ve kanalları kullanıyoruz.

	//Uygulamada 3 işçi ve 9 tane iş oluşturuyoruz. İşleri for döngüsü ile "tamsayı biçiminde" oluşturup iş_kuyruğu isimli 9 tamponlu kanala gönderiyoruz. İş verip de sonucunu takip etmemek olmaz. Bu nedenle işlerin sonucunu da beklememiz gerekiyor. Bu amaçla kullanmak üzere iş_cevapları isimli bir kuyruk daha oluşturuyoruz. İşleri yapmak üzere, işçi isimli bir fonksiyon geliştirilmiştir. Bu fonksiyona argüman olarak; işçi numarası(işçi_no), yapılacak işlerin bilgisini taşıyan kanalın ismi (işin_tanımı) ve tamamlanan işlerin bilgisini alabilmek için kullanılan kanalın ismi(iş_cevapları) gönderiliyor.
	iş_sayısı:=9
	işçi_sayısı:=3
	//İş kuyruğu ve iş cevapları için
	//İş sayısı kadar tamponlu kanal oluşturalım
	iş_tanımı:=make(chan int, iş_sayısı)
	iş_cevabı:=make(chan int,iş_sayısı)

	//3 tane işçi oluşturalım
	for i := 1; i <= işçi_sayısı; i++ {
		go işçi(i,iş_tanımı,iş_cevabı)
	}

	//9 tane iş uydurup, kuyruğa koyalım
	for iş := 1; iş <= iş_sayısı; iş++ {
		iş_tanımı<-iş
	}

	//işlerin cevaplarını bekleyelim
	for a := 1; a <= iş_sayısı; a++ {
		<-iş_cevabı
	}
	fmt.Println("----\nTüm İşler Tamamlandı.")
}

func işçi(işçi_no int, işin_tanımı chan int, iş_cevabı chan int){
	for iş:=range işin_tanımı{
		fmt.Println(işçi_no, ". İŞÇİ", iş ,". İŞİ YAPIYOR")
		//Bir saniye sürsün
		time.Sleep(time.Second)
		iş_cevabı<-iş
	}
	//Süreçler eş zamanlı olarak işletildiği için, iş numaraları sıralı olarak ilerlememektedir. İşlerin tamamı aynı sürede bitecek şekilde planlandığı için, işçiler eşit sayıda iş yapmaktadır. İşlerin süresi farklı farklı olduğunda, İşçiler de farklı sayıda iş yapabilmektedir. Ancak her işçini süre olarak çalışma miktarı aynı seviyelerdedir.
}
