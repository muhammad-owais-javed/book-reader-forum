package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"forum/cmd/web/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := database.InitDB("./app.db", "./schema.sql")
	if err != nil {
		log.Fatalf(">> ERROR: Database initialization failed: %v", err)
	}
	defer db.Close()

	log.Println(">> INFO: Connection to Database Established!")

	//
	// ----- just to temporarily check what is inside the db -----
	//rows, _ := db.Query("SELECT * FROM users") //--
	//defer rows.Close()                         //--
	//for rows.Next() {                          //--
	//	var id, username, email, passwordHash string           //--
	//	err = rows.Scan(&id, &username, &email, &passwordHash) //--
	//	fmt.Println(id, username, email, passwordHash)         //--
	//} //--
	// ----- just to temporarily check what is inside the db -----
	//

	app := &handlers.Application{DB: db}

	mux := app.Routes()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,  // there is no default timeout for super slow requests or
		WriteTimeout: 10 * time.Second, // super slow reading of responses that could hold connections open indefinitely, which is why these are needed
	}

	log.Printf("server running at http://localhost%s", server.Addr)
	err = server.ListenAndServe()
	log.Fatal(err)

}
