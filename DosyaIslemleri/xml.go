package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Takım struct {
	Oyuncular []Oyuncu `xml:"oyuncu"`
}
type Oyuncu struct {
	Isim   string `xml:"isim"`
	Mevki  string `xml:"mevki"`
	Gol    int    `xml:"gol"`
	Sosyal struct {
		Facebook string `xml:"facebook"`
		Twitter  string `xml:"twitter"`
		Youtube  string `xml:"youtube"`
	} `xml:"sosyal"`
}
type JsonOyuncu struct {
	Isim   string `json:"isim"`
	Mevki  string `json:"mevki"`
	Gol    int    `json:"gol"`
	Sosyal struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
		Youtube  string `json:"youtube"`
	} `json:"sosyal"`
}

func main() {
	var takim Takım
	file, err := ioutil.ReadFile("DosyaIslemleri/veri.xml")
	checkErr(err)
	err = xml.Unmarshal(file, &takim)
	checkErr(err)
	for i := 0; i < len(takim.Oyuncular); i++ {
		for j := i + 1; j < len(takim.Oyuncular); j++ {
			if takim.Oyuncular[j].Gol > takim.Oyuncular[i].Gol {
				gecici := takim.Oyuncular[i]
				takim.Oyuncular[i] = takim.Oyuncular[j]
				takim.Oyuncular[j] = gecici
			}
		}
	}
	takim.Oyuncular = takim.Oyuncular[0:11]
	newfile, err := os.Create("DosyaIslemleri/atakimi.xml")
	checkErr(err)
	bytearr, err := xml.Marshal(takim)
	checkErr(err)
	_, err = newfile.Write(bytearr)
	checkErr(err)

	var JsonOyunculari []JsonOyuncu
	var JsonOyuncusu JsonOyuncu
	for i := 0; i < len(takim.Oyuncular); i++ {
		JsonOyuncusu.Isim = takim.Oyuncular[i].Isim
		JsonOyuncusu.Mevki = takim.Oyuncular[i].Mevki
		JsonOyuncusu.Gol = takim.Oyuncular[i].Gol
		JsonOyuncusu.Sosyal.Facebook = takim.Oyuncular[i].Sosyal.Facebook
		JsonOyuncusu.Sosyal.Youtube = takim.Oyuncular[i].Sosyal.Youtube
		JsonOyuncusu.Sosyal.Twitter = takim.Oyuncular[i].Sosyal.Twitter
		JsonOyunculari = append(JsonOyunculari, JsonOyuncusu)

	}
	jsonfile, err := os.Create("DosyaIslemleri/atakimi.json")
	checkErr(err)
	jsonbyte, err := json.Marshal(JsonOyunculari)
	checkErr(err)
	_, err = jsonfile.Write(jsonbyte)
	checkErr(err)
}
