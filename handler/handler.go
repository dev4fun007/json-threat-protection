package handler

import (
	"log"
	"net/http"
)

const(
	port = ":8090"
)

func main() {

	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/json", jsonThreatDetection)
	log.Fatal(http.ListenAndServe(port, nil))

}


func configHandler(writer http.ResponseWriter, request *http.Request) {

}

func jsonThreatDetection(writer http.ResponseWriter, request *http.Request) {

}