package main

import "fmt"

func main() {
	//Özel fonksiyonlardır.
		murat:=insan{boy: 172}
		fmt.Println(murat.idealkilo())
}

type insan struct {
	boy float32
}

func (birisi insan) idealkilo() float32 {//Metotun şekil açıdan fonksiyondan tek farkı; func sözcüğü ile metotun ismi arasında parantez içinde yazılan kısımdır. Bu kısım,fonksiyonu bir yapıya bağlayarak metot haline getirir. (birisi insan) diye bir kısım olmasaydı normal fonksiyon olmuş olucaktı.o halde (birisi insan) kısmı, ideal kilo() fonksiyonun bağımsız bir fonksiyon olması yerine, bunu insan yapısına bağlayan ve bu fonksiyonu metot haline getiren kısımdır.Bu kısma, metodun alıcsı(receiver) denir.
	kilo:=0.9*birisi.boy-85
	return kilo
}



