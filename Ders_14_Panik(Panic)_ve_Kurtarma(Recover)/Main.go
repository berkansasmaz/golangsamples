package main

import (

	"fmt"
)

func main() {
	dizi:=make([]int,5)

	defer func() {//defer func en son hata varsa çalışması için yazdık
		err:=recover()
		if err!=nil {
			fmt.Println(err)
		}
	}()

	dizi[7]=6//burada bilerek hata yaptık burak selim senyurt defer panic recover alıntı

	var dizitop int
	for _,toplam:=range dizi{
		dizitop+=toplam
	}
	fmt.Println(dizitop)

}
