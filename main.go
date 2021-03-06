package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"sort"
)

func main() {

	// ~~~~~ GLOBAL SETTINGS ~~~~~
	var colorDistanceReqArg int
	colorDistanceFlag := "cdr"
	colorDistanceDefault := 40
	colorDistanceDescription := "The distance requirement between colors for them to be distinct"
	colorDistanceMin := 0
	colorDistanceMax := 255
	flag.IntVar(&colorDistanceReqArg, colorDistanceFlag, colorDistanceDefault, colorDistanceDescription)
	if colorDistanceReqArg < colorDistanceMin || colorDistanceReqArg > colorDistanceMax {
		fmt.Println("Color distance requirement must be between 0 and 255, defaulting to ", colorDistanceDefault)
		colorDistanceReqArg = colorDistanceDefault
	}

	var colorCloseToReqArg int
	colorCloseToFlag := "cctr"
	colorCloseToDefault := 80
	colorCloseToDescription := "The distance requirement between colors for them to be close"
	colorCloseToMin := 0
	colorCloseToMax := 255
	flag.IntVar(&colorCloseToReqArg, colorCloseToFlag, colorCloseToDefault, colorCloseToDescription)
	if colorCloseToReqArg < colorCloseToMin || colorCloseToReqArg > colorCloseToMax {
		fmt.Println("Color close to requirement  must be between 0 and 255, defaulting to ", colorCloseToDefault)
		colorCloseToReqArg = colorCloseToDefault
	}

	// ~~~~~ PARSE ARGS ~~~~~
	var imagePathArg string
	var maxWidthArg int
	flag.StringVar(&imagePathArg, "image", "", "The path to the image you want to process")
	flag.IntVar(&maxWidthArg, "mw", 0, "The maximum width of the image you want to process")

	flag.Parse()

	colorDistanceRequirement := colorDistanceReqArg
	colorCloseToRequirement := colorCloseToReqArg
	maxPixelWidth := maxWidthArg

	// ~~~~~ VALIDATE ARGS ~~~~~

	// ~~~~~ GET IMAGE ~~~~~

	// imagePointer, err := os.Open("C:\\zHolderFolder\\color-wheel.png")
	imagePointer, err := os.Open(imagePathArg)
	if err != nil {
		log.Fatal("error opening image", err)
	}
	defer imagePointer.Close() //defer to close the file when program is done

	decodedImage, _, err := image.Decode(imagePointer)
	if err != nil {
		log.Fatal(err)
	}

	if maxPixelWidth == 0 {
		maxPixelWidth = decodedImage.Bounds().Max.X / 2
	}
	scaleDownBy := decodedImage.Bounds().Max.X / maxPixelWidth

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

// Returns a channel that the rowProcessor will return it's value through
func rowWorker(imagePointer image.Image, Ycoord, scaleBy, colorDistanceRequirement, colorCloseToRequirement int) <-chan processedRow {
	c := make(chan processedRow)
	go rowProcessor(c, imagePointer, Ycoord, scaleBy, colorDistanceRequirement, colorCloseToRequirement)
	return c
}

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
