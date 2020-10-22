package main

import (
	"flag"
	"log"
	"net/http"
)

var workDir string

func init() {
	initFlag()
}

func main() {
	flag.Parse()

	initAuth()

	var err error
	workDir, err = getWorkDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	host := flagHost + ":" + flagPort

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(basicAuth(mainHandler)))

	log.Print("Start serve ", workDir, " at ", host, ".")
	log.Fatal(http.ListenAndServe(host, mux))
}
