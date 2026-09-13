package utils

import (
	"github.com/klauspost/compress/zstd"
)

func CompressFile(data string) ([]byte, error) {
	encoder, err := zstd.NewWriter(nil)

	if err != nil {
		return nil, err
	}
	defer encoder.Close()
	compressed := encoder.EncodeAll([]byte(data), nil)

	return compressed, nil
}