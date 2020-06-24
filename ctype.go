package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	cType := resp.Header.Get("Content-Type")
	if cType == "" {
		return "", fmt.Errorf("cant find the Content-Type Header")
	}
	return cType, nil
}
func main() {
	ctype, err := contentType("https://www.google.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
