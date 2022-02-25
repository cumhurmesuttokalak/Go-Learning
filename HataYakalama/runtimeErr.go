package main

import "fmt"

//run time hata da hatanın oldugu satırdan sonra program çalışmaz
//ve direkt hatayı ekrana verir.Fakat defer funclar hata olsa bile çalışabildikleri icin yakalamada işimize yaralar
// recover fonksiyonu defer fonksiyona gelmeden önce (programda) bir runtimeErr oluşmusşa o hatayı yakalayıp değişkene atar
func main() {
	fmt.Println("runtime Err")

	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("Hiç bir sayı 0 a bölünemez", error)
		}

	}()
	a, b := 2, 0
	c := a / b
	fmt.Println(c)
}
