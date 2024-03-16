package utils

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// ResizeImage mengecilkan ukuran gambar sesuai dengan dimensi yang ditentukan.
func ResizeImage(inputPath string, outputPath string, width int, height int) error {
	// Buka file gambar
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Dekode gambar
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Resize gambar
	resized := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	// Buat file output
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Tulis gambar yang telah diresize ke file output
	err = jpeg.Encode(outputFile, resized, nil)
	if err != nil {
		return err
	}

	return nil
}
