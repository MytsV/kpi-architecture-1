package main

import (
	"fmt"
	"net/http"
)

const port = 8795

func main() {
	http.HandleFunc("/time", HandleTime)

	fmt.Printf("Starting server at port %d\n", port)
	portFmt := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(portFmt, nil); err != nil {
		panic(err)
	}
}
