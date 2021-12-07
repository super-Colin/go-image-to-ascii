package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {

	// ~~~~~ GLOBAL SETTINGS ~~~~~

	// colorDistanceRequirement := 80 // 0-255; Distance between colors for them to be distinct
	// colorCloseToRequirement := 60  // 0-255;  Distance ... to be close to each other, for blending to secondary colors

	maxPixelWidth := 100

	doubleWide := true // Useful for monospace fonts, each pixel on the x-axis is printed twice
	// doubleWide = false

	// Double wide respects the maxPixelWidth setting
	if doubleWide {
		maxPixelWidth /= 2 //half the width of pixels actually desired since we will print them all twice
	}

	// ~~~~~ GET ASCII REPRESENTATION OF PIXEL INTENSITY ~~~~~
	// intensityLevels := getIntensityLevelsSlice(0)

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
	defer imagePointer.Close() //defer to close the file when program is done

	decodedImage, _, err := image.Decode(imagePointer)
	if err != nil {
		log.Fatal(err)
	}

	// ~~~~~ CREATE COLOR MAPS ~~~~~

	// ~~~~~ CALC THE INTENSITY INCREASE REQUIRED FOR EACH CHARACTER IN intensityLevels ~~~~~
	// var valPerIntensityLevel int = 255 / len(intensityLevels)
	// maxIntensity, currentIntensity := len(intensityLevels)-1, 0
	// Make our loop iterator jump by this factor to scale down the image. The same value works for both axis
	scaleDownBy := decodedImage.Bounds().Max.X / maxPixelWidth

	// ~~~~~ DECLARE GLOBALS FOR THE LOOP ~~~~~
	// mainOutput, htmlColorCode, ansiColorCode, lastColorUsed := "", "", "", ""

	// ~~~~~ START THE LOOP! ~~~~~
	// For each row we'll create a worker and pool all the returned channels together
	var workerChannels []<-chan string
	for y := decodedImage.Bounds().Min.Y; y < decodedImage.Bounds().Max.Y; y += scaleDownBy {
		workerChannels = append(workerChannels, rowWorker(decodedImage, y, scaleDownBy))
	}

}

// Returns a channel that the rowProcessor will return it's value through
func rowWorker(imagePointer image.Image, Ycoord, scaleBy int) <-chan string {
	c := make(chan string)
	go rowProcessor(c, imagePointer, Ycoord, scaleBy)
	return c
}

func fanInStringChannels(channelsIn ...<-chan string) <-chan string {
	channelOut := make(chan string)
	for _, channelIn := range channelsIn {
		go func() {
			channelOut <- <-channelIn
		}()
	}
	return channelOut
}

// ~~~~~ THE CHARACTER REPRESENTATIONS OF INCREASING LEVELS OF "BOLDNESS"/"INTENSITY" OF A PIXEL ~~~~~
// ~~~~~ Grouped together for clarity ~~~~~
func getIntensityLevelsSlice22(num int) []string {
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
