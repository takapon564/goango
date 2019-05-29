package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	var plainText string
	var decodeStringAsPlain string
	scanner := bufio.NewScanner(os.Stdin)
	// encoder
	fmt.Println("base64 encoder/decoder tool")
	fmt.Print("enter plain text:")
	for scanner.Scan() {
		plainText = scanner.Text()
		if len(plainText) == 0 {
			fmt.Println("Empty character is not allowed. Enter character.")
			fmt.Print("enter plain text:")
			continue
		} else {
			break
		}
	}
	dataForEncode := []byte(plainText)
	encodeString := base64.StdEncoding.EncodeToString(dataForEncode)
	fmt.Println("result: ", encodeString)

	// decoder
	fmt.Print("enter decode string: ")
	for scanner.Scan() {
		decodeStringAsPlain = scanner.Text()
		if len(decodeStringAsPlain) == 0 {
			fmt.Println("Empty character is not allowed. Enter character.")
			fmt.Print("enter decode string: ")
			continue
		} else {
			break
		}
	}
	decodeStringAsByte, err := base64.StdEncoding.DecodeString(decodeStringAsPlain)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("reslut: ", string(decodeStringAsByte))
}
