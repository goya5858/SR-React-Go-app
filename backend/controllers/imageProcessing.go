package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"backend/onnxReference"
)

func proccessingImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("image submit")
	inpath := "/tmp/sample.png" //提出されたファイルを保存するPATH
	outpath := "/tmp/output.png"

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
	ReqToImg(&item, inpath)

	// Reference Img and Save OutputImg
	onnxReference.OnnxRef(inpath, outpath)

	// Read OutImgFile and convert to []byte
	file, err := os.Open(outpath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, _ := file.Stat()
	byte_data := make([]byte, fileInfo.Size())
	file.Read(byte_data)

	// Convert Byte_Image to Base64String
	base64img := base64.StdEncoding.EncodeToString(byte_data)

	fmt.Fprintf(w, base64img)
}

//var header map[string]string = map[string]string{
//	"Access-Control-Allow-Headers": "Content-Type",
//	"Access-Control-Allow-Origin":  "*",
//	"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
//	"Content-Type":                 "image/*",
//}
//
//type Response struct {
//	StatusCode      int
//	Headers         map[string]string
//	Body            string
//	IsBase64Encoded bool
//}
