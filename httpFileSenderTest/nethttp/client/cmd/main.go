package main

import (
	"fmt"
	"path/filepath"

	nethttpClient "httpFileSenderTest/nethttp/client"
	"httpFileSenderTest/pathvalidator"
)

const filePath = "./testdata/file-sample_150kB.pdf"

// 284.142042ms - 100 MB
// 2.590718291s - 589 MB
// 44.12584125s - 10 GB
func main() {
	filename := filepath.Join(pathvalidator.GetBaseDir(), filePath)
	targetUrl := "http://localhost:8081/upload"

	// t1 := time.Now()

	err := nethttpClient.UploadFile(filename, targetUrl)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
	}

	// t2 := time.Now()
	// fmt.Println(t2.Sub(t1))
}
