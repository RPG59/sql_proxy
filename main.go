package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func getDump(w http.ResponseWriter, r *http.Request) {
    connStr := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	    w.WriteHeader(http.StatusInternalServerError)
	}

	// age := 21
	db.Query("CREATE DATABASE users")
	// row, err := db.Query("SELECT name FROM users WHERE age = $1", age)

    // if err != nil {
	// 	log.Fatal(err)
	//     w.WriteHeader(http.StatusInternalServerError)
    // }

    // fmt.Println(row)
	w.WriteHeader(http.StatusOK)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/dump", getDump).Methods("POST")

	fmt.Println("Listen server at port 8080")

	http.ListenAndServe(":8080", router)

}
