package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var plainText string
	fmt.Print("enter plain text:")
	fmt.Scan(&plainText)
	data := []byte(plainText)
	encodeText := base64.StdEncoding.EncodeToString(data)
	fmt.Println("result: ", encodeText)
}
