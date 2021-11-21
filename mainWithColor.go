package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	debugSwitch := false

	var maxPixelWidth int = 60
	colorDistance := 80
	doubleWide := false

	pngImage, err := os.Open("C:\\zHolderFolder\\color-wheel.png")
	if err != nil {
		log.Fatal(err)
	}
	defer pngImage.Close()

	decodedImage, err := png.Decode(pngImage)
	if err != nil {
		log.Fatal(err)
	}

	// ANSI color codes:
	colorReset := "\033[0m"

	colorRed := "\033[31m"
	// colorBrightRed := "\033[91m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"
	colorBlack := "\033[30m"

	levels := []string{" ", ".", "▁", "▂", "▃", "▄", "▅", "▆", "▇", "▉", "█"}

	var valPerLevel int = 255 / len(levels)
	maxLevel := len(levels) - 1
	scaleDownBy := decodedImage.Bounds().Max.X / maxPixelWidth
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
			rCloseToB := withinRangeOf(red, blue, colorDistance)
			rCloseToG := withinRangeOf(red, green, colorDistance)
			gCloseToB := withinRangeOf(green, blue, colorDistance)

			rgbMax := maxUint8(red, green, blue)
			rgbMin := minUint8(red, green, blue)
			if debugSwitch {
				fmt.Println(colorWhite, "Max, Min // R G B ", rgbMax, rgbMin, "//", colorRed, red, colorGreen, green, colorBlue, blue)
			}
			switch {

			// no color is very bright
			case lastColorUsed != "black" && rgbMax < 80:
				fmt.Print(colorBlack)
				lastColorUsed = "black"

			// all colors are bright
			case lastColorUsed != "white" && rgbMin > 200:
				fmt.Print(colorWhite)
				lastColorUsed = "white"

			// red and blue are close, green isn't
			case lastColorUsed != "purple" && rgbMin == green && rCloseToB && !rCloseToG:
				fmt.Print(colorPurple)
				lastColorUsed = "purple"

			// red and green are close, blue isn't
			case lastColorUsed != "yellow" && rgbMin == blue && rCloseToG && !rCloseToB:
				fmt.Print(colorYellow)
				lastColorUsed = "yellow"

			// green and blue are close, red isn't
			case lastColorUsed != "cyan" && rgbMin == red && gCloseToB && !rCloseToG:
				fmt.Print(colorCyan)
				lastColorUsed = "cyan"

			// red is dominant
			case lastColorUsed != "red" && rgbMax == red && !rCloseToG && !rCloseToB:
				fmt.Print(colorRed)
				lastColorUsed = "red"
			// green is dominant
			case lastColorUsed != "green" && rgbMax == green && !rCloseToG && !gCloseToB:
				fmt.Print(colorGreen)
				lastColorUsed = "green"
			// blue is dominant
			case lastColorUsed != "blue" && rgbMax == blue && !gCloseToB && !rCloseToB:
				fmt.Print(colorBlue)
				lastColorUsed = "blue"
			}

			if debugSwitch {
				fmt.Println("Chose Color: ", lastColorUsed)
				fmt.Println("")
			}

			// decide how bold the ascii "pixel" should be
			if level > maxLevel {
				level = maxLevel
			}
			if doubleWide { // print it an extra time
				fmt.Print(levels[level])
			}
			fmt.Print(levels[level]) // print the ascii "pixel"
		}
		fmt.Print("\n")
	}
	fmt.Print(colorReset)

}
func withinRangeOf(a, b uint8, distance int) bool {
	intA, intB := int(a), int(b)
	return intA > intB-distance && intA < intB+distance

}
func minUint8(values ...uint8) uint8 {
	min := uint8(255)
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}
func maxUint8(values ...uint8) uint8 {
	max := uint8(0)
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
