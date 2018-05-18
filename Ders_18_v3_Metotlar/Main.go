package main

import "fmt"

func main() {
	//Metotlar yapı türünün haricindede kullanılabilmektedir.
	var a ,b sayı=5, 2
	c:=a-b
	//a,b ve c sayı türünde değişkendir
	fmt.Println(a.karesi(),b.karesi(),c.karesi())

}

type sayı int //sayı int türündedir

func (x sayı) karesi() int{
	return int(x*x)
}


