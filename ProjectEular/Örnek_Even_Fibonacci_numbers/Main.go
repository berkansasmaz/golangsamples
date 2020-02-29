package main

import "fmt"



var max int=1
var min int=0
var toplam int
func main() {
	for a := 1; a <=35; a++ {

		var fib int

		fib=min+max
		min=max
		max=fib

		if fib%2 == 0 && fib<=4000000 {
			toplam+=fib
		}



	}
	fmt.Println(toplam)


}

