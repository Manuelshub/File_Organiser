package helpers

import (
	"log"
	"os"
	"path/filepath"
)

// Expand function expands the '~' to the home path of the user
func Expand(path string) (string, error) {
	if path == "" || path[0] != '~' {
		return path, nil
	}
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, path[1:]), nil
}

// Downloads expands the home directory '~' and appends the downloads folder to it.
func Downloads(dir string) (string, error) {
	downloads, err := Expand(dir)
	if err != nil {
		log.Panic(err)
		return "", err
	}
	log.Println("Downloads PATH:", downloads)
	return downloads, nil
}

// CheckAndCreateFolder checks if a folder exists, if it does not it creates it.
func CheckAndCreateFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, 0750); err != nil {
			return err
		}
	}
	return nil
}

// CreateDirectory creates a directory and appends it to the user path
func CreateDirectory(dir string) (string, error) {
	// create each directory that files with different extension would be saved in
	downloads, err := Downloads("~/Downloads")
	if err != nil {
		log.Panic(err)
		return dir, err
	}
	newPath := downloads + "/" + dir
	if err = CheckAndCreateFolder(newPath); err != nil {
		return "", err
	}
	log.Printf("%s Path: %s", dir, newPath)
	return newPath, nil
}

// Contains checks if a string is contained in a given slice of string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}
