package main

import "fmt"

func main() {
//Diziden farkı kapasitesini değiştirilebilir olmasıdır dizi yerine slice kesit kullanmak daha mantıklı.

	//Slice oluşturma yöntemleri.
	var iller []string //string depolayabilen boş bir slıce.
	plaka_no :=make([]int,82) //82 tane int elemanı olan bir kesit.
	cocuklar:=make([]string,2,10) //2 string içeren ,10 elemana kadar genişleyebilen slıce.
	// MAKE fonksiyonu kesit almak için kullanılır.
        fmt.Println(iller,plaka_no,cocuklar)

	fmt.Println("========================================================================")

	//Bir dizi üzerinden kesit almak.
	var hayat [75]string
	cocukluk:=hayat[:15]//başa kadar.
	genclik:=hayat[16:30]//kesit alırken dizi isminden sonra köşeli parantezler için [... : ...] içine,kesit alınacak olan kısmın başlangıc indeksi ile bitiş indeksi (bitiş indeksi dahil edilmez) belirtilir.
	ortayas:=hayat[31:45]
	ileriyas:=hayat[45:]//sona kadar anlamında.
	fmt.Println(cocukluk, genclik, ortayas, ileriyas)

	fmt.Println("========================================================================")

	//Slice tüm elemanlarını ve her alamanın ineksini(eleman numarasını) yazandıran örnek.
	slice:=[]int{1,0,7,1}//bir kesit tanımladık ve kesitin ilk 4 elemanını tanımladık.

	for i, değer:=range slice{//range dizinin tüm elemanlarını teker teker döndürüyor.

		fmt.Printf("Elaman Numarası: %d Değer: %d\n",i,değer)
	}



}
