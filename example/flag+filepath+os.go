package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	fmt.Println("Hello World!")

	var dir string

	flag.StringVar(&dir, "d", "", "directory path")

	flag.Parse()

	fmt.Println(dir)

	paths := make([]string, 0)
	filepath.Walk("C:\\Users\\Leison\\Documents\\GO", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	for _, path := range paths {
		newPath := strings.Replace(path, ".txt", ".go", 1)
		os.Rename(path, newPath)
	}

	fmt.Printf("%v\n", paths)
}
