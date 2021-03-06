package controllers

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

// Json or Struct形式になっているRequestから、画像部分を取り出してByte形式で出力する
func ReqToImg(request *ItemParams, filepath string) []byte {
	imgAndTag := request.Text                     // Base64になってる画像部分のみを取り出す
	base64img := strings.Split(imgAndTag, ",")[1] // 頭に "data:image/png;base64," という余計な部分がくっついてるので取り除く
	fmt.Println("DeocdeToBase64Data")

	data, _ := base64.StdEncoding.DecodeString(base64img)
	file, _ := os.Create(filepath)
	defer file.Close()
	file.Write(data) //encode_and_decode.jpgに対して、画像のデータを書き込み
	return data
}
