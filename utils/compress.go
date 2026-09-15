package utils

import (
	"io"

	"github.com/klauspost/compress/zstd"
)

func CompressFile(src io.Reader, dst io.Writer) error {
	encoder, err := zstd.NewWriter(dst)
	if err != nil {
		return err
	}
	defer encoder.Close()

	_, err = io.Copy(encoder, src)
	return err
}
