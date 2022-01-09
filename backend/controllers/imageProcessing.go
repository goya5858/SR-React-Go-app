package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func proccessingImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image submit")

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

	// Convert Struct to Image & save Image
	byte_data := ReqToImg(&item, "/tmp/sample.png")

	base64img := base64.StdEncoding.EncodeToString(byte_data)

	fmt.Fprintf(w, base64img)
}

var header map[string]string = map[string]string{
	"Access-Control-Allow-Headers": "Content-Type",
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
	"Content-Type":                 "image/*",
}

type Response struct {
	StatusCode      int
	Headers         map[string]string
	Body            string
	IsBase64Encoded bool
}
