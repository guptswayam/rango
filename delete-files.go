package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DeleteFiles() {
	folderName := "./"

	// Read the directory contents
	entries, err := os.ReadDir(folderName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate over the files and remove .txt files
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".txt") {
			filePath := filepath.Join(folderName, entry.Name())
			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf("Error deleting file %s: %v\n", entry.Name(), err)
			} else {
				fmt.Printf("Deleted file: %s\n", entry.Name())
			}
		}
	}

	fmt.Println("Congrats all txt files have been deleted!")
}
