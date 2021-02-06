package main

import (
	"fmt"
	"path/filepath"
)

// It works if we run `go run playground/main.go`
func main() {
	files := getFilesWithGlob()
	fmt.Println(files)
}

func getFilesWithGlob() []string {
	files, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

	return files
}
