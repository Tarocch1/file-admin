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

	var err error
	workDir, err = getWorkDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	host := flagHost + ":" + flagPort

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(mainHandler))

	log.Print("Start serve ", workDir, " at ", host, ".")
	log.Fatal(http.ListenAndServe(host, mux))
}
