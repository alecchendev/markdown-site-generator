package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var srcDir string
var buildDir = "build"
var staticDir = "static"
var contentDir = "content"

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
	cmd := exec.Command("cp", "-r", strings.Join([]string{srcDir, staticDir}, "/"), strings.Join([]string{buildDir, staticDir}, "/"))
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through .md files in content

	// createEmptyFile(strings.Join([]string{buildDir, filename}, "/"))

}
