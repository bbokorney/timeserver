package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bbokorney/goutils"
)

var hostport string

func requestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received from %v", r.RemoteAddr)
	fmt.Fprintf(w, "The time is %s from %s\n", time.Now(), hostport)
}

func main() {
	hostport = goutils.GetRequiredEnv("HOSTPORT")

	log.Println(fmt.Sprintf("Starting server on %s...", hostport))
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(hostport, nil))
}
