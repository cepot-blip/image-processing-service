package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"
)

// CompressImage mengompres gambar untuk mengurangi ukuran file sambil mempertahankan kualitas yang wajar.
func CompressImage(inputPath string, quality int) ([]byte, error) {
	// Buka file gambar
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Dekode gambar
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Kompresi gambar
	var opt jpeg.Options
	opt.Quality = quality

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, &opt)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
