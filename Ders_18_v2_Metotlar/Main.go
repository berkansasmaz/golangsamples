package main

import "fmt"

func main() {
	kare:=dörtgen{5,5,10,10}
	fmt.Println("Kare",kare.alan(),"birim kare")
	dikdörtgen:=dörtgen{0,0,50,10,}
	fmt.Println("Dikdörtgen", dikdörtgen.alan(),"birim kare")
}

type dörtgen struct {
	x1,y1,x2,y2 float64
}

func (D *dörtgen) alan() float64  {//metot oluşturduk (D *dörtgen) de metotdun kaynağı metotun ismide aladır
	kenar1:=D.x2-D.x1//alanı bulmak için yaptık
	kenar2:=D.y2-D.y1//alanı bulma için koordinat sistemini düşün
	return kenar1*kenar2
}