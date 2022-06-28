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
	router.HandleFunc("/name/{PARAM}", myParamHandler).Methods("GET")
	router.HandleFunc("/bad", myBadHandler).Methods("GET")
	router.HandleFunc("/data", myBodyPostHandler).Methods("POST")
	router.HandleFunc("/headers", myHeadersPostHandler).Methods("POST")
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func myParamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["PARAM"]
	fmt.Fprintf(w, "Hello, %s!\n", param)
}

func myBadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func myBodyPostHandler(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "I got message:\n%s", d)
}

func myHeadersPostHandler(w http.ResponseWriter, r *http.Request) {
	var res int
	aStr := r.Header.Get("a")
	if aStr != "" {
		a, err := strconv.Atoi(aStr)
		if err != nil {
			log.Fatal(err)
		}
		res += a
	}
	bStr := r.Header.Get("b")
	if bStr != "" {
		b, err := strconv.Atoi(bStr)
		if err != nil {
			log.Fatal(err)
		}
		res += b
	}
	w.Header().Add("a+b", strconv.Itoa(res))
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
