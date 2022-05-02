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
	mux.HandleFunc("/api/touch", middleware.Middleware(api.TouchHandler))
	mux.HandleFunc("/api/mkdir", middleware.Middleware(api.MkdirHandler))
	mux.HandleFunc("/api/cat", middleware.Middleware(api.CatHandler))
	mux.HandleFunc("/api/edit", middleware.Middleware(api.EditHandler))
	mux.HandleFunc("/api/mv", middleware.Middleware(api.MvHandler))
	mux.HandleFunc("/api/rm", middleware.Middleware(api.RmHandler))
	mux.HandleFunc("/api/download", middleware.Middleware(api.DownloadHandler))
	mux.HandleFunc("/api/upload", middleware.Middleware(api.UploadHandler))

	log.Printf("Starting serve %s at %s", common.RootDir, host)

	if common.FlagHTTPSCert != "" && common.FlagHTTPSKey != "" {
		log.Fatal(http.ListenAndServeTLS(host, common.FlagHTTPSCert, common.FlagHTTPSKey, mux))
	} else {
		log.Fatal(http.ListenAndServe(host, mux))
	}
}
