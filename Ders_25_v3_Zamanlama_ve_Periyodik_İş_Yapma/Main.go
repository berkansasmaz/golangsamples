package main

import (
	"time"
	"fmt"
)

func main() {
	//Periyodik olarak iş yapmak için ticker isimli işlev kullanalım.
	//ticker işlevi, belirli aralıklarla zaman çizelgesi oluşturur. Bu şekilde oluşturulan bir zaman çizelgesine range işlevi uygulanarak belirli aralıklarla iş yapılması sağlanabilir.
	//1 saniyelik zaman çizelgesi oluşturalım.
	peryot:=time.NewTicker(time.Second)//bir saniyelik peryotlardan oluşan zaman çizelgesi oluşturuldu.
		go func() {
		fmt.Println("Saniyede 1 tane iş yapabilirim:")
		for zaman:=range peryot.C{//range peryot.C ifadesi ile belirtilen peryotlardan oluşan bir dizi oluşturulmuştur ve bu dizi for döngüsüne sokulmuştur. Döngünü sonunda 4 saniyelik uyuma süresinden sonra peryot.Stop() ifadesi ile zaman çizelgesi sonlandırılmıştır.
			fmt.Println(zaman)
		}
	}()
	time.Sleep(time.Second*4)
	peryot.Stop()
}
