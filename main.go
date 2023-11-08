package main

import (
	"os"
)

func main() {
	//CheckArguments()

	// First command line argument is the directory to scan
	dir := os.Args[1]

	GetDirectory(dir)

	GetDirectoryRecursively(dir)

}
