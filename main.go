package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var workDir string
var (
	version   = ""
	goVersion = ""
	buildTime = ""
	commitID  = ""
)

func init() {
	initFlag()
}

func main() {
	flag.Parse()

	if flagVersion {
		fmt.Println("Version:", version)
		fmt.Println("Go Version:", goVersion)
		fmt.Println("Build Time:", buildTime)
		fmt.Println("Git Commit ID:", commitID)
		return
	}

	initAuth()

	var err error
	workDir, err = getWorkDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	host := flagHost + ":" + flagPort

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(basicAuth(mainHandler)))

	log.Print("Starting serve ", workDir, " at ", host, ".")

	if flagHTTPSCert != "" && flagHTTPSKey != "" {
		log.Fatal(http.ListenAndServeTLS(host, flagHTTPSCert, flagHTTPSKey, mux))
	} else {
		log.Fatal(http.ListenAndServe(host, mux))
	}
}
