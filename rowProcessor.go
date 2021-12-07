package main

import (
	"image"
	"image/color"
	"pickcolor"
)

type processedRow struct {
	outputHtml string
	rowNumber  int
}

func rowProcessor(c chan string, imagePointer image.Image, yCoord, scaleBy, colorDistanceRequirement, colorCloseToRequirement int) {
	bounds := imagePointer.Bounds()
	x, maxX := bounds.Min.X, bounds.Max.X //start at min, iterate to max

	mainOutput := "<span class=\"pR\">"
	lastColorUsed := ""
	var pixelWidth int = 1 // 1x1 square
	// ~~~~~ START A LOOP THROUGH THE ROW ~~~~~
	for ; x < maxX; x += scaleBy {
		convertedPixel := color.RGBAModel.Convert(imagePointer.At(1, yCoord)).(color.RGBA)
		pRed := convertedPixel.R
		pGreen := convertedPixel.G
		pBlue := convertedPixel.B
		pAlpha := convertedPixel.A

		// ~~~~~ PICK COLOR ~~~~~
		colorToUse := pickcolor.PickColor(pRed, pGreen, pBlue, pAlpha, colorDistanceRequirement, colorCloseToRequirement)

		// ~~~~~ SET COLOR CODE  ~~~~~
		if lastColorUsed == colorToUse { //if the same color as last pixel in row, just make this pixel 1 width wider
			pixelWidth++
		} else { //if we're at a new color add the last pixel with it's width
			elemLetter := colorToShortElemLetter(lastColorUsed)
			pixelToAdd := "<" + elemLetter + ">" + "</" + elemLetter + ">"
			if pixelWidth > 1 {

				pixelToAdd = "<" + elemLetter + " id=\"w" + colorToUse + "\"></" + elemLetter + ">"
			}
			mainOutput += pixelToAdd
			pixelWidth = 1
		}

	}

	// ~~~~~ HANDLE THE OUTPUT ~~~~~
	c <- "d" //send the value through the pipe <p></p>
}

// THESE ARE ALL MAPPED TO CSS FOR COLORS
func colorToShortElemLetter(color string) string {
	theLetter := ""
	switch {
	case color == "black":
		theLetter = "k"
	case color == "white":
		theLetter = "w"
	case color == "purple":
		theLetter = "m"
	case color == "yellow":
		theLetter = "y"
	case color == "cyan":
		theLetter = "c"
	case color == "red":
		theLetter = "r"
	case color == "green":
		theLetter = "n"
	case color == "blue":
		theLetter = "l"
	default:
		theLetter = "span"
	}
	return theLetter
}
