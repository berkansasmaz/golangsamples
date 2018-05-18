package main

import "fmt"

func main() {
	//Fonksiyon sonlanırken çalışmasını istediğimiz komutları defer ifadesi ile belirtebiliriz. Kullanımı evden çıkarken elektiriği kapatayım diye düşünüpte çıktıktan sonra tereddüt ederiz ya bazen, tam da bu durumlar için tasarlanmış defer ifadesi.Siz istediğiniz zaman yazıyorsunuz ama o mutlaka çıkarken yapıyor o işi.
	//Dosya işlemleri yaparken,işlem yapmadan önce dosyanın açılması,işi bitince de kapatılması gerekir. Kapatılmazsa bellekte yer kaplamasından tutun da dosyanın bozulmasına kadar farklı sorunlar yaşanabilir. Bu durumlarda mesela dosya açılı açılmaz arkasından defer ifadesi ile kapatma fonksiyonu yazılır. Bu sayede, iş bitince dosyanın kapatılması garanti edilmiştir

	// *defer ifadesi kullanıldığı anda, ifadede kullanılan değişkenlerin değeri ne ise o değerler defer için sabitlenirç Yani değer değişse bile, fonksiyon sonunda ifade çalışırken, ilk baştaki değerlerle çalışır.
	// *Birden fazla defer ifadesi varsa,son giren ilk çıkar mantığı işler.
	
	//Uygulama
	işlemyap()
	
}

func işlemyap()  {
	i:=0
	defer fmt.Println(i) //2. sırada işletilecek ilk giren bu olduğu için son çıkıcak.
	i++
	defer fmt.Println(i) // 1.sırada işletilecek.
	fmt.Println("Bu satır defer' lerden önce işletilir. ")

}
