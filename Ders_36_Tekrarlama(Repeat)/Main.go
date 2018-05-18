package main

import (
	"fmt"
	"strings"
)

func main() {
	//Bir ifadeyi istenen sayıda tekrarlama içi kullanılır.

	for i:=1 ; i<=5;i++{
		fmt.Println(strings.Repeat("*",i))
	}
}
