package main

import (
	"net/http"
	"github.com/tomme87/IMT2681-assignment3/dialogflow"
)

func main() {

	http.HandleFunc("/dialogflow", dialogflow.HandleDialogflowRequest)
	http.ListenAndServe("localhost:8080", nil)
}


