package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	//Veritabani başlat
	a.DB, _ = gorm.Open("sqlite3", "test.db")
	a.DB.AutoMigrate(&Todo{})

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/todos", a.Todos).Methods("GET")
	a.Router.HandleFunc("/todos", a.TodoCreate).Methods("POST")
	a.Router.HandleFunc("/todos/{id}", a.Todo).Methods("PUT")

}

func (a *App) Run(addr string) {
	fmt.Println("Listening server")
	http.ListenAndServe(addr, a.Router)

}
func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
func (a *App) Todos(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Cumhur "},
		Todo{Name: "Mesut "},
	}
	responseWithJSON(w, http.StatusOK, todos)

}
func (a *App) TodoCreate(w http.ResponseWriter, r *http.Request) {

	var todo Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		responseWithError(w, http.StatusOK, err.Error())
	}
	a.DB.Create(&todo)

	responseWithJSON(w, http.StatusOK, todo)

}
func (a *App) Todo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "This is Todo func %s", vars["id"])
}
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(code)
	w.Write(response)
}
func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"error": message})
}
