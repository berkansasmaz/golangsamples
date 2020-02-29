package main

import "fmt"

func main() {
	var toplam int
	for x := 0; x < 1000; x++ {
		if x%3 == 0 || x%5 == 0 {
			toplam+=x
		}
	}
	fmt.Println(toplam)
}
