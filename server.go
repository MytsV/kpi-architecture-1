package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Payload struct {
	Time string `json:"time"`
}

func HandleTime(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method is not supported", http.StatusNotFound)
		return
	}
	res.Header().Set("Content-Type", "application/json")

	payload := Payload{Time: time.Now().Format(time.RFC3339)}
	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(res, string(data))
}

const port = 8795

func main() {
	http.HandleFunc("/time", HandleTime)

	fmt.Printf("Starting server at port %d\n", port)
	portFmt := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(portFmt, nil); err != nil {
		panic(err)
	}
}
