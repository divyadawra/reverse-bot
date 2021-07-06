package utils

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/hunterlong/gifs"
)

// Reverse a GIF file
func Reverse(mp4FilePath string) {

	gifPath, err := convertToGif(mp4FilePath)
	if err != nil {
		fmt.Println(err)
	}
	err = downloadFile("./files/original.gif", gifPath)

	originalGif := &gif.GIF{}
	reversedGif := &gif.GIF{}
	fmt.Println("Starting...")
	f, err := os.Open("./files/original.gif")
	if err != nil {
		fmt.Println("Error")
		return
	}
	originalGif, _ = gif.DecodeAll(f)

	images := originalGif.Image

	delay := originalGif.Delay
	fmt.Println(delay)
	for i := len(images) - 1; i > 0; i-- {
		name := strconv.Itoa(i) + ".jpg"
		file, _ := os.Create(name)

		opt := jpeg.Options{
			Quality: 100,
		}
		err = jpeg.Encode(file, images[i], &opt)
		reversedGif.Image = append(reversedGif.Image, images[i])
		reversedGif.Delay = append(reversedGif.Delay, delay[i])
	}

	f, _ = os.OpenFile("./files/reversed.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, reversedGif)
	if err != nil {
		fmt.Println(err)
	}
}

// This is time taking TODO: make it fast
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

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
