package main

import "fmt"

func Üreteç(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filtre(içine, çıkış chan int, asalSayı int) {
	for {
		i := <-içine
		if i%asalSayı != 0 {
			çıkış <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go Üreteç(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filtre(ch, ch1, prime)
		ch = ch1
	}
}
