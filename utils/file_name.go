package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func fileNameMaker(fileName string, lastEdit time.Time) string {
	ext := filepath.Ext(fileName)
	name := strings.TrimSuffix(fileName, ext)

	return name +
		"_" +
		lastEdit.Format("20060102_150405") +
		ext +
		".zst"
}

func uniqueBackupPath(base string) string {
	if _, err := os.Stat(base); errors.Is(err, os.ErrNotExist) {
		return base
	}

	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	for i := 1; ; i++ {
		candidate := fmt.Sprintf("%s_%d%s", name, i, ext)
		if _, err := os.Stat(candidate); errors.Is(err, os.ErrNotExist) {
			return candidate
		}
	}
}
