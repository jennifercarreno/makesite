package main

import (
	"fmt"
	"html/template"
	"os"
	"flag"
	"io/ioutil"
	"path/filepath"
	"github.com/gomarkdown/markdown"

)

type Page struct {
    TextFilePath string
    TextFileName string
    HTMLPagePath string
    Content      template.HTML
}

func main() {
	dirFlag := flag.String("dir", "", "directory path to search for files")
	flag.Parse()

	// Check if the dir flag value is provided
	if *dirFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Get the absolute path of the directory
	dirPath, err := filepath.Abs(*dirFlag)
	if err != nil {
		panic(err)
	}

	// Read the directory to find files
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text files in the directory:")
	for _, fileInfo := range fileInfos {
		//searching for md files
		if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == ".md" {
			fmt.Println(fileInfo.Name())
			
			//gets file path and reads contents
			filePath := filepath.Join(dirPath, fileInfo.Name())
			fileContents, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			//gets the contents of the md file and renders to html
			html := markdown.ToHTML(fileContents, nil, nil)
			content := template.HTML(html)

			page := Page {
			TextFilePath: filePath,
			TextFileName: "test",
			HTMLPagePath: fileInfo.Name() + ".html",
			Content:      content,
			}

			// creates a new file from the template and adds the content from page
			t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
			newFile, err := os.Create(page.HTMLPagePath)
			if err != nil {
				panic(err)
			}
			t.Execute(newFile, page)


		}
	}

}