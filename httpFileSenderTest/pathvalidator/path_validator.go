package pathvalidator

import (
	"log"
	"os"
	"path/filepath"
)

func GetBaseDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}
