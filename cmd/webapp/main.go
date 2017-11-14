package main

import (
	"net/http"
	"github.com/tomme87/IMT2681-assignment3/dialogflow"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/dialogflow", dialogflow.HandleDialogflowRequest)
	http.ListenAndServe(":"+port, nil)
}


