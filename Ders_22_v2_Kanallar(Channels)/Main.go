package main

import "fmt"
	//Kanal üzerinden canlı iletişim sağlayabilmek için, kanalın iki tarafında aynı anda gorutinler hazır olmalıdır
func main() {
	//Kanal adı dedikodu
	//dedikodu kanalına veriyi yerleştiren fonksiyonun ismi müzevir
	//dedikodu kanalındaki veriyi aktardığımız değikeninde Mahalle Gastesi olsun.
	dedikodu:=make(chan string)
	go müzevir(dedikodu)//bu satırda gorutin kullanmassak fatal error alırız.
	MahalleGastesi:= <- dedikodu
	fmt.Println(MahalleGastesi)
}

func müzevir(dedikodu chan string) {

	dedikodu<-"Leyle Mecnun'u seviyormuş"
}