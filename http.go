package main

import (
	"html/template"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	targetPath, err := getTargetPath(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	switch r.Method {
	case http.MethodGet:
		getHandler(w, r, targetPath)
		break
	case http.MethodPost:
		break
	case http.MethodPut:
		break
	case http.MethodDelete:
		deleteHandler(w, r, targetPath)
		break
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request, targetPath string) {
	if pathNotExist(targetPath) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	isDir, err := pathIsDir(targetPath)
	if err != nil {
		errorHandler(w, err)
		return
	}
	if isDir {
		files, err := ioutil.ReadDir(targetPath)
		if err != nil {
			errorHandler(w, err)
			return
		}
		var data listTemplateDataStruct
		data.Title = targetPath
		data.Files = []string{}
		for _, fileInfo := range files {
			data.Files = append(data.Files, fileInfo.Name())
		}
		t, err := template.New("listTemplate").Parse(listTemplate)
		if err != nil {
			errorHandler(w, err)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		t.Execute(w, data)
		return
	}
	mimeType := mime.TypeByExtension(filepath.Ext(targetPath))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	fileBytes, err := ioutil.ReadFile(targetPath)
	if err != nil {
		errorHandler(w, err)
		return
	}
	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, targetPath string) {
	if pathNotExist(targetPath) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := os.RemoveAll(targetPath)
	if err != nil {
		errorHandler(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func errorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
