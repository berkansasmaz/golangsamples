package main

import "fmt"

func main() {
	//Kesite yeni bir eleman eklemek için kullanılır.dizide yeni bir eleman eklemek için yer varsa sona eklenir ve kesitin boyutu bir arttırılır.Eğer kesitin bağlı olduğu dizide yer yoksa dizi kapasitesi 2 katına çıkartılıp eklenit

	//4 kapasiteli ve 4 elemanlı bir dizi oluşturuluyor.Sonra bu dizinin 3 elemanı temel alınarak kesit isminde yeni bir kesit oluşturuluyor
	dizi:=[4]string{"a","b","c","d"}
	kesit:=dizi[:3]//baştan 3 e kadar

	fmt.Printf(" 1: eleman sayısı=%d, kapasite=%d, %v\n ",len(kesit),cap(kesit),kesit)

	kesit=append(kesit,"x")
	fmt.Printf("2: eleman sayısı=%d, kapasite=%d, %v\n ",len(kesit),cap(kesit),kesit)//kapasite aynı yeni eleman sona eklendi yer var diye

	kesit=append(kesit,"y")
	fmt.Printf("3: eleman sayısı=%d, kapasite=%d, %v\n ",len(kesit),cap(kesit),kesit)//kapasite 2 katına çıkarıldı çünkü yer yoktu

	kesit=kesit[2:]
	fmt.Printf("4: eleman sayısı=%d, kapasite=%d, %v\n ",len(kesit),cap(kesit),kesit)//kesitin 2 indisinden sonraki  elemanları,kesitin üzerine kaydedilde yani 0,1 indisleri atıldı ve bellekte boşuna yer kaplamaması için kapasite 2 düşürüldü eleman sayının 2 katı oldu

	//kesitteki eleman sayısı len ,kapasitesi cap fonsiyonu ile öğrenilir

fmt.Println("===========================================================================")

	//Append fonksiyonu, iki kesit eklemek içinde kullanılabilir.
	ilk:=[]int{1,4}
	son:=[]int{5,3}
	butun:=append(ilk,son...)//3 nokta ... konulmassa hata verir amacı değişken sayıda argüman alabilen (variadic functions) "tüm argümanlar ,tüm elemanlar"
	fmt.Printf("%v\n",butun)

}
