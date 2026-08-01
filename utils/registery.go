package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var lastModified = make(map[string]string)

func RegisterFiles(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if err := os.MkdirAll("./.backup", 0755); err != nil {
		return err
	}


	if !info.IsDir() {
		last, exists := lastModified[path]

		hash, err := fileHash(path)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		if !exists {
			lastModified[path] = hash

			if err := makeBackup(path); err != nil {
				return fmt.Errorf("Failed to create backup: %w", err)
			}

			return nil
		}

		if hash != last {
			if err := makeBackup(path); err != nil {
				return fmt.Errorf("Failed to create backup: %w", err)
			}

			lastModified[path] = hash
		}

		return nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		err := RegisterFiles(filepath.Join(path, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}