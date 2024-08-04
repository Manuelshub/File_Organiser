package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Manuelshub/File_Organiser/helpers"
)

func main() {
	videoExtensions := []string{".mp4", ".mkv", ".avi"}
	musicExtensions := []string{".mp3", ".wav", ".flac"}
	imageExtensions := []string{".jpg", ".png", ".jpeg", "gif", ".svg", ".webp", ".jfif"}
	docExtensions := []string{".pdf", ".doc", ".docx", ".txt"}
	otherExtensions := []string{".json", ".deb", ".zip", ".gz"}
	
	downloadsFolder, err := helpers.Downloads("~/Downloads")
	if err != nil {
		log.Fatalf("Failed to get the Downloads folder: %v\n", err)
	}

	folders := map[string]string{
		"Videos": "Videos",
		"Music": "Music",
		"Images": "Images",
		"Documents": "Documents",
		"Others": "Others",
	}

	for folderName, folderPath := range folders {
		if _, err := helpers.CreateDirectory(folderPath); err != nil {
			log.Printf("Failed to create the %s folder: %v\n", folderName, err)
		}
	}


	items, err := os.ReadDir(downloadsFolder)
	if err != nil {
		log.Fatalf("Failed to read the Downloads folder: %v\n", err)
	}

	for _, item := range items {
		if item.IsDir() {
			continue
		}

		filePath := filepath.Join(downloadsFolder, item.Name())
		extension := strings.ToLower(filepath.Ext(item.Name()))

		var targetFolder string
		switch {
		case helpers.Contains(videoExtensions, extension):
			targetFolder = filepath.Join(downloadsFolder, "Videos")
		case helpers.Contains(musicExtensions, extension):
			targetFolder = filepath.Join(downloadsFolder, "Music")
		case helpers.Contains(imageExtensions, extension):
			targetFolder = filepath.Join(downloadsFolder, "Images")
		case helpers.Contains(docExtensions, extension):
			targetFolder = filepath.Join(downloadsFolder, "Documents")
		case helpers.Contains(otherExtensions, extension):
			targetFolder = filepath.Join(downloadsFolder, "Others")
		default:
			continue
		}

		destinationPath := filepath.Join(targetFolder, item.Name())
		if err := os.Rename(filePath, destinationPath); err != nil {
			log.Printf("Failed to move %s to %s folder: %v\n", item.Name(), targetFolder, err)
		}
	}
}
