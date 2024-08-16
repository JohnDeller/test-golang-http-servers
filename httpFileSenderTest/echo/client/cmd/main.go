package main

import (
	"fmt"
	"path/filepath"

	echoClient "httpFileSenderTest/echo/client"
	"httpFileSenderTest/pathvalidator"
)

const filePath = "./testdata/file-sample_150kB.pdf"

// 553.152541ms - 100 MB
// 1.770650417s - 589 MB
// 59.672896s - 10 GB
func main() {
	filename := filepath.Join(pathvalidator.GetBaseDir(), filePath)
	targetUrl := "http://localhost:8084/upload"

	// t1 := time.Now()

	err := echoClient.UploadFile(filename, targetUrl)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
	}

	// t2 := time.Now()
	// fmt.Println(t2.Sub(t1))
}
