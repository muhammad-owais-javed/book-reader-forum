package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/logic", LoginHandler)

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	err = server.ListenAndServe()
	log.Fatal(err)

}
