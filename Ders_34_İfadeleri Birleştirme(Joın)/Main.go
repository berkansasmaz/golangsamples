package main

import (
	"strings"
	"fmt"
)

func main() {
	//Kesit biçiminde verilen metin ifadelerini birleştirmek için, Joın fonksiyonu kullanılabilir.

	sözcükler:= []string{"süt","Bal","yumurta"}
	fmt.Println(strings.Join(sözcükler,", "))
	fmt.Println(strings.Join(sözcükler,"\n"))

	//Join fonksinu, kendisine ikinci argüman olarak verilen karakterleri ayraç olarak kullanır ve birinci argüman olarak verilen kesitin tüm elemanlarını birleştirir. Yukarıdaki programda ayraç olarak, önce virgül kullanılmış sonra da yeni satır ifadesi kullanılmıştır.

}
