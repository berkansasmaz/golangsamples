package main

import "fmt"

func main() {//Hata soruda int değer kullanılması int kullanıldığı içinde değer 2 olarak gözüküyor.
	var a,b float32= 5,2
	fmt.Println("Bölme",bölme(a,b))
}

func bölme(a, b float32) float32 {
	return a/b
}