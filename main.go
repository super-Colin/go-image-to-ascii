package main

import (
	"flag"
	"fmt"
	"image"

	// _ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"sort"
)

var imagePath string
var colorDistanceRequirement int
var colorCloseToRequirement int
var maxPixelWidth int
var scaleFactor float64

func main() {

	// Verify args and set them as globals
	parseArgs()

	// ~~~~~ GET IMAGE ~~~~~
	imagePointer, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("error opening image", err)
	}
	defer imagePointer.Close() // go ahead and defer closing the file for when program is done

	decodedImage, _, err := image.Decode(imagePointer)
	if err != nil {
		log.Fatal(err)
	}

	// ~~~~~ FIGURE OUT THE EXACT WIDTH ~~~~~
	imgBounds := decodedImage.Bounds()
	scaleDownBy := imgBounds.Max.X / maxPixelWidth
	calcWidth(imgBounds.Max.X)

	// ~~~~~ TRANSFROM IMAGE BINARY ~~~~~

	// ~~~~~ DECLARE GLOBALS FOR THE LOOP ~~~~~
	var rowWorkerChannels []<-chan processedRow
	var rowsTotal int = 0

	// ~~~~~ START THE LOOP! ~~~~~
	// FOR EACH ROW START A WORKER AND ADD THE CHANNEL TO THE POOL
	for y := decodedImage.Bounds().Min.Y; y < decodedImage.Bounds().Max.Y; y += scaleDownBy {
		rowWorkerChannels = append(rowWorkerChannels, rowWorker(decodedImage, y, scaleDownBy, colorDistanceRequirement, colorCloseToRequirement))
		rowsTotal++
	}

	// ~~~~~ COMBINE THE WORKER CHANNELS ~~~~~
	workerOutputChannel := FanInProcessedRows(rowWorkerChannels...)

	// ~~~~~ RECIEVE THE PROCESSED ROWS  ~~~~~
	var finishedRows []processedRow
	for i := 0; i < rowsTotal; i++ { // this loop will block until it has recieved all rows back
		theProcessedRow := <-workerOutputChannel //blocks to listen
		finishedRows = append(finishedRows, theProcessedRow)
	}
	// ~~~~~ PUT THEM BACK IN ORDER ~~~~~
	theWidths := getAllRequiredWidths(finishedRows...)
	// create css for all the necessary widths
	theCss := generateCssForWidths(theWidths)
	theOut := reorderRows(finishedRows)

	// ~~~~~ HANDLE REASSEMBLED OUTPUT ~~~~~
	WriteToHtmlFile(theOut, "go img-to-ascii output!!!", "", theCss)
}

//

//

func calcWidth(imageBounds_maxX int) int {
	// ~~~~~ SET SCALE FACTOR & MAX WIDTH ~~~~~
	if scaleFactor != 0 || maxPixelWidth != 0 {
		if scaleFactor > 0 && scaleFactor < 1 {
			scaledWidth = imageBounds_maxX * scaleFactor
		}
		if maxPixelWidth != 0 {
			maxPixelWidth = imageBounds_maxX / 2
		}
		maxPixelWidth = imageBounds_maxX / 2

	}
	return 0
}

//

// Returns a channel that the rowProcessor will return it's value through
func rowWorker(imagePointer image.Image, Ycoord, scaleBy, colorDistanceRequirement, colorCloseToRequirement int) <-chan processedRow {
	c := make(chan processedRow)
	go rowProcessor(c, imagePointer, Ycoord, scaleBy, colorDistanceRequirement, colorCloseToRequirement)
	return c
}

//

// Will take an amount of channels, listen to them all and return all results through a single out channel
func FanInProcessedRows(chans ...<-chan processedRow) <-chan processedRow {
	newOutChannel := make(chan processedRow)
	for _, channelIn := range chans {
		go func(cOut chan processedRow, cIn <-chan processedRow) {
			for val := range cIn {
				cOut <- val
			}
		}(newOutChannel, channelIn)
	}
	return newOutChannel
}

//

func reorderRows(processedRows []processedRow) string {
	// theMap := make(map[int]processedRow)
	var theSlice []processedRow
	var mainOutput string
	for _, theRow := range processedRows {
		// theMap[theRow.rowNumber] = theRow
		theSlice = append(theSlice, theRow)
	}
	// fmt.Println("about to output this map:", theMap)
	// for _, theRow := range theMap {
	sort.Slice(theSlice, func(x, n int) bool { return theSlice[x].rowNumber < theSlice[n].rowNumber })
	// fmt.Println(theSlice)

	for _, theRow := range theSlice {
		// fmt.Println("about to output this row from the map:", theRow)
		mainOutput += theRow.rowHtml
	}

	return mainOutput
}

// Verify flag arguments
func parseArgs() {
	// Defaults
	var colorDistanceDefault int = 40
	var colorCloseToDefault int = 80

	// ~~~~~ DEFINE FLAGS ~~~~~

	// Image Path - the image to use
	var arg_imagePath string
	flag.StringVar(&arg_imagePath, "img", "", "The path to the image you want to process")

	// Max Width - default 0/infinite
	var arg_maxWidth int
	flag.IntVar(&arg_maxWidth, "mw", 0, "The maximum width of the image you want to process")

	// Scale - default 0/infinite
	var arg_scaleFactor float64
	flag.Float64Var(&arg_scaleFactor, "scale", 0, "Whether to scale the image, must be less than 1, 0=No scaling")

	// Color Distance - How far to be independant colors
	var arg_colorDistanceReq int
	flag.IntVar(&arg_colorDistanceReq, "cdr", colorDistanceDefault, "The distance requirement between colors for them to be distinct")
	// Verirfy
	if arg_colorDistanceReq < 0 || arg_colorDistanceReq > 255 {
		fmt.Println("Color distance requirement must be between 0 and 255, defaulting to ", colorDistanceDefault)
		arg_colorDistanceReq = colorDistanceDefault
	}

	// Color Closeness - How close to create mixed colors
	var arg_colorCloseToReq int
	flag.IntVar(&arg_colorCloseToReq, "ccd", colorCloseToDefault, "The distance requirement between colors for them to be close")
	// Verirfy
	if arg_colorCloseToReq < 0 || arg_colorCloseToReq > 255 {
		fmt.Println("Color close to requirement  must be between 0 and 255, defaulting to ", colorCloseToDefault)
		arg_colorCloseToReq = colorCloseToDefault
	}

	// Parse Flag Args
	flag.Parse()

	if arg_imagePath == "" {
		log.Fatal("Empty image path provided")
	}

	imagePath = arg_imagePath
	colorDistanceRequirement = arg_colorDistanceReq
	colorCloseToRequirement = arg_colorCloseToReq
	maxPixelWidth = arg_maxWidth
	scaleFactor = arg_scaleFactor
}
