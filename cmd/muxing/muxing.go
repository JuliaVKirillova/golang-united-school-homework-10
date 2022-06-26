package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, err := w.Write([]byte("Hello, " + vars["PARAM"] + "!"))
		if err != nil {
			return
		}
	}).Methods("GET")

	router.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}).Methods("GET")

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		param, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		defer r.Body.Close()

		_, err = w.Write([]byte("I got message:\n" + string(param)))
		if err != nil {
			return
		}
	}).Methods("POST")

	router.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("A")
		aNum, _ := strconv.Atoi(a)
		b := r.Header.Get("B")
		bNum, _ := strconv.Atoi(b)
		w.Header().Set("a+b", strconv.Itoa(aNum+bNum))
	}).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
