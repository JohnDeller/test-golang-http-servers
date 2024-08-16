package main

import (
	"fmt"
	"path/filepath"

	chiClient "httpFileSenderTest/chi/client"
	"httpFileSenderTest/pathvalidator"
)

const filePath = "./testdata/file-sample_150kB.pdf"

// 544.922042ms - 100 MB
// 1.830936125s - 589 MB
// 1m2.72820175s - 10 GB
func main() {
	filename := filepath.Join(pathvalidator.GetBaseDir(), filePath)
	targetUrl := "http://localhost:8085/upload"

	// t1 := time.Now()

	err := chiClient.UploadFile(filename, targetUrl)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
	}

	// t2 := time.Now()
	// fmt.Println(t2.Sub(t1))
}
