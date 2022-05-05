package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Simple API creation example

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func SimpleAPIExample() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8086", router))
}
