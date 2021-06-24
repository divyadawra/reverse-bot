package utils

import "fmt"

// ReverseInterface to implement reverse function for various types of files
type ReverseInterface interface {
	Reverse()
}

// GifFile structure refers to gif File
type GifFile struct {
	fileName string
}

// Mp4File structure refers to mp4 File
type Mp4File struct {
	fileName string
}

// Reverse is implemented for gif file
func (f *GifFile) Reverse() {
	fmt.Println("in gif reverse")
}

// Reverse is implemented for mp4 file
func (f *Mp4File) Reverse() {
	fmt.Println("in mp4 reverse")
}
