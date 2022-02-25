package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Eleman struct {
	Isim string `json:"isim"`
	Maas int    `json:"maas"`
	Puan int    `json:"puan"`
}

func main() {
	file, err := ioutil.ReadFile("DosyaIslemleri/veri.json")
	checkerr(err)
	var elemanlar []Eleman
	err = json.Unmarshal(file, &elemanlar)
	checkerr(err)
	var eniyi Eleman = elemanlar[0]

	for i := 0; i < len(elemanlar); i++ {
		if elemanlar[i].Puan > eniyi.Puan {
			eniyi = elemanlar[i]
		}
	}
	for i := 0; i < len(elemanlar); i++ {
		if eniyi == elemanlar[i] {
			elemanlar[i].Maas += elemanlar[i].Maas * 10 / 100
		}
	}
	newfile, err := json.Marshal(elemanlar)
	checkerr(err)
	err = ioutil.WriteFile("DosyaIslemleri/veri.json", newfile, 0666)
	checkerr(err)
	fmt.Println(elemanlar)

}
