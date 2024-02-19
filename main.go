package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const connectionString = "postgres://postgres:root@localhost:5433/postgres?sslmode=disable"

func getDump(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("dump")
	service := r.FormValue("service")

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println("Service:" + service)

	buf := bytes.NewBuffer(nil)

	// Copy the contents of the file to the form field
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer file.Close()

	dump := buf.String()

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(dump)

	row, err := db.Exec(dump)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(row)

	w.WriteHeader(http.StatusOK)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/dump", getDump).Methods("POST")

	fmt.Println("Listen server at port 8080")

	http.ListenAndServe(":8080", router)

}
