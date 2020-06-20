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
	c := content{Description: readFile(data)}
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

func isTxtFile(filename string) bool {
	if strings.Contains(filename, ".") {
		return strings.Split(filename, ".")[1] == "txt"
	} else {
		return false
	}
}

func main() {
	filePtr := flag.String("file", "", "txt file to be converted to html file")
	dirPtr := flag.String("dir", "", "directory to search for txt files")
	flag.Parse()
	if *dirPtr != "" {
		files, err := ioutil.ReadDir(*dirPtr)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			name := f.Name()
			if isTxtFile(name) == true {
				renderTemplate("template.tmpl", readFile(name))
				writeTemplateToFile("template.tmpl", name)
			}		
		}
	}
	
	if *filePtr != "" {
		renderTemplate("template.tmpl", readFile(*filePtr))
		writeTemplateToFile("template.tmpl", *filePtr)
	}
	// } else {
	// 	renderTemplate("template.tmpl", readFile("test.txt"))
	// 	writeTemplateToFile("template.tmpl", "test.txt")
	// }
}
