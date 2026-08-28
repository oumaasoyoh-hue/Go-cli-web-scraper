package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL")
		return
	}

	url := os.Args[1]

	fmt.Println("URL:", url)

}
