package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	//LOG//
	l, err := os.OpenFile("pnc.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkErr(err)

	defer l.Close()

	log.SetOutput(l)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	gorilla := mux.NewRouter()
	gorilla.HandleFunc("/", indexHandler)
	gorilla.HandleFunc("/casper", casperHandler)
	gorilla.HandleFunc("/donate", donateHandler)
	gorilla.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	http.ListenAndServe(":8808", gorilla)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
