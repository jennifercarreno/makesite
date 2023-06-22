package main

import (
	"fmt"
	"html/template"
	"os"
	"flag"
	"io/ioutil"
	"path/filepath"

)

type Page struct {
    TextFilePath string
    TextFileName string
    HTMLPagePath string
    Content      string
}

func main() {
	dirFlag := flag.String("dir", "", "Directory path to search for .txt files")
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

	// Read the directory to find .txt files
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text files in the directory:")
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == ".txt" {
			fmt.Println(fileInfo.Name())

			filePath := filepath.Join(dirPath, fileInfo.Name())
			fileContents, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			page := Page{
			TextFilePath: filePath,
			TextFileName: "test",
			HTMLPagePath: fileInfo.Name() + ".html",
			Content:      string(fileContents),
		}

			t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
			newFile, err := os.Create(page.HTMLPagePath)
			if err != nil {
				panic(err)
			}
			t.Execute(newFile, page)

		}
	}

	// file := flag.String("file", "", "Path to the text file")
	// flag.Parse()

	// fileContents, err := ioutil.ReadFile(*file)
	// if err != nil {
	// 	// A common use of `panic` is to abort if a function returns an error
	// 	// value that we don’t know how to (or want to) handle. This example
	// 	// panics if we get an unexpected error when creating a new file.
	// 	panic(err)
	// }
	// fmt.Print(fileContents)


	// page := Page{
	// 	TextFilePath: *file,
	// 	TextFileName: "test",
	// 	HTMLPagePath: *file+".html",
	// 	Content:      string(fileContents),
	// }

	// // Create a new template in memory named "template.tmpl".
	// // When the template is executed, it will parse template.tmpl,
	// // looking for {{ }} where we can inject content.
	// t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// // Create a new, blank HTML file.
	// newFile, err := os.Create(page.HTMLPagePath)
	// if err != nil {
	// 		panic(err)
	// }

	// // Executing the template injects the Page instance's data,
    //     // allowing us to render the content of our text file.
    //     // Furthermore, upon execution, the rendered template will be
    //     // saved inside the new file we created earlier.
    // t.Execute(newFile, page)


	// if err != nil {
	// 	// A common use of `panic` is to abort if a function returns an error
	// 	// value that we don’t know how to (or want to) handle. This example
	// 	// panics if we get an unexpected error when creating a new file.
	// 	panic(err)
	// }
	
}