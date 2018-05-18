package main

import "fmt"

func main() {
	//Bir dizi içerisinde farklı türde veri tutulamaz.

	var il [82]string//82 kullanmamızın sebebi dizi indisleri sıfırdan başlar ama plakalar 1 den başlar
	il[1]="Adana"//dizinin sıfırıncı elemanı kullanılmamıştır
	il[54]="Sakarya"
	il[16]="Bursa"
	il[47]="Mardin"

	fmt.Println(il)//indis numarasına göre sıra ile yazılmıştır
	fmt.Println(il[54])//sadece bir eleman yazdırmak için
	//dizinin tanımlanmayan elemanlarıda bellekte yer tutar
        fmt.Println("==================================")
	//sınava giren 5 kişinin ortalaması
	var sınavagiren [5]float32//ortalamanın sonucunun ondalık sayı cıkma olasılığı da var
	var toplam float32
	sınavagiren[0]=81
	sınavagiren[1]=100
	sınavagiren[2]=65
	sınavagiren[3]=95
	sınavagiren[4]=45

	for i := 0; i <5; i++ {
		toplam += sınavagiren[i]
}
	fmt.Println("Toplam Aday Sayısı:",len(sınavagiren))// len dizinin tüm elemanlarının sayısını yazdırmada
   	fmt.Println("Ortalama=",toplam/5)
	fmt.Println("==================================")

	//5 kişini başarılı olup olmadığını söyleyen uygulama
	//ama len fonksiyonun bir değişkene atanmış hali

	aday:=[5]float32{81,100,27,95,45,}// En sona virgül koymamızın amacı yazılım ihtiyacına göre istenen satırların kolayca değre dışı bırakılmasıdır
        //dizi elemanları böylede kolayca atanabilir
	var adaysayisi int =len(aday)//5 tane  aday olduğu adaysayına eşitlendi
	for i := 0; i < adaysayisi; i++ {
		if aday[i]<50 {
			fmt.Println(i,". aday başarısız oldu:",aday[i])
		}
	}
	fmt.Println("==================================")
	//üsteki uygulamanın aynısı 5 kişinin başarılı başarızı

	sınavadayı:=[5]int{81,100,27,95,45}
	for i,puan:=range sınavadayı {//range fonksiyonu çok sayıda veri üzerinde foru döndürmede kullanılır
	//i döngü değişkeni
			if puan < 50 {
			fmt.Println(i,". aday başarısız oldu",puan)
		}
 	//i döngü değişkenini kulanmıycaksak go da kullanılmayan değikenlerin yerine _ konulabilir

	}


	fmt.Println("==================================")

	//Dizilerde Çok boyutlu Veri kaydetme Matris
	nokta:=[3][2]int{{11, 12},{21, 22},{31, 32}}//istersek 3*3 lük matrisler bile tanımlayabiliriz
	fmt.Println("Dizinin Tamamı:",nokta)//dizinin tamamını yazdırdı
	fmt.Println("\n3*2 Matris Biçiminde:")
	for satır := 0; satır < 3; satır++ {
		for sutun := 0; sutun < 2; sutun++ {
			fmt.Print(nokta[satır][sutun],"  ")
		}
		fmt.Println()//her döngü sonunda alt satıra geçmesi için
	}

}

