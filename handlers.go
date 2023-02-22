package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Payload struct {
	Time string `json:"foo"`
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

	fmt.Fprint(res, string(data))
}
