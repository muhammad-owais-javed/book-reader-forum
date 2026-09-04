package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"forum/cmd/web/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.LoginHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,  // no default timeout for super slow requests that
		WriteTimeout: 10 * time.Second, // could hold connections open indefinitely
	}

	err = server.ListenAndServe()
	log.Fatal(err)

}
