package main

import "fmt"

//DEFER
// defer ertelemek demektir.Yani bir func yazıyoruz onu en sona ertele anlamında.
//eger birden fazla defer varsa ilk yazılan defer en son çağırılır mantıgıyla calısır.
//neden defer kullanılır peki;1-Hataları yönetmek icin,2-Veritabanı bağlantısında kolaylık sağladıgı icin ...
//defer funclar program akışı esnasında erteleme yaptıktan sonra hata ile karşılaşsa da en sonda yine çalışırlar.

func main() {
	fmt.Println("Defer dersi")

	func() {
		fmt.Println("1.func")
	}()
	defer func() {
		fmt.Println("1. defer func")
	}()
	a, b := 2, 0
	c := a / b
	fmt.Println(c)
	defer func() {
		fmt.Println("2. defer func")
	}()

	func() {
		fmt.Println("2.func")
	}()
	fmt.Println("Son print")

}
