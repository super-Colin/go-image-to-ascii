package pickcolor

// "image/color"

// func PickColor(red, green, blue, alpha uint8, colorDistanceRequirement, colorCloseToRequirement int, lastColorUsed string) string {
func PickColor(red, green, blue, alpha uint8, colorDistanceRequirement, colorCloseToRequirement int, lastColorUsed string) string {

	colorToReturn := ""
	rCloseToB := withinRangeOf(red, blue, colorCloseToRequirement)
	rCloseToG := withinRangeOf(red, green, colorCloseToRequirement)
	gCloseToB := withinRangeOf(green, blue, colorCloseToRequirement)

	rgbMax := maxUint8(red, green, blue)
	rgbMin := minUint8(red, green, blue)

	// -- debug line --
	// fmt.Println(ansiColorWhite, "Max, Min // R G B ", rgbMax, rgbMin, "//", ansiColorRed, red, ansiColorGreen, green, ansiColorBlue, blue)

	switch {
	// no color is very bright
	case lastColorUsed != "black" && rgbMax < 80:
		colorToReturn = "black"

	// all colors are bright
	case lastColorUsed != "white" && rgbMin > 200:
		colorToReturn = "white"

	// red and blue are close, green isn't
	case lastColorUsed != "purple" && rgbMin == green && rCloseToB && !rCloseToG:
		colorToReturn = "purple"

	// red and green are close, blue isn't
	case lastColorUsed != "yellow" && rgbMin == blue && rCloseToG && !rCloseToB:
		colorToReturn = "yellow"

	// green and blue are close, red isn't
	case lastColorUsed != "cyan" && rgbMin == red && gCloseToB && !rCloseToG:
		colorToReturn = "cyan"

	// red is dominant
	case lastColorUsed != "red" && rgbMax == red && !rCloseToG && !rCloseToB:
		colorToReturn = "red"
	// green is dominant
	case lastColorUsed != "green" && rgbMax == green && !rCloseToG && !gCloseToB:
		colorToReturn = "green"
	// blue is dominant
	case lastColorUsed != "blue" && rgbMax == blue && !gCloseToB && !rCloseToB:
		colorToReturn = "blue"
	}

	// colorStrength := color.GrayModel.Convert(decodedImage.At(x, y)).(color.Gray)
	// // colorStrength := rgbMin
	// level := int(colorStrength.Y) / valPerLevel
	// // level := int(colorStrength) / valPerLevel

	return colorToReturn
}

// func PickColor(r, g, b, a uint8) uint32 {
// 	return uint32(r) | uint32(g)<<8 | uint32(b)<<16 | uint32(a)<<24
// }

func decideIntensity() int {
	return 1
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
