package main

import (
	"fmt"
	"htmlcreator"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"pickcolor"
)

func main() {

	// ~~~~~ GLOBAL SETTINGS ~~~~~

	useColor := false // true || false; Grayscale or no
	useColor = true
	colorDistanceRequirement := 80 // 0-255; Distance between colors for them to be distinct
	colorCloseToRequirement := 60  // 0-255;  Distance ... to be close to each other, for blending to secondary colors

	ansiOrHtml := "ansi" // "ansi" || "html"; color coding for terminals etc. or html span tags for webpages
	ansiOrHtml = "html"

	var maxPixelWidth int = 1000

	doubleWide := true // Useful for monospace fonts, each pixel on the x-axis is printed twice
	// doubleWide = false

	// Double wide respects the maxPixelWidth setting
	if doubleWide {
		maxPixelWidth /= 2 //half the width of pixels actually desired since we will print them all twice
	}

	// ~~~~~ GET ASCII REPRESENTATION OF PIXEL INTENSITY ~~~~~

	intensityLevels := getIntensityLevelsSlice(0)

	// ~~~~~ GET IMAGE ~~~~~

	// imagePointer, err := os.Open("C:\\zHolderFolder\\cat1.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\windowPainting.jpg")
	imagePointer, err := os.Open("C:\\zHolderFolder\\ColinPicture3.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\color1.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\color2.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\color-wheel.png")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\ODDicon.png")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\sc-diamond-noTxt.png")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\colorSquares.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imagePointer.Close()

	decodedImage, _, err := image.Decode(imagePointer)
	if err != nil {
		log.Fatal(err)
	}

	// ~~~~~ CREATE COLOR MAPS ~~~~~

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

	// HTML color codes:
	htmlColorEnd := "</span>" // ~~~~~ ALL THE OTHER HTML COLOR CODES ALSO INCLUDE A CLOSING SPAN TAG ~~~~~
	htmlColorRed := "</span><span style=\"color: #ff0000;\">"
	htmlColorGreen := "</span><span style=\"color: #00ff00;\">"
	htmlColorYellow := "</span><span style=\"color: #ffff00;\">"
	htmlColorBlue := "</span><span style=\"color: #0000ff;\">"
	htmlColorPurple := "</span><span style=\"color: #ff00ff;\">"
	htmlColorCyan := "</span><span style=\"color: #00ffff;\">"
	htmlColorWhite := "</span><span style=\"color: #ffffff;\">"
	htmlColorBlack := "</span><span style=\"color: #000000;\">"

	// ~~~~~ CALC THE INTENSITY INCREASE REQUIRED FOR EACH CHARACTER IN intensityLevels ~~~~~
	var valPerIntensityLevel int = 255 / len(intensityLevels)
	maxIntensity, currentIntensity := len(intensityLevels)-1, 0
	scaleDownBy := decodedImage.Bounds().Max.X / maxPixelWidth

	// ~~~~~ SAFETY CHECKS ~~~~~
	// if maxPixelWidth < 2 {
	// 	maxPixelWidth = 2
	// }

	// ~~~~~ DECLARE GLOBALS FOR THE LOOP ~~~~~
	mainOutput, htmlColorCode, ansiColorCode, lastColorUsed := "", "", "", ""

	// ~~~~~ START THE LOOP! ~~~~~
	for y := decodedImage.Bounds().Min.Y; y < decodedImage.Bounds().Max.Y; y += scaleDownBy {
		// iterate through all X's first at each Y for printing format
		for x := decodedImage.Bounds().Min.X; x < decodedImage.Bounds().Max.X; x += scaleDownBy {

			// ~~~~~ PICK INTENSITY ~~~~~
			grayScaleIntensity := color.GrayModel.Convert(decodedImage.At(x, y)).(color.Gray)
			// fmt.Println("grayScaleIntensity:", grayScaleIntensity)//DEBUG LINE

			if !useColor {
				// currentIntensity = pickcolor.DecideIntensityWithGray(grayScaleIntensity, valPerIntensityLevel,maxIntensity)
				currentIntensity = int(grayScaleIntensity.Y) / valPerIntensityLevel
				// fmt.Println("currentIntensity:", currentIntensity)
			}

			if useColor { // ~~~~~ START USE COLOR IF STATEMENT ~~~~~
				// decide what color to use
				pRed := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).R
				pGreen := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).G
				pBlue := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).B
				pAlpha := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).A

				currentIntensity = pickcolor.DecideIntensityWithColor(pRed, pGreen, pBlue, pAlpha, valPerIntensityLevel, maxIntensity)
				// fmt.Println("currentIntensity was set to:", currentIntensity) //DEBUG LINE

				// ~~~~~ PICK COLOR ~~~~~
				colorToUse := pickcolor.PickColor(pRed, pGreen, pBlue, pAlpha, colorDistanceRequirement, colorCloseToRequirement)
				// fmt.Println("picked color:", colorToUse)//DEBUG LINE

				// ~~~~~ SET COLOR CODE BEFORE PRINTING PIXEL ~~~~~
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
				default: // will get used if colorToUse == lastColorUsed, so print nothing
					ansiColorCode = ""
					htmlColorCode = ""
				}

				// ~~~~~ SET THE COLOR OF THE PIXEL ~~~~~
				if ansiOrHtml == "ansi" {
					// fmt.Print(ansiColorCode)
					mainOutput += ansiColorCode
				} else if ansiOrHtml == "html" {
					// fmt.Print(htmlColorCode)
					mainOutput += htmlColorCode
				}
				lastColorUsed = colorToUse

			} // ~~~~~ END USE COLOR IF STATEMENT ~~~~~

			// correct intensity if needed
			if currentIntensity > maxIntensity {
				currentIntensity = maxIntensity
				// fmt.Println("fixed intensity:", currentIntensity) //DEBUG LINE
			}

			// ~~~~~ PRINT THE PIXEL ~~~~~
			// fmt.Println("about to use intensity:", currentIntensity) //DEBUG LINE
			mainOutput += intensityLevels[currentIntensity]
			if doubleWide { // print an extra time if dobule wide
				mainOutput += intensityLevels[currentIntensity]
			}
		}

		// ~~~~~ PRINT A NEW LINE AFTER EACH ROW ~~~~~
		if ansiOrHtml == "ansi" {
			// fmt.Print("\n")
			mainOutput += "\n"
		} else if ansiOrHtml == "html" {
			// fmt.Print("<br />")
			mainOutput += "<br />"
		}
	}

	// ~~~~~ RESET COLOR WHEN THE LOOP IS DONE ~~~~~
	if ansiOrHtml == "ansi" {
		// fmt.Print(ansiColorReset)
		mainOutput += ansiColorReset
	} else if ansiOrHtml == "html" {
		// fmt.Print(htmlColorEnd)
		mainOutput += htmlColorEnd
	}

	// ~~~~~ HANDLE THE OUTPUT ~~~~~
	if ansiOrHtml == "ansi" {
		fmt.Print(mainOutput)
	} else if ansiOrHtml == "html" {
		htmlcreator.WriteToHtmlFile(mainOutput, "go img-to-ascii output!!!", "")
	}

}

// ~~~~~ THE CHARACTER REPRESENTATIONS OF INCREASING LEVELS OF "BOLDNESS"/"INTENSITY" OF A PIXEL ~~~~~
func getIntensityLevelsSlice(num int) []string {
	var intensityLevels []string
	switch {
	case num == 0:
		intensityLevels = []string{"░", "▒", "▓", "█"}
	case num == 1:
		intensityLevels = []string{" ", "░", "▒", "▓", "█"}
	case num == 2:
		intensityLevels = []string{"▁", "░", "▒", "▓", "█"}
	case num == 3:
		intensityLevels = []string{" ", "-", "+", "#"}
	case num == 4:
		intensityLevels = []string{" ", ".", "▁", "▂", "▃", "▄", "▅", "▆", "▇", "▉", "█"}
	case num == 5:
		intensityLevels = []string{"_", "a", "b", "c", "d", "e", "f"}
	default:
		intensityLevels = []string{" ", "░", "▒", "▓", "█"}
	}
	return intensityLevels
}
