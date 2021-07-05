package utils

import (
	"fmt"

	"github.com/hunterlong/gifs"
)

func Reverse(mp4FilePath string) string {

	gifPath, err := convertToGif(mp4FilePath)
	if err != nil {
		fmt.Println(err)
		// return "", err
	}
	// fmt.Println(gifPath)
	return gifPath
}

func convertToGif(mp4FilePath string) (string, error) {
	input := &gifs.New{
		Source: mp4FilePath,
	}

	response, err := input.Create()
	if err != nil {
		return "", err
	}
	fmt.Println("Gifs.com Gif URL: ", response.Files.Gif)
	return response.Files.Gif, nil
}
