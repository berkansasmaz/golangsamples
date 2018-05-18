package main

import "fmt"

func main() {
	kanal1:=make(chan string)//kanal1 isminde kanal oluşturduk
	go func() {kanal1 <- "Merhaba"}()//kanal1 kanalına merhaba metnini koyduk.Veri aktarma işlemi için solo ok işareti (<-)kullandık.Bu satırdı isimsiz bir fonksiyon kullandık.kanal üzerinde haberleşme yeteneği sadece gorutinlerde vardı.
	mesaj:=<-kanal1//kanal1 deki veri mesaj isimli değişkene aktarılıyor.
	fmt.Println(mesaj)

}
