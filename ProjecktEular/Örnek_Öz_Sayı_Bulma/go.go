package main

import "fmt"

func main() {

	var özsayı int

	for i := 100; i < 1000; i++ {
		sayi := i * 9

		binler := sayi / 1000
		sayi = sayi % 1000
		yüzler := sayi / 100
		sayi = sayi % 100
		onlar := sayi / 10
		birler := sayi % 10

		toplam := binler + yüzler + onlar + birler

		özsayı += toplam

	}
	fmt.println(özsayı)

}
