package main

import (
	"fmt"
	"path/filepath"

	fiberClient "httpFileSenderTest/fiber/client"
	"httpFileSenderTest/pathvalidator"
)

const filePath = "./testdata/file-sample_150kB.pdf"

// 373.194875ms - 100 MB
// 1.70000125s - 589 MB
// 46.119087083s - 10 GB
func main() {
	filename := filepath.Join(pathvalidator.GetBaseDir(), filePath)
	targetUrl := "http://localhost:8082/upload"

	// t1 := time.Now()

	err := fiberClient.UploadFile(filename, targetUrl)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
	}

	// t2 := time.Now()
	// fmt.Println(t2.Sub(t1))
}
