package main

import (
	"fmt"
	"path/filepath"

	fasthttpClient "httpFileSenderTest/fasthttp/client"
	"httpFileSenderTest/pathvalidator"
)

const filePath = "./testdata/file-sample_150kB.pdf"

// 424.184541ms - 100 MB
// 1.720674s - 589 MB
// 48.616326s - 10 GB
func main() {
	filename := filepath.Join(pathvalidator.GetBaseDir(), filePath)
	targetUrl := "http://localhost:8083/upload"

	// t1 := time.Now()

	err := fasthttpClient.UploadFile(filename, targetUrl)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
	}

	// t2 := time.Now()
	// fmt.Println(t2.Sub(t1))
}
