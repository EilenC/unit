package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chain/src"
)

//init chain
var chain *src.Chain

const host = ":8888"

func main() {
	chain = src.InitChain()
	http.HandleFunc("/chain/get", get)
	http.HandleFunc("/chain/send", send)
	fmt.Printf("Server [%s] Start...\n", host)
	_ = http.ListenAndServe(host, nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	marshal, err := json.Marshal(chain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(marshal)
}

func send(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	if data == "" {
		http.Error(w, "data empty", http.StatusBadRequest)
		return
	}
	chain.SendData(data)
	get(w, r)
}
