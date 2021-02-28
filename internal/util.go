package internal

import "os"

func CreateFolderIfNotExists(folder string) error {
	if dir, err := os.Open(folder); os.IsNotExist(err) {
		return os.MkdirAll(folder, 0700)
	} else {
		_ = dir.Close()
	}
	return nil
}
