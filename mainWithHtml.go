package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	debugSwitch := false

	ansiOrHtml := "ansi"
	// ansiOrHtml := "html"
	colorDistance := 80
	doubleWide := true
	var maxPixelWidth int = 60

	levels := []string{" ", "░", "▒", "▓", "█"}
	// levels := []string{"▁", "░", "▒", "▓", "█"}
	// levels := []string{" ", ".", "▁", "▂", "▃", "▄", "▅", "▆", "▇", "▉", "█"}

	if doubleWide {
		maxPixelWidth /= 2
	}

	// imagePointer, err := os.Open("C:\\zHolderFolder\\color1.jpg")
	imagePointer, err := os.Open("C:\\zHolderFolder\\color-wheel.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imagePointer.Close()

	// decodedImage, fileType, err := image.Decode(imagePointer)
	decodedImage, _, err := image.Decode(imagePointer)
	if err != nil {
		log.Fatal(err)
	}

	// ANSI color codes:
	ansiColorReset := "\033[0m"

	ansiColorRed := "\033[31m"
	// ansiColorBrightRed := "\033[91m"
	ansiColorGreen := "\033[32m"
	ansiColorYellow := "\033[33m"
	ansiColorBlue := "\033[34m"
	ansiColorPurple := "\033[35m"
	ansiColorCyan := "\033[36m"
	ansiColorWhite := "\033[37m"
	ansiColorBlack := "\033[30m"

	// html color codes:
	// htmlColorReset := "<span style=\"color: #000000;\">"
	htmlColorReset := "</span>"
	htmlColorRed := "</span><span style=\"color: #ff0000;\">"
	htmlColorGreen := "</span><span style=\"color: #00ff00;\">"
	htmlColorYellow := "</span><span style=\"color: #ffff00;\">"
	htmlColorBlue := "</span><span style=\"color: #0000ff;\">"
	htmlColorPurple := "</span><span style=\"color: #ff00ff;\">"
	htmlColorCyan := "</span><span style=\"color: #00ffff;\">"
	htmlColorWhite := "</span><span style=\"color: #ffffff;\">"
	htmlColorBlack := "</span><span style=\"color: #000000;\">"

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
				fmt.Println(ansiColorWhite, "Max, Min // R G B ", rgbMax, rgbMin, "//", ansiColorRed, red, ansiColorGreen, green, ansiColorBlue, blue)
			}
			switch {

			// no color is very bright
			case lastColorUsed != "black" && rgbMax < 80:
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorBlack)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorBlack)
				}
				lastColorUsed = "black"

			// all colors are bright
			case lastColorUsed != "white" && rgbMin > 200:
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorWhite)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorWhite)
				}
				lastColorUsed = "white"

			// red and blue are close, green isn't
			case lastColorUsed != "purple" && rgbMin == green && rCloseToB && !rCloseToG:
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorPurple)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorPurple)
				}
				lastColorUsed = "purple"

			// red and green are close, blue isn't
			case lastColorUsed != "yellow" && rgbMin == blue && rCloseToG && !rCloseToB:
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorYellow)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorYellow)
				}
				lastColorUsed = "yellow"

			// green and blue are close, red isn't
			case lastColorUsed != "cyan" && rgbMin == red && gCloseToB && !rCloseToG:
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorCyan)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorCyan)
				}
				lastColorUsed = "cyan"

			// red is dominant
			case lastColorUsed != "red" && rgbMax == red && !rCloseToG && !rCloseToB:
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorRed)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorRed)
				}
				lastColorUsed = "red"
			// green is dominant
			case lastColorUsed != "green" && rgbMax == green && !rCloseToG && !gCloseToB:
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorGreen)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorGreen)
				}
				lastColorUsed = "green"
			// blue is dominant
			case lastColorUsed != "blue" && rgbMax == blue && !gCloseToB && !rCloseToB:
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorBlue)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorBlue)
				}
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

		if ansiOrHtml == "ansi" {
			fmt.Print("\n")
		} else if ansiOrHtml == "html" {
			// fmt.Print("</span><br />")
			fmt.Print("<br />")
		}
	}
	if ansiOrHtml == "ansi" {
		fmt.Print(ansiColorReset)
	} else if ansiOrHtml == "html" {
		fmt.Print(htmlColorReset)
	}

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
