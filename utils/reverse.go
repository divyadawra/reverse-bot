package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// Reverse a GIF file
func Reverse(mp4FilePath string) {
	
	os.Remove("./files/reversed.mp4")
	cmd := exec.Command("ffmpeg", "-i", mp4FilePath, "-vf", "reverse", "./files/reversed.mp4")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	}
	fmt.Println(string(output))
}
