package main

import (
	"fmt"
	"strings"
)

func main() {
	//Metni belirli bir ifadeye(ayraç) göre bölmek için Split fonksiyonu kullanılır. Bu fonksionun bir kaç farklı türevi bulunmaktadır.

	fmt.Printf("%q\n", strings.Split("a-b-c-d-e","-"))//kesit parçalama işlem tire(-) kullanıldır
	fmt.Printf("%q\n", strings.SplitAfter("a-b-c-d-e","-"))//SplitAfter fonksiyonu metni ayracın hemen sonrasından ayırır.
	fmt.Printf("%q\n", strings.SplitN("a-b-c-d-e","-",3))//SplitN fonksiyonu metni N tane parçaya ayırır. N sayısı, argüman olarak belitilmelidir. Örnekte 3 olarak belirtildiği için, metni 3 parçaya ayırmıştır. Kalanların tamamı son parçaya dahil edilmiştir.
	fmt.Printf("%q\n", strings.SplitAfterN("a-b-c-d-e","-",3))//SplitAfterN fonksiyonu metni N tane parçaya ayırır ve ayraçları da parçalara dahil eder.
}
