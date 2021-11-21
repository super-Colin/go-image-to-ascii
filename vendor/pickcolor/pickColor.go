package pickcolor

import (
	"fmt"
	"image/color"
)

func PickColor(r, g, b, a uint8) uint32 {
	// decide what color to use
	red := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).R
	green := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).G
	blue := color.RGBAModel.Convert(decodedImage.At(x, y)).(color.RGBA).B
	rCloseToB := withinRangeOf(red, blue, colorDistance)
	rCloseToG := withinRangeOf(red, green, colorDistance)
	gCloseToB := withinRangeOf(green, blue, colorDistance)

	rgbMax := maxUint8(red, green, blue)
	rgbMin := minUint8(red, green, blue)

	colorStrength := color.GrayModel.Convert(decodedImage.At(x, y)).(color.Gray)
	// colorStrength := rgbMin
	level := int(colorStrength.Y) / valPerLevel
	// level := int(colorStrength) / valPerLevel

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
}

// func PickColor(r, g, b, a uint8) uint32 {
// 	return uint32(r) | uint32(g)<<8 | uint32(b)<<16 | uint32(a)<<24
// }
