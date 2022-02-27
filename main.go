package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Misi struct {
	ID   string `json:"id"`
	Misi string `json:"title"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Misi
}

var db *sql.DB
var err error

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/misi", getMisi).Methods("GET")
	http.ListenAndServe(":8000", router)
}
