package main

//https://itnext.io/webassemply-with-golang-by-scratch-e05ec5230558
import (
	"log"
	"net/http"
)

const (
	AddSrv       = ":8080"
	TemplatesDir = "."
)

func main() {
	log.Printf("listening on %q...", AddSrv)
	fileSrv := http.FileServer(http.Dir(TemplatesDir))
	if err := http.ListenAndServe(AddSrv, fileSrv); err != nil {
		log.Fatal(err)
	}
}
