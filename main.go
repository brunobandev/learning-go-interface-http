package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// io.Copy(os.Stdout, resp.Body)

	// My own implementation
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// creating my own function that implements the interface (io.Writer)
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bites", len(bs))

	return len(bs), nil
}
