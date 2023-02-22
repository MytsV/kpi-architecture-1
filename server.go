package main

import (
	"fmt"
	"net/http"
)

func HandleTime(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "I'm not ready yet >.<")
}

func main() {
	http.HandleFunc("/time", HandleTime)

	fmt.Println("Starting server at port 8795")
	http.ListenAndServe(":8795", nil)
}
