package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"forum/cmd/web/handlers"
	"forum/internal/repositories"
	"forum/internal/services"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	//creating the tables
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}

	// ----- just to temporarily check what is inside the db -----
	//rows, _ := db.Query("SELECT * FROM users") //--
	//defer rows.Close()                         //--
	//for rows.Next() {                          //--
	//	var id, username, email, passwordHash string           //--
	//	err = rows.Scan(&id, &username, &email, &passwordHash) //--
	//	fmt.Println(id, username, email, passwordHash)         //--
	//} //--
	// ----- just to temporarily check what is inside the db -----

	// -- repos --
	userRepository := &repositories.UserRepository{DB: db}
	// -- services --
	registrationService := &services.RegistrationService{Repository: userRepository}
	authService := &services.AuthService{Repository: userRepository}

	app := &handlers.Application{
		Auth:         authService,
		Registration: registrationService,
	}

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
