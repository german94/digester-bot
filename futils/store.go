package futils

import (
	"fmt"
	"os"
	"path/filepath"
)

func (*Service) StoreFileLocally(data []byte, chatID int64, fileName string) (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(currentPath, "chats", fmt.Sprintf("%v", chatID))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(dir, fileName), data, 0644); err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, fileName), nil
}
