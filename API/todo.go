package main

import "time"

type Todo struct {
	Name      string
	Complated bool
	Due       time.Time
}

type Todos []Todo
