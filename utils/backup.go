package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func makeBackup(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("file does not exist")
		}

		return err
	}

	backupPath := uniqueBackupPath(filepath.Join("./.backup", fileNameMaker(info.Name(), time.Now())))

	backup, err := os.Create(backupPath)
	if err != nil {
		return err
	}
	defer backup.Close()

	src, err := os.Open(path)
	if err != nil {
		return err
	}
	defer src.Close()

	if err := CompressFile(src, backup); err != nil {
		return fmt.Errorf("compress failed: %w", err)
	}

	fmt.Printf("%s has been backed up successfully to %s\n", info.Name(), backupPath)

	return nil
}
