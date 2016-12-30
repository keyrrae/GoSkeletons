package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {

	output, err := json.MarshalIndent(&albums, "", "\t")
	if err != nil {
		fmt.Println("Converting json to output error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
