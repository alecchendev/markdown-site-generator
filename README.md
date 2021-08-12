# Marvon
A simple markdown to static site build tool. Essentially Hugo for minimalists.

### Motivation
I was too lazy to use Hugo or Gatsby. Wanted less complexity.

### Function
Given a src directory with markdown (.md) files, static assets, and a layout.html file holding the element `<div id="content"></div>`, marvon will create a build directory holding html files built from replacing the content div with the content in your provided markdown files. Each html page built corresponds to a provided markdown file. See the src and build directories in this repository for examples.

### How to use this tool to build your site
1. [Install Go](https://golang.org/doc/install)
2. Clone this repository
3. Build the executable by running `go build .` in the root directory
4. Either populate the src directory with your content, or place the executable in a directory following a similar structure, i.e.
```
root/
    src/
        content/
        static/
        layout.html
    marvon(.exe)
```
**Note: To ensure proper build, place .md files in content. place static assets in static, and ensure layout.html includes `<div id="content"></div>`.**

5. Build your site by running `./marvon src` for mac/linux or `marvon.exe src` for windows (src is just an argument providing the name of the directory holding your markdown content, static assets, and layout.html)