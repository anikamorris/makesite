package main

import (
	"io/ioutil"
	// "fmt"
	"html/template"
	"os"
	"flag"
	"strings"
)

type content struct {
	Description string
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(filename string, data string) {
	c := content{Description: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = t.Execute(os.Stdout , c)
	if err != nil {
		panic(err)
	}
}

func addExtHTML(filename string) string {
	ext := ".html"
	withExt := strings.Split(filename, ".")[0] + ext
	return withExt
}

func writeTemplateToFile(tmplName string, data string) {
	c := content{Description: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(tmplName))

	file := addExtHTML(data)
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, c)
	if err != nil {
		panic(err)
	}
}

func main() {
	filePtr := flag.String("file", "", "txt file to be converted to html file")
	flag.Parse()
	if *filePtr != "" {
		renderTemplate("template.tmpl", readFile(*filePtr))
		writeTemplateToFile("template.tmpl", *filePtr)
	} else {
		renderTemplate("template.tmpl", readFile("test.txt"))
		writeTemplateToFile("template.tmpl", "test.txt")
	}
}
