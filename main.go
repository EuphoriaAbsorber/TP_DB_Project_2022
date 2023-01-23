package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	conf "dbproject/config"
)

func main() {
	myRouter := mux.NewRouter()

	err := http.ListenAndServe(conf.Port, myRouter)
	if err != nil {
		log.Println("cant serve", err)
	}
}
