package main

import "fmt"

//Faktöriyel uygulaması
func main() {
	fmt.Println(faktöriyel(4))

}

func faktöriyel(sayi int) int {

	if sayi == 0 {//Eğer değer 0 ise 1 yaz çünkü 0 ın faktöriyeli 1 dir.
			return 1
	}

	return sayi*faktöriyel(sayi-1)
}
//Aynı uygulama for ile de yapılabilir
//func faktöriyel(sayi int) int {
//	çarpım:=1
//	for i:1; i<=sayi ;i++ {
//		çarpım=çarpım*i
//	}
//	return çarpım
//}
