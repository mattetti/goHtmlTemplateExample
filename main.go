package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var templates *template.Template
var templateDelims = []string{"{{%", "%}}"}

func init() {
	// initialize the templates
	basePath := "resources/templates/"
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// don't process folders themselves
		if info.IsDir() {
			return nil
		}
		templateName := path[len(basePath):]
		if templates == nil {
			templates = template.New(templateName)
			templates.Delims(templateDelims[0], templateDelims[1])
			_, err = templates.ParseFiles(path)
		} else {
			_, err = templates.New(templateName).ParseFiles(path)
		}
		log.Printf("Processed template %s\n", templateName)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("resources/images"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("resources/javascripts"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("resources/stylesheets"))))
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":1313", nil)
}
