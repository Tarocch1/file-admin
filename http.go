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
		postHandler(w, r, targetPath)
		break
	case http.MethodPut:
		putHandler(w, r, targetPath)
		break
	case http.MethodDelete:
		deleteHandler(w, r, targetPath)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
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
		allFiles, err := ioutil.ReadDir(targetPath)
		if err != nil {
			errorHandler(w, err)
			return
		}
		var data listTemplateDataStruct
		data.Title = targetPath
		data.Files = []listTemplateDataFileStruct{}
		var dirs, files []listTemplateDataFileStruct
		for _, fileInfo := range allFiles {
			if fileInfo.IsDir() {
				dirs = append(dirs, listTemplateDataFileStruct{
					Name:  fileInfo.Name(),
					IsDir: fileInfo.IsDir(),
					Time:  fileInfo.ModTime().Unix(),
					Size:  fileInfo.Size(),
				})
			} else {
				files = append(files, listTemplateDataFileStruct{
					Name:  fileInfo.Name(),
					IsDir: fileInfo.IsDir(),
					Time:  fileInfo.ModTime().Unix(),
					Size:  fileInfo.Size(),
				})
			}
		}
		data.Files = append(data.Files, dirs...)
		data.Files = append(data.Files, files...)
		t, err := template.New("listTemplate").Funcs(template.FuncMap{"formatSize": formatSize}).Parse(listTemplate)
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

func postHandler(w http.ResponseWriter, r *http.Request, targetPath string) {
	if !pathNotExist(targetPath) {
		w.WriteHeader(http.StatusConflict)
		return
	}
	fileBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, err)
		return
	}
	err = ioutil.WriteFile(targetPath, fileBytes, 0666)
	if err != nil {
		errorHandler(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func putHandler(w http.ResponseWriter, r *http.Request, targetPath string) {
	flag := pathNotExist(targetPath)
	fileBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, err)
		return
	}
	err = ioutil.WriteFile(targetPath, fileBytes, 0666)
	if err != nil {
		errorHandler(w, err)
		return
	}
	if flag {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}
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
