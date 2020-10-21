package main

type listTemplateDataStruct struct {
	Title string
	Files []listTemplateDataFileStruct
}

type listTemplateDataFileStruct struct {
	Name  string
	IsDir bool
	Time  int64
	Size  int64
}
