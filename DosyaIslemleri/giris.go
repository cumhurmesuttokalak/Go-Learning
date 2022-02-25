package main

import (
	"log"
	"os"
)

//log.fatal'ın printten farkı: log.fatal calıştıktan sonra sistemden cıkış yapılır defer funclar da çalışmaz.

//Dosya Oluşturma
//Dosyaya veriYazma
//Dosya okuma
//Dosyanın adını ve yolunu değiştirme
//Dosyanın özelliklerini oluşturma
//Dosyayı silme

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*
		file, err := os.Create("deneme.txt")
		check(err)

		_, err = file.WriteString("Merhaba")
		check(err)
	*/
	/*
		content, err := ioutil.ReadFile("deneme.txt")
		check(err)
		fmt.Println(string(content))
	*/
	/*
		err := os.Mkdir("yeniklasör", 0777)
		check(err)
		err = os.Rename("yeni.txt", "yeniklasör/yenitxt")
		check(err)
	*/
	/*
		file, err := os.Stat("yeniklasör/yenitxt")
		check(err)
		fmt.Println(file.Mode())
	*/
	err := os.RemoveAll("yeniklasör")
	check(err)
}
