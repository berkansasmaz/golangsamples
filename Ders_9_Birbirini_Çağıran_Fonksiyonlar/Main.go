package main

import (
	"fmt"
)

func main() {
	isim:=isimal()
	yaş:=yaşal()
	fmt.Printf("\nHoşgeldin %s, yaşınız %d\n",isim,yaş)
	kontrolet(yaş)

}

func isimal()(isim string) {
	fmt.Println("İsminiz: ")
	fmt.Scanln(&isim)
	return
}

func yaşal()(yaş int)  {
	fmt.Println("Yaşınız: ")
	fmt.Scanln(&yaş)
	return
}

func kontrolet(yaş int)()  {
	if yaş<18 {
		fmt.Println("Sisteme giriş için yaşını müsait değil. ")
	} else if yaş<120 {
		fmt.Println("Sisteme başarılı olarak giriş yaptınız. ")
		oturumac()
	}else{
		fmt.Println("Şaka yapıyorsunuz değil mi? ")
	}
}

func oturumac()  {
	fmt.Println("İşiniz bitince oturumunuzu kapatmayı unutmayın. ")
}