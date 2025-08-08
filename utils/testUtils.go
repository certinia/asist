package utils

import (
	"log"
	"os"
)

func DeleteFolder(dir string) {

	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Printf("Directory does not exist: %s", dir)
		return
	}
	if err != nil {
		log.Fatalf("Error accessing directory %s: %v", dir, err)
	}

	if !info.IsDir() {
		log.Fatalf("Path is a file, not a directory: %s", dir)
	}

	err = os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteFile(path string) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Printf("File does not exist: %s", path)
		return
	}
	if err != nil {
		log.Fatalf("Error accessing file %s: %v", path, err)
	}

	if info.IsDir() {
		log.Fatalf("Path is a directory, not a file: %s", path)
	}

	if err := os.Remove(path); err != nil {
		log.Fatalf("Failed to delete file %s: %v", path, err)
	}
}

func CreateFolder(dir string) {
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFile(file string) *os.File {
	destinationFile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer destinationFile.Close()
	return destinationFile
}

func WriteFile(file string, content []byte) {
	err := os.WriteFile(file, content, 0660)
	if err != nil {
		log.Fatal(err)
	}
}
