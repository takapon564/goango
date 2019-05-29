package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func encoder() {
	var plainText string
	scanner := bufio.NewScanner(os.Stdin)
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
}

func decoder() {
	var decodeStringAsPlain string
	scanner := bufio.NewScanner(os.Stdin)
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
	fmt.Println("result: ", string(decodeStringAsByte))
}

func main() {
	fmt.Println("base64 encoder/decoder tool")
	fmt.Print("Do you want to encode, or decode? Please enter: e/d: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		options := scanner.Text()
		if options == "e" {
			encoder()
			break
		} else if options == "d" {
			decoder()
			break
		} else {
			fmt.Println("Please enter: e/d")
			continue
		}
	}

}
