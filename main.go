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
	"pickcolor"
	// "htmlcreator"
)

func main() {

	useColor := false
	useColor = true
	colorDistanceRequirement := 80
	colorCloseToRequirement := 60

	ansiOrHtml := "ansi"
	// ansiOrHtml := "html"

	var maxPixelWidth int = 80
	doubleWide := true
	// doubleWide = false

	// levels := []string{"_", "a", "b", "c", "d", "e", "f"}
	// levels := []string{" ", "-", "+", "#", "%", "&"}
	// levels := []string{" ", "-", "+", "#"}
	levels := []string{" ", "░", "▒", "▓", "█"}
	// levels := []string{"▁", "░", "▒", "▓", "█"}
	// levels := []string{" ", ".", "▁", "▂", "▃", "▄", "▅", "▆", "▇", "▉", "█"}

	if doubleWide {
		maxPixelWidth /= 2 //half the width of pixels actually scanned
	}
	// if maxPixelWidth < 2 {
	// 	maxPixelWidth = 2
	// }

	// ~~~ image to use
	// imagePointer, err := os.Open("C:\\zHolderFolder\\cat1.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\windowPainting.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\ColinPicture3.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\color1.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\color2.jpg")
	imagePointer, err := os.Open("C:\\zHolderFolder\\color-wheel.png")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\colorSquares.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imagePointer.Close()

	decodedImage, _, err := image.Decode(imagePointer)
	if err != nil {
		log.Fatal(err)
	}

	// ANSI color codes:
	ansiColorReset := "\033[0m"

	ansiColorRed := "\033[31m"
	ansiColorGreen := "\033[32m"
	ansiColorYellow := "\033[33m"
	ansiColorBlue := "\033[34m"
	ansiColorPurple := "\033[35m"
	ansiColorCyan := "\033[36m"
	ansiColorWhite := "\033[37m"
	ansiColorBlack := "\033[30m"

	// html color codes:
	htmlColorEnd := "</span>"
	htmlColorRed := "</span><span style=\"color: #ff0000;\">"
	htmlColorGreen := "</span><span style=\"color: #00ff00;\">"
	htmlColorYellow := "</span><span style=\"color: #ffff00;\">"
	htmlColorBlue := "</span><span style=\"color: #0000ff;\">"
	htmlColorPurple := "</span><span style=\"color: #ff00ff;\">"
	htmlColorCyan := "</span><span style=\"color: #00ffff;\">"
	htmlColorWhite := "</span><span style=\"color: #ffffff;\">"
	htmlColorBlack := "</span><span style=\"color: #000000;\">"

	var valPerIntensityLevel int = 255 / len(levels)
	maxIntensity, currentIntensity := len(levels)-1, 0
	scaleDownBy := decodedImage.Bounds().Max.X / maxPixelWidth

	htmlColorCode, ansiColorCode := "", ""
	lastColorUsed := ""
	for y := decodedImage.Bounds().Min.Y; y < decodedImage.Bounds().Max.Y; y += scaleDownBy {
		// iterate through all X's first at each Y for printing format
		for x := decodedImage.Bounds().Min.X; x < decodedImage.Bounds().Max.X; x += scaleDownBy {

			// PICK INTENSITY
			grayScaleIntensity := color.GrayModel.Convert(decodedImage.At(x, y)).(color.Gray)
			// fmt.Println("grayScaleIntensity:", grayScaleIntensity)//DEBUG LINE

			if !useColor {
				// currentIntensity = pickcolor.DecideIntensityWithGray(grayScaleIntensity, valPerIntensityLevel,maxIntensity)
				currentIntensity = int(grayScaleIntensity.Y) / valPerIntensityLevel
				// fmt.Println("currentIntensity:", currentIntensity)
			}
			if useColor {
				// decide what color to use
				pRed := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).R
				pGreen := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).G
				pBlue := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).B
				pAlpha := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).A

				currentIntensity = pickcolor.DecideIntensityWithColor(pRed, pGreen, pBlue, pAlpha, valPerIntensityLevel, maxIntensity)
				// fmt.Println("currentIntensity was set to:", currentIntensity) //DEBUG LINE
				// ~~~ PICK COLOR ~~~
				colorToUse := pickcolor.PickColor(pRed, pGreen, pBlue, pAlpha, colorDistanceRequirement, colorCloseToRequirement)
				// fmt.Println("picked color:", colorToUse)//DEBUG LINE
				// set color code variables
				switch {
				case colorToUse == "black" && lastColorUsed != "black":
					ansiColorCode = ansiColorBlack
					htmlColorCode = htmlColorBlack
				case colorToUse == "white" && lastColorUsed != "white":
					ansiColorCode = ansiColorWhite
					htmlColorCode = htmlColorWhite
				case colorToUse == "purple" && lastColorUsed != "purple":
					ansiColorCode = ansiColorPurple
					htmlColorCode = htmlColorPurple
				case colorToUse == "yellow" && lastColorUsed != "yellow":
					ansiColorCode = ansiColorYellow
					htmlColorCode = htmlColorYellow
				case colorToUse == "cyan" && lastColorUsed != "cyan":
					ansiColorCode = ansiColorCyan
					htmlColorCode = htmlColorCyan
				case colorToUse == "red" && lastColorUsed != "red":
					ansiColorCode = ansiColorRed
					htmlColorCode = htmlColorRed
				case colorToUse == "green" && lastColorUsed != "green":
					ansiColorCode = ansiColorGreen
					htmlColorCode = htmlColorGreen
				case colorToUse == "blue" && lastColorUsed != "blue":
					ansiColorCode = ansiColorBlue
					htmlColorCode = htmlColorBlue
				default:
					ansiColorCode = ""
					htmlColorCode = ""
				}

				// set the color of the pixel
				if ansiOrHtml == "ansi" {
					fmt.Print(ansiColorCode)
				} else if ansiOrHtml == "html" {
					fmt.Print(htmlColorCode)
				}
				lastColorUsed = colorToUse

			} // END use color if statement

			// correct intensity if needed
			if currentIntensity > maxIntensity {
				currentIntensity = maxIntensity
				// fmt.Println("fixed intensity:", currentIntensity) //DEBUG LINE
			}

			// PRINT THE PIXEL
			// fmt.Println("about to use intensity:", currentIntensity) //DEBUG LINE
			fmt.Print(levels[currentIntensity])
			if doubleWide {
				fmt.Print(levels[currentIntensity]) // print an extra time if dobule wide
			}
		}

		// print a newline after each row
		if ansiOrHtml == "ansi" {
			fmt.Print("\n")
		} else if ansiOrHtml == "html" {
			fmt.Print("<br />")
		}
	}
	// reset the color when done
	if ansiOrHtml == "ansi" {
		fmt.Print(ansiColorReset)
	} else if ansiOrHtml == "html" {
		fmt.Print(htmlColorEnd)
	}

}
