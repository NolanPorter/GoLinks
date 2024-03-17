package main

import (
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("http://example.com/")

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	println(body)
	println(err)
}
