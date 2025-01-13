package shared

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	FolderPath string
	FileName   string
}

func FindFiles(extension string) []FileInfo {
	// folderPath := "./Documents"
	folderPath := getDocumentsDirectoryPath()

	// Read the directory contents
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		panic("Invalid Directory!")
	}

	var fileList []FileInfo

	// Iterate over the files and remove .txt files
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), extension) {
			file := FileInfo{FolderPath: folderPath, FileName: entry.Name()}
			fileList = append(fileList, file)
		}
	}

	return fileList
}

func getDocumentsDirectoryPath() string {

	usr, _ := user.Current()

	documents := filepath.Join(usr.HomeDir, "Documents")

	return documents
}
