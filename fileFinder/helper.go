package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CheckArguments() {
	// Check if the user has provided a directory argument
	if len(os.Args) < 2 {
		log.Fatal("Usage: program <directory>")
	}
}

func GetDirectory(dir string) {
	// First command line argument is the directory to scan
	// Read the directory contents
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		// Check if it's a regular file (not a directory)
		if !file.IsDir() {
			fileName := file.Name()
			fileLength := len(fileName)
			fileextension := filepath.Ext(fileName)

			// Check if the file name length is greater than 10
			if fileLength > 10 {
				newName := fileName[:10] + fileextension // Truncate file name to 10 characters

				// Construct the full old and new file paths
				oldPath := filepath.Join(dir, fileName)
				newPath := filepath.Join(dir, newName)

				// Rename the file
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Printf("Failed to rename %s to %s: %v\n", oldPath, newPath, err)
				} else {
					fmt.Printf("Renamed %s to %s\n", oldPath, newPath)
				}
			} else {
				fmt.Println("no files found to rename that are greater than 10 characters")
			}
		}
	}

}

func GetDirectoryRecursively(dirToWalk string) {
	err := filepath.WalkDir(dirToWalk, func(path string, d os.DirEntry, err error) error {

		if err != nil {

			fmt.Printf("Error acccesing path %q: %v\n", path, err)
			return err
		}

		if !d.IsDir() {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", dirToWalk, err)
		return
	}
}
