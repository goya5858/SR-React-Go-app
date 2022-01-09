package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func proccessingImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image submit")

	//var header map[string]string = map[string]string{
	//	"Access-Control-Allow-Headers": "Content-Type",
	//	"Access-Control-Allow-Origin":  "*",
	//	"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
	//	"Content-Type":                 "image/*",
	//}

	// Convert r.Body to []byte
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	// Convert []byte to Struct
	var item ItemParams
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err)
	}
	ReqToImg(&item, "/tmp/sample.png")
}
