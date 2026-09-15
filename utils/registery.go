package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var lastModified = make(map[string]string)

const stateFile = ".backup/state.json"

var stateOnce sync.Once
var stateLoadErr error

func loadState() error {
	stateOnce.Do(func() {
		data, err := os.ReadFile(stateFile)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return
			}
			stateLoadErr = err
			return
		}
		stateLoadErr = json.Unmarshal(data, &lastModified)
	})
	return stateLoadErr
}

func saveState() error {
	if err := os.MkdirAll(filepath.Dir(stateFile), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(lastModified, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(stateFile, data, 0644)
}

func CountFiles(path string) (int, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return 1, nil
	}

	count := 0
	err = filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			count++
		}
		return nil
	})
	return count, err
}

func RegisterFiles(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if err := loadState(); err != nil {
		return fmt.Errorf("failed to load state: %w", err)
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
				return fmt.Errorf("failed to create backup: %w", err)
			}

			if err := saveState(); err != nil {
				return fmt.Errorf("failed to save state: %w", err)
			}

			return nil
		}

		if hash != last {
			if err := makeBackup(path); err != nil {
				return fmt.Errorf("failed to create backup: %w", err)
			}

			lastModified[path] = hash

			if err := saveState(); err != nil {
				return fmt.Errorf("failed to save state: %w", err)
			}
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
