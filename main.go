package main

import (
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

	colorDistanceRequirement := 40 // 0-255; Distance between colors for them to be distinct
	colorCloseToRequirement := 20  // 0-255;  Distance ... to be close to each other, for blending to secondary colors

	maxPixelWidth := 400

	// ~~~~~ GET IMAGE ~~~~~

	imagePointer, err := os.Open("C:\\zHolderFolder\\cat1.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\windowPainting.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\triangle.png")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\ColinPicture3.jpg")
	// imagePointer, err := os.Open("C:\\zHolderFolder\\Frame16_2.png")
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

	scaleDownBy := decodedImage.Bounds().Max.X / maxPixelWidth

	// ~~~~~ DECLARE GLOBALS FOR THE LOOP ~~~~~
	var rowWorkerChannels []<-chan processedRow
	var rowsTotal int = 0

	// ~~~~~ START THE LOOP! ~~~~~
	// FOR EACH ROW START A WORKER AND ADD THE CHANNEL TO POOL
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
	// create css for all the necesarry
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
