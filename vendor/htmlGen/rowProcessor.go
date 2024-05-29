package htmlGen

import (
	"colormath"
	"fmt"
	"image"
	"image/color"
)

type processedRow struct {
	rowHtml        string
	rowNumber      int
	widthsRequired map[int]struct{} // Set of ints, int keys with empty vals
}

func rowProcessor(c chan processedRow, imagePointer image.Image, yCoord, scaleBy, colorDistanceRequirement, colorCloseToRequirement int) {
	bounds := imagePointer.Bounds()
	x, maxX := bounds.Min.X, bounds.Max.X //start at min, iterate to max

	rowOutput := "<span class=\"pR\">"
	lastColorUsed := ""
	widthsRequired := NewSet()
	var pixelWidth int = 1 // start as 1x1 square vs 2x1, 14x1 etc..
	// ~~~~~ START A LOOP THROUGH THE ROW ~~~~~
	for ; x < maxX; x += scaleBy {
		convertedPixel := color.RGBAModel.Convert(imagePointer.At(x, yCoord)).(color.RGBA)
		pRed := convertedPixel.R
		pGreen := convertedPixel.G
		pBlue := convertedPixel.B
		pAlpha := convertedPixel.A

		// ~~~~~ PICK COLOR ~~~~~
		colorToUse := colormath.PickColor(pRed, pGreen, pBlue, pAlpha, colorDistanceRequirement, colorCloseToRequirement)
		//if the same color as last pixel in row, just make this pixel 1 width wider
		if lastColorUsed == colorToUse {
			pixelWidth++
		} else { //if we're at a new color add the last pixel with it's width
			elemLetter := colorToShortElemLetter(lastColorUsed)
			pixelToAdd := "<" + elemLetter + ">" + "</" + elemLetter + ">"
			if pixelWidth > 1 { // if its wide add a class(..id) to lengthen it
				pixelToAdd = "<" + elemLetter + " id=\"w" + fmt.Sprintf("%d", pixelWidth) + "\"></" + elemLetter + ">"
			}

			rowOutput += pixelToAdd //add the pixel to the output

			widthsRequired.Add(pixelWidth)
			pixelWidth = 1             //reset pixel width
			lastColorUsed = colorToUse // reset color
		}

	}
	widthsRequired.Add(pixelWidth)
	elemLetter := colorToShortElemLetter(lastColorUsed)
	pixelToAdd := "<" + elemLetter + ">" + "</" + elemLetter + ">"
	if pixelWidth > 1 { // if its wide add a class(..id) to lengthen it
		pixelToAdd = "<" + elemLetter + " id=\"w" + fmt.Sprintf("%d", pixelWidth) + "\"></" + elemLetter + ">"
	}

	rowOutput += pixelToAdd //add the pixel to the output
	rowOutput += "</span>"
	theRow := processedRow{
		rowHtml:        rowOutput,
		rowNumber:      yCoord,
		widthsRequired: widthsRequired, //basically a set
	}

	// ~~~~~ HANDLE THE OUTPUT ~~~~~
	c <- theRow //send the value through the pipe <p></p>
}

// THESE ARE ALL MAPPED TO CSS FOR COLORS
func colorToShortElemLetter(color string) string {
	theLetter := ""
	// fmt.Println("check elem to use for color: ", color)
	switch {
	case color == "black":
		theLetter = "k" // <k></k>
	case color == "white":
		theLetter = "w" //<w></w>
	case color == "purple":
		theLetter = "m" // ..
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

func recombineRows(finishedRows []processedRow) string {
	orderedRows := make(map[int]string)
	theOut := ""
	// ~~~~~ PUTTING THEM INTO theOut WILL PUT THEM IN THE CORRECT ORDER ~~~~~
	for _, theRow := range finishedRows {
		orderedRows[theRow.rowNumber] = theRow.rowHtml
	}
	for _, theRowHtml := range orderedRows {
		theOut += theRowHtml
	}
	return theOut
}

// func getRequiredWidths(finishedRows []processedRow) map[int]struct{} {
func getRequiredWidths(finishedRows []processedRow) *Set {
	// theSet := make(map[int]struct{})
	theSet := NewSet()
	for _, theRow := range finishedRows {
		// theSet[theRow.rowNumber] = struct{}{}
		theSet.Add(theRow.rowNumber)
	}

	return theSet
}
func getAllRequiredWidths(rows ...processedRow) *Set {
	theSet := NewSet()
	for _, row := range rows { //for each row
		for key := range row.widthsRequired.list { // for each list of widths
			theSet.Add(key)
		}
	}
	return theSet
}
