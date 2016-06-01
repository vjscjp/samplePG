package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	_, err = db.Query("SELECT name FROM users WHERE age = $1", age)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/data", getdata).Methods("GET")
	router.HandleFunc("/data", postdata).Methods("POST")
	port := ":8888"
	log.Printf("Listening at %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

type Response struct {
	Success bool
	Message string
	Data    map[string]interface{}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to postgres sample, Datapoints available: \n /data[GET] \n /data[POST]")
}
func getdata(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to get request")
	log.Println(r.UserAgent())

	resp := new(Response)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "GET:", resp)
}

func postdata(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to post request")
	log.Println(r.UserAgent())

	resp := new(Response)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "POST:", resp)
}
