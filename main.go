package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func renameIndexFiles(mdPath string) {
	files, err := os.ReadDir(mdPath)
	check(err)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		// If there exists a file with the same name, then rename that file.
		potentialFile := fmt.Sprintf("%s\\%s.md", mdPath, file.Name())

		if _, err := os.Stat(potentialFile); err == nil {
			targetPath := strings.Trim(potentialFile, ".md") + "\\README.md"
			os.Rename(potentialFile, targetPath)
		}

		renameIndexFiles(fmt.Sprintf("%s\\%s", mdPath, file.Name()))
	}
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: ./convert <exported-dendron-path> <markdown-path>")
		return
	}

	dendronPath := args[1]
	mdPath := args[2]

	if !isDirectory(dendronPath) {
		fmt.Println("Error:", dendronPath, "is not a directory")
		return
	}

	if !isDirectory(mdPath) {
		fmt.Println("Error:", mdPath, "is not a directory")
		return
	}

	renameIndexFiles(mdPath)
}
