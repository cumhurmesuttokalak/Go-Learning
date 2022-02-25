package main

import (
	"errors"
	"fmt"
)

//HATA DÖNEN FUNCLAR
//Neden kullanılır;Veri tabanına bağlanmak icin,dosya okumak istediğimiz... kullanılır
//Compiletime error  =Program daha akışa başlamadan onceki error. //Derleme Zamanında oluşan hata
//Runtime error. =Programın akış esnasında olan error.          //Calısma Zamanında oluşan hata

func emeklilikHesapla(yas, calisma int) (string, error) {
	if yas < 1 || calisma < 1 || calisma > yas {
		return " ", errors.New("Matematiksel bir hata!")
	} else if yas > 65 || calisma > 25 {
		return "Tebrikler emeklisiniz", nil
	} else {
		return "Emekli olmak icin 25 yıl calışmalısınız ve 65 yaşından büyük olmalısınız", nil
	}
}

func main() {
	var yas, calisma int
	fmt.Println("Lütfen yaşınızı giriniz")
	fmt.Scan(&yas)
	fmt.Println("Lütfen çalışma saatinizi yıl olarak giriniz")
	fmt.Scan(&calisma)
	str, err := emeklilikHesapla(yas, calisma)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(str)

}
