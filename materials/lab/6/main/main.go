package main

import (
	"net/http"
	"scrape/logging"
	"scrape/scrape"

	"github.com/gorilla/mux"
)

var LOG_LEVEL = 0

// DONE
//TODO_1: Logging right now just happens, create a global constant integer LOG_LEVEL
//TODO_1: When LOG_LEVEL = 0 DO NOT LOG anything
//TODO_1: When LOG_LEVEL = 1 LOG API details only
//TODO_1: When LOG_LEVEL = 2 LOG API details and file matches (e.g., everything)

func main() {
	logging.IfLevel("starting API server", 1)
	//create a new router
	router := mux.NewRouter()
	logging.IfLevel("creating routes", 1)
	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")
	router.HandleFunc("/addsearch/{regex}", scrape.AddReg).Methods("GET")
	router.HandleFunc("/clear", scrape.ClearFilesAndReg).Methods("GET")
	router.HandleFunc("/reset", scrape.ResetFilesAndReg).Methods("GET")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}
