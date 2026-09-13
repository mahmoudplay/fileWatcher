package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)


func makeBackup(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("file does not exist")
		}

		return err
	}

	src, err := os.Open(path)
	if err != nil {
		return err
	}
	defer src.Close()

	content, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	compressed, err := CompressFile(string(content))
	if err != nil {
		return fmt.Errorf("compress failed: %w", err)
	}

	backup, err := os.Create(filepath.Join("./.backup", fileNameMaker(info.Name(), info.ModTime())))
	if err != nil {
		return err
	}
	defer backup.Close()

	_, err = backup.Write(compressed)
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}

	fmt.Printf("%s has been backed up successfully to %s\n",
		info.Name(),
		filepath.Join("./.backup", fileNameMaker(info.Name(), info.ModTime())))

	return nil
}