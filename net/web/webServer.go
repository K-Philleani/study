package web

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func WebServer() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
