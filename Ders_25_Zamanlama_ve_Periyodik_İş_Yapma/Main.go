package main

import (
	"time"
	"fmt"
)

func main() {
	//Go dilinde, ileriki zamanda çalışmak üzere bir iş zamanlaması yapılabilir. Bu amaçla, timer isimli işlev kullanılır. Bir timer oluşturulurken, onun çalışmadan önce bekleyeceği süre belirtilir. Zamanı geldiğinde ilgili timer bir kanal üzerinden mesaj göndererek, zamanın geldiğini haber verir.
	timer1:=time.NewTimer(time.Second*2)//time.NewTimer yerine time.Sleep' de kullanılabilir.
	<-timer1.C//Bir timer'ın süres, dolduktan sonra, o anki zaman bilgisi timer'ın C metoduna gönderilir.Bu veriyi timer.C şeklinde erişilir.
	fmt.Println("2 saniye geçti")
}
