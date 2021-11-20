package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {

	// ANSI color codes:
	// colorReset := "\033[0m"

	// colorRed := "\033[31m"
	// colorGreen := "\033[32m"
	// colorYellow := "\033[33m"
	// colorBlue := "\033[34m"
	// colorPurple := "\033[35m"
	// colorCyan := "\033[36m"
	// colorWhite := "\033[37m"

	// pngImage, err := os.Open("C:\\zHolderFolder\\anime-cat2Small.png")
	// pngImage, err := os.Open("C:\\zHolderFolder\\triangle.png")
	pngImage, err := os.Open("C:\\zHolderFolder\\colorSquaresSmall.png")
	if err != nil {
		log.Fatal(err)
	}
	defer pngImage.Close()

	decodedImage, err := png.Decode(pngImage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(decodedImage.At(100, 100))

	levels := []string{" ", "░", "▒", "▓", "█"}
	// levels := []string{"▔", "▝", "▞", "▛", "█"}
	// doubleWide := false
	doubleWide := true
	scaleDownBy := 4

	for y := decodedImage.Bounds().Min.Y; y < decodedImage.Bounds().Max.Y; y += scaleDownBy {
		for x := decodedImage.Bounds().Min.X; x < decodedImage.Bounds().Max.X; x += scaleDownBy { // iterate through all X's first at each Y for printing
			colorStrength := color.GrayModel.Convert(decodedImage.At(x, y)).(color.Gray)
			level := colorStrength.Y / 51 // 51 * 5 = 255
			if level > 4 {
				level = 4
			}
			if doubleWide {
				fmt.Print(levels[level], levels[level])
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}

}
