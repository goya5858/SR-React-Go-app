package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartBackendServer() error {
	fmt.Println("Start Go Server")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rootPage)
	router.HandleFunc("/submit", proccessingImage)

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome ot the Go API Server")
	fmt.Println("Root is fooked")
}
