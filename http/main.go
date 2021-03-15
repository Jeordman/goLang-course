package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// create an empty byte slice with space for x elements
	bs := make([]byte, 999999)
	// take the body and read it into the bs
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}
