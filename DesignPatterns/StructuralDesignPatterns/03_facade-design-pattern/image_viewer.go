package main

import "fmt"

type ImageViewer struct {
}

func (i *ImageViewer) LoadImage(fileName string) {
	// Load the image from the file
	fmt.Println("Loading image:", fileName)
}

func (i *ImageViewer) DisplayImage() {
	// Display the loaded image
	fmt.Println("Displaying image")
}
