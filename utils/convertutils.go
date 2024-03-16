package utils

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

// ConvertPNGToJPEG converts a PNG image to JPEG format.
func ConvertPNGToJPEG(pngPath string, jpegPath string) error {
	// Open PNG file
	pngFile, err := os.Open(pngPath)
	if err != nil {
		return fmt.Errorf("failed to open PNG file: %v", err)
	}
	defer pngFile.Close()

	fmt.Println("Successfully opened PNG file")

	// Decode PNG image
	pngImg, err := png.Decode(pngFile)
	if err != nil {
		return fmt.Errorf("failed to decode PNG image: %v", err)
	}

	fmt.Println("Successfully decoded PNG image")

	// Create JPEG file
	jpegFile, err := os.Create(jpegPath)
	if err != nil {
		return fmt.Errorf("failed to create JPEG file: %v", err)
	}
	defer jpegFile.Close()

	fmt.Println("Successfully created JPEG file")

	// Encode image to JPEG
	err = jpeg.Encode(jpegFile, pngImg, nil)
	if err != nil {
		return fmt.Errorf("failed to encode image to JPEG: %v", err)
	}

	fmt.Println("Successfully encoded image to JPEG")

	return nil
}
