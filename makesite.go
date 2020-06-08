package main

import (
	"io/ioutil"
	"fmt"
)

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate() {
	
}

func main() {
	fmt.Println(readFile())
}
