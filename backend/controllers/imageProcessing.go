package controllers

import (
	"fmt"
	"net/http"
)

func proccessingImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image submit")
}
