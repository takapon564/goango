package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("base64 encoder/decoder tool")
	fmt.Print("enter plain text:")
	var plainText string
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
	data := []byte(plainText)
	encodeText := base64.StdEncoding.EncodeToString(data)
	fmt.Println("result: ", encodeText)
}
