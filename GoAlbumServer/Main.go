package main

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var albums []Album

var db *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "albums.db")

	if err != nil {
		panic(err)
		return
	}
	db.Find(&albums)
	//fmt.Println(albums[0])
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", handleRequest)
	server.ListenAndServe()
}
