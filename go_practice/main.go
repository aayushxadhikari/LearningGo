package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w,"Hello, world!")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Wrote %d bytes", n))
	})
}
