package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// using real client
	client := &http.Client{Timeout: time.Second}
	parser := NewParser(client)

	// get full contents of example.com html
	html, err := parser.GetContents("http://example.com/", 11)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(html)
}
