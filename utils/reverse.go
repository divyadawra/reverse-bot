package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

// Reverse a GIF file
func Reverse(mp4FilePath string) {

	breakIntoImages(mp4FilePath)
	cmd := exec.Command("goanigiffy", "-src=./reversed-images/*.png", "-dest=./files/reversed.gif", "-delay=30")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	}
	fmt.Println(string(output))
}

func breakIntoImages(mp4FilePath string) {
	err := downloadFile("./files/original.mp4", mp4FilePath)
	if err != nil {
		fmt.Println("There is error")
		fmt.Println(err)
	}
	fmt.Println("----here----")
	os.RemoveAll("./images")
	os.MkdirAll("./images", 0755)

	cmd := exec.Command("ffmpeg", "-i", "./files/original.mp4", "-vf", "fps=6", "./images/%03d.png")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	}
	fmt.Println(string(output))

	files, _ := ioutil.ReadDir("./images")

	// Save the images in reversed order in another folder
	os.RemoveAll("./reversed-images")
	os.MkdirAll("./reversed-images", 0755)
	for i := 0; i < len(files); i++ {
		fileName := strconv.Itoa(len(files)-i) + ".png"
		data, err := ioutil.ReadFile("./images/" + files[i].Name())
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("./reversed-images/"+fileName, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
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
