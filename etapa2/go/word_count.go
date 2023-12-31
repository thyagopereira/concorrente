package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings" 
	"sync" 
)

// Count the number of words in the file at `filePath`.
func wc_file(filePath string){
	defer wg.Done()
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	
	mu.Lock()
	numberOfWords += len(strings.Fields(string(fileContent)))
	mu.Unlock()
}

// Count the number of words in all files directly within `directoryPath`.
// Files in subdirectories are not considered.
func wc_dir(directoryPath string){
	defer dirWg.Done()

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() { 
			filePath := directoryPath + "/" + file.Name()
			wg.Add(1)
			go wc_file(filePath)
		}

	}
	
	wg.Wait()
}

// Calculate the number of words in the files stored under the directory name
// available at argv[1].
//
// Assume a depth 3 hierarchy:
//   - Level 1: root
//   - Level 2: subdirectories
//   - Level 3: files
//
// root
// ├── subdir 1
// │     ├── file
// │     ├── ...
// │     └── file
// ├── subdir 2
// │     ├── file
// │     ├── ...
// │     └── file
// ├── ...
// └── subdir N
// │     ├── file
// │     ├── ...
// │     └── file

// Definindo variaveis globais
var numberOfWords = 0
var wg sync.WaitGroup
var dirWg sync.WaitGroup
var mu sync.Mutex

func main() {
	rootPath := os.Args[1] 

	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			dirWg.Add(1)
			directoryPath := rootPath + "/" + file.Name()
			
			go wc_dir(directoryPath)
		}
	}
	
	dirWg.Wait() 
	fmt.Println(numberOfWords)
}
