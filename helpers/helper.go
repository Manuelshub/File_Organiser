package helpers

import (
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
