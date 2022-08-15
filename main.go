package main

import (
	"fmt"
	"os"

	"github.com/mr-tron/base58"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: base58 <data>")
		return
	}

	data := os.Args[1]

	if os.Args[1] == "--decode" {
		data = os.Args[2]
		decoded, err := base58.Decode(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(decoded))
		return
	}

	encoded := base58.Encode([]byte(data))
	fmt.Println(encoded)
}
