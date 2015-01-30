package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("modstat: Error: Expected file.")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fail(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fail(err)
	}

	mode := fileInfo.Mode()
	if mode.IsDir() {
		mode = mode ^ os.ModeDir
	}

	fmt.Printf("%04o\t%v\n", mode, path.Clean(fileName))
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}
