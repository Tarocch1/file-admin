package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/Tarocch1/file-admin/api"
	"github.com/Tarocch1/file-admin/common"
	"github.com/Tarocch1/file-admin/middleware"
)

func StartServer() {
	host := common.FlagHost + ":" + common.FlagPort
	mux := http.NewServeMux()

	fsRoot, _ := fs.Sub(static, "static")

	mux.HandleFunc("/", middleware.Middleware(http.FileServer(http.FS(fsRoot)).ServeHTTP))

	mux.HandleFunc("/api/ls", middleware.Middleware(api.LsHandler))

	log.Printf("Starting serve %s at %s", common.WorkDir, host)

	if common.FlagHTTPSCert != "" && common.FlagHTTPSKey != "" {
		log.Fatal(http.ListenAndServeTLS(host, common.FlagHTTPSCert, common.FlagHTTPSKey, mux))
	} else {
		log.Fatal(http.ListenAndServe(host, mux))
	}
}
