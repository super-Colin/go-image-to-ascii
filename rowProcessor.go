package main

import "image"

func rowProcessor(c chan string, imagePointer image.Image, Ycoord, scaleBy int) {
	c <- "d"

	// ~~~~~ PICK INTENSITY ~~~~~

	// ~~~~~ PRINT THE PIXEL ~~~~~

	// ~~~~~ PRINT A NEW LINE AFTER EACH ROW ~~~~~

	// ~~~~~ RESET COLOR WHEN THE LOOP IS DONE ~~~~~

	// ~~~~~ HANDLE THE OUTPUT ~~~~~

}
