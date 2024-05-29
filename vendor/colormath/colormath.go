package colormath

import (
	"math"
)

// "image/color"

func PickColor(red, green, blue, alpha uint8, colorDistanceRequirement, colorCloseToRequirement int) string {
	colorToReturn := ""

	rCloseToB := withinRangeOf(red, blue, colorCloseToRequirement)
	rCloseToG := withinRangeOf(red, green, colorCloseToRequirement)
	gCloseToB := withinRangeOf(green, blue, colorCloseToRequirement)

	rgbMax := maxUint8(red, green, blue)
	rgbMin := minUint8(red, green, blue)

	// CHECK FOR OVERALL BRIGHTNESS / DARKNESS
	switch {
	// no color is very bright
	case rgbMax < 40:
		colorToReturn = "black"
	// all colors are bright
	case rgbMin > 230:
		colorToReturn = "white"
	// red and blue are close, green isn't
	case rgbMin == green && rCloseToB && !rCloseToG:
		colorToReturn = "purple"
	// red and green are close, blue isn't
	case rgbMin == blue && rCloseToG && !rCloseToB:
		colorToReturn = "yellow"
	// green and blue are close, red isn't
	case rgbMin == red && gCloseToB && !rCloseToG:
		colorToReturn = "cyan"
	// red is dominant
	case rgbMax == red && !rCloseToG && !rCloseToB:
		colorToReturn = "red"
	// green is dominant
	case rgbMax == green && !rCloseToG && !gCloseToB:
		colorToReturn = "green"
	// blue is dominant
	case rgbMax == blue && !gCloseToB && !rCloseToB:
		colorToReturn = "blue"
	default:
		// fmt.Println("switch chose default")//DEBUG LINE
	}

	return colorToReturn
}

func DecideIntensityWithGrayscale(perLevelThreshold, maxIntensity int) int {
	return 1
}
func DecideIntensityWithColor(red, green, blue, alpha uint8, perLevelThreshold, maxIntensity int) int {
	brightness := math.Sqrt(float64(0.299*math.Pow(float64(red), 2) + 0.587*math.Pow(float64(green), 2) + 0.114*math.Pow(float64(blue), 2)))
	// ^ will return a value between 0 and 255
	returnIntensity := int(int(brightness) / perLevelThreshold)
	if returnIntensity > maxIntensity {
		returnIntensity = maxIntensity
	}
	// fmt.Println("returnIntensity:", returnIntensity) //DEBUG LINE
	return returnIntensity
}

func distanceFrom(a, b uint8) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
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
