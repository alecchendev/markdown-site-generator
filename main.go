package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/otiai10/copy"
)

var srcDir string
var buildDir = "build"
var staticDir = "static"
var contentDir = "content"
var layoutFile = "layout.html"

func createEmptyFile(filename string) {
	data := []byte("")
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {

	// Get src directory
	if len(os.Args) < 2 {
		srcDir = "."
	} else {
		if len(os.Args) > 2 {
			fmt.Println("Unnecessary extra arguments.")
		}
		srcDir = os.Args[1]
	}

	// Create build directory
	_, err := os.Stat(buildDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(buildDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	// Copy src/static to build/static
	err = copy.Copy(strings.Join([]string{srcDir, staticDir}, "/"), strings.Join([]string{buildDir, staticDir}, "/"))
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through .md files in content
	layout, err := ioutil.ReadFile(strings.Join([]string{srcDir, layoutFile}, "/")) // []uint8
	if err != nil {
		log.Fatal(err)
	}

	c, err := ioutil.ReadDir(strings.Join([]string{srcDir, contentDir}, "/"))
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
		if entry.IsDir() {
			continue
		}
		// - Read file
		fileBytes, err := ioutil.ReadFile(strings.Join([]string{srcDir, contentDir, entry.Name()}, "/")) // []uint8
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T\n", fileBytes)
		fmt.Println(string(fileBytes))

		// - Generate html corresponding to the markdown
		extensions := parser.CommonExtensions | parser.Attributes
		parser := parser.NewWithExtensions(extensions)
		html := markdown.ToHTML(fileBytes, parser, nil) // []uint8
		fmt.Println(string(html))

		// - Create html file from layout.html
		filename := strings.Join([]string{buildDir, strings.Replace(entry.Name(), ".md", ".html", 1)}, "/")
		layoutStr := string(layout)
		// Check if div content is even there
		output := strings.Replace(layoutStr, "<div id=\"content\"></div>", string(html), 1)
		err = ioutil.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			log.Fatal(err)
		}

		// - Fill "content" div with content
	}

	// createEmptyFile(strings.Join([]string{buildDir, filename}, "/"))

}
