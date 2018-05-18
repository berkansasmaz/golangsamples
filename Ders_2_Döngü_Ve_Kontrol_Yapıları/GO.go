package main

import "fmt"

func main() {
	//go da tek bir döngü vardır oda For dur
	//for kullanımı 3 ifadede gerçekleşir değişkenin başlangıcı ,kontrol değişkeni ve ilerleme kuralı
	for i := 1; i <= 4; i++ {
		fmt.Println(i)
	}

	//foru while döngüsü gibi kullanma
	i:=1
	toplam:=0
	for i<=10  {//Başlangıç ifadesi ve ileleme kuralı pas geçilebilir

		i++
		toplam=toplam+i
	}

	fmt.Println("Toplam:",toplam)

	//IF Karşılaştırma Fonksiyonu
	for i:=1;i<=10 ;i++  {
		if i % 2==0 {//2'ye bölümünden kalan sıfır mı?
			//sayı çift ise çalışsın
			fmt.Printf("%d ",i)//println yerine printf kullanmamızın sebebi sayıların alt alta değil yan yana yazılmasını sağlamak.
		}

	}

	//else if
	var ı, b2,b3,b4  int

	for ı = 1; ı <= 100; ı++ {
		if ı%2==0 {
			b2++
		} else if ı%3==0{
			b3++

		} else if ı%4==0{//çalıştırdığımızda 4 ile bölünebilen sayı 0 olarak gösteriyor sebebi 4 ile kalansız bölünebilen sayı 2 ilede kalansız bölünebilir
			b4++
		}
	}
	fmt.Printf("2:%d,3:%d,4:%d",b2,b3,b4)

	//switch Kullanımı
        sınavnotu:=80
	switch  {
	case sınavnotu < 40:
		fmt.Println("Çok Düşük!")
	case sınavnotu < 55:
		fmt.Println("İyi Değil!")
	case sınavnotu < 70:
		fmt.Println("İdare Eder!")
	case sınavnotu < 80:
		fmt.Println("Güzel")
	default:
		fmt.Println("Çok Güzel")

	}



}


