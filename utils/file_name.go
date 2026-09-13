package utils

import (
	"path/filepath"
	"strings"
	"time"
)

func fileNameMaker(fileName string, lastEdit time.Time) string {

	ext := filepath.Ext(fileName)
	name := strings.TrimSuffix(fileName, ext)

	backupName := name +
		"_" +
		lastEdit.Format("20060102_150405") +
		ext +
		".zst"

	return backupName
}