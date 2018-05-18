package main

import "fmt"

/* 3 */func main() {
		//İnsanların boyuna göre ideal ağırlıklarını hesaplayan fonksiyon geliştirelim.Fonksiyona yapı türünde bir değişke gönderelim.
	nigar:=insan{boy: 165,cinsiyet:"Kadın"}
	murat:=insan{boy:172,cinsiyet:"Erkek"}

	ideal_kilo(&murat)
	ideal_kilo(&nigar)
	fmt.Println(murat.kilo)
	fmt.Println(nigar.kilo)
}

/* 1 */ type insan struct {
	boy, kilo float32
	cinsiyet string
}
/* 2 */func ideal_kilo(birisi *insan)  {//burada ki çarpı işaretci için
	switch birisi.cinsiyet {
	case "Erkek":
		birisi.kilo=0.9*birisi.boy-85
	case "Kadın":
		birisi.kilo=0.9*birisi.boy-90
	default:
		fmt.Println("Cinsiyet alanında sorun var")
	}
}