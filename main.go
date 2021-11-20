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
	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"
	colorBlack := "\033[30m"

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
	// fmt.Println(decodedImage.At(100, 100))

	doubleWide := true
	scaleDownBy := 4

	levels := []string{" ", "░", "▒", "▓", "█"}
	// levels := []string{" ", ".", "▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▉"}
	// levels := []string{"▔", "▝", "▞", "▛", "█"}

	var valPerLevel int = 255 / len(levels)

	maxLevel := len(levels) - 1
	lastColorUsed := ""

	for y := decodedImage.Bounds().Min.Y; y < decodedImage.Bounds().Max.Y; y += scaleDownBy {
		// iterate through all X's first at each Y for printing
		for x := decodedImage.Bounds().Min.X; x < decodedImage.Bounds().Max.X; x += scaleDownBy {
			colorStrength := color.GrayModel.Convert(decodedImage.At(x, y)).(color.Gray)
			level := int(colorStrength.Y) / valPerLevel

			// decide what color to use
			red := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).R
			green := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).G
			blue := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).B

			rgbMax := maxUnint8(red, green, blue)
			rgbMin := minUnint8(red, green, blue)
			switch {

			// no color is very bright
			case lastColorUsed != "black" && rgbMax < 50:
				fmt.Print(colorBlack)
				lastColorUsed = "black"

			// all colors are bright
			case lastColorUsed != "white" && rgbMin > 200:
				fmt.Print(colorWhite)
				lastColorUsed = "white"

			// red and blue are close, green isn't
			case lastColorUsed != "purple" && rgbMin == green && within50(red, blue) && !within50(red, green):
				fmt.Print(colorPurple)
				lastColorUsed = "purple"

			// red and green are close, blue isn't
			case lastColorUsed != "yellow" && rgbMin == blue && within50(red, green) && !within50(red, blue):
				fmt.Print(colorYellow)
				lastColorUsed = "yellow"

			// green and blue are close, red isn't
			case lastColorUsed != "cyan" && rgbMin == red && within50(green, blue) && !within50(green, red):
				fmt.Print(colorCyan)
				lastColorUsed = "cyan"

			// red is dominant
			case lastColorUsed != "red" && rgbMax == red:
				fmt.Print(colorRed)
				lastColorUsed = "red"
			// green is dominant
			case lastColorUsed != "green" && rgbMax == green:
				fmt.Print(colorGreen)
				lastColorUsed = "green"
			// blue is dominant
			case lastColorUsed != "blue" && rgbMax == blue:
				fmt.Print(colorBlue)
				lastColorUsed = "blue"

			}

			// decide how bold the ascii "pixel" should be
			if level > maxLevel {
				level = maxLevel
			}
			if doubleWide { // print it an extra time
				fmt.Print(levels[level], levels[level])
			}
			fmt.Print(levels[level]) // print the ascii "pixel"
		}
		fmt.Print("\n")
	}
	fmt.Print(colorReset)

}
func within50(a, b uint8) bool {
	return a > b-50 && a < b+50
}
func minUnint8(values ...uint8) uint8 {
	min := uint8(255)
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}
func maxUnint8(values ...uint8) uint8 {
	max := uint8(0)
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
