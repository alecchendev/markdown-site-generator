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
var fillDiv = "<div id=\"content\"></div>"
var routerFile = "router.js"
var scriptTags = `
<script src="static/router.js"></script>
<script src="static/routes.js"></script>
<script src="static/index.js"></script>
`

func main() {

	// Get src directory
	fmt.Println(os.Args)
	ccr := false // client side routing
	if len(os.Args) < 2 {
		srcDir = "./src"
	} else {
		if os.Args[1] == "-r" {
			if len(os.Args) == 2 {
				srcDir = "./src"
			} else {
				srcDir = os.Args[2]
			}
			ccr = true
			if len(os.Args) > 3 {
				fmt.Println("Unnecessary extra arguments")
			}
		} else {
			srcDir = os.Args[1]
			if len(os.Args) > 2 {
				fmt.Println("Unnecessary extra arguments")
			}
		}
	}

	fmt.Printf("Source directory: %v\n", srcDir)

	// Create build directory
	_, err := os.Stat(buildDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(buildDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Created build directory")

	// Copy src/static to build/static
	err = copy.Copy(strings.Join([]string{srcDir, staticDir}, "/"), strings.Join([]string{buildDir, staticDir}, "/"))
	if err != nil {
		log.Fatal(err)
	}
	// err = copy.Copy(routerFile, strings.Join([]string{buildDir, staticDir, routerFile}, "/"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Copied static assets")

	// Read in layout.html
	layout, err := ioutil.ReadFile(strings.Join([]string{srcDir, layoutFile}, "/")) // []uint8
	if err != nil {
		log.Fatal(err)
	}
	layoutStr := string(layout)

	// Add js scripts if client-side routing
	if ccr {
		idx := strings.Index(layoutStr, "</body>")
		layoutStr = layoutStr[:idx] + scriptTags + layoutStr[idx:]
	}

	// Check if div content is even there
	if !strings.Contains(layoutStr, fillDiv) {
		log.Fatalf("layout.html does not contain %v", fillDiv)
	}

	// Read in content directory
	contentFiles, err := ioutil.ReadDir(strings.Join([]string{srcDir, contentDir}, "/"))
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through .md files in content
	for _, entry := range contentFiles {
		if entry.IsDir() {
			continue
		}

		filename := strings.Join([]string{buildDir, strings.Replace(entry.Name(), ".md", ".html", 1)}, "/")
		fmt.Printf("Creating %v...", filename)

		// Read file
		fileBytes, err := ioutil.ReadFile(strings.Join([]string{srcDir, contentDir, entry.Name()}, "/")) // []uint8
		if err != nil {
			log.Fatal(err)
		}

		// Generate html corresponding to the markdown
		extensions := parser.CommonExtensions | parser.Attributes
		parser := parser.NewWithExtensions(extensions)
		html := markdown.ToHTML(fileBytes, parser, nil) // []uint8

		if !ccr {
			// Populate routes.js with markdown content

		} else {
			// Create html file from layout.html - fill with markdown content
			output := strings.Replace(layoutStr, fillDiv, string(html), 1)
			err = ioutil.WriteFile(filename, []byte(output), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Print("Done\n")

	}
	fmt.Print("Build complete\n")

}
