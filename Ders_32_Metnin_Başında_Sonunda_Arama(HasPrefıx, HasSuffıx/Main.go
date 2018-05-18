package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.HasPrefix("Ankara","a"))
	fmt.Println(strings.HasSuffix("Ankara","a"))

	//Fonksiyon büyük ve küçük harf duyarlı çalışmaktadır. Bu nedenle, A ve a harfleri farklı karakterlerdir.
}
