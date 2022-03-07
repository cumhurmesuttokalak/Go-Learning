package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"-"`
	Likes   []Like `json:"like,omitempty"`
	// Likes     []struct {
	// 	Date time.Time
	// }
}

type Like struct {
	ID   string
	Date time.Time
}

func main() {
	u := User{
		Name:    "Cumhur Mesut",
		Surname: "Tokalak",
		Likes: []Like{
			{
				Date: time.Now(),
			},
			{
				Date: time.Now(),
			},
			{
				Date: time.Now(),
			},
			{
				Date: time.Now(),
			},
		},
	}

	arr, _ := json.Marshal(u)
	fmt.Println(string(arr))

	fmt.Println(u)
}
