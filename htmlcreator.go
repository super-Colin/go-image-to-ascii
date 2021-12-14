package main

import (
	"fmt"
	"log"
	"os"
)

func generateCssForWidths(setOfWidths *Set) string {
	theWidthsCss := ""
	for width := range setOfWidths.list {
		theWidthsCss += fmt.Sprintf("#w%v{width: calc(%v * var(--pixel-width));}", width, width)
	}
	return theWidthsCss
}

func WriteToHtmlFile(htmlBodyToWrite, pageTitle, fileName, additionalCss string) {
	// set a name for the file
	fName := ""
	if fileName == "" {
		fName = "index"
	}
	// create / open the file
	file, err := os.Create("./htmlCreatorOutput-" + fName + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// print to the file
	fmt.Fprintln(file, WrapWithBoilerplateHTML(htmlBodyToWrite, pageTitle, additionalCss))
}

func WrapWithBoilerplateHTML(htmlBodyToWrap, pageTitle, additionalCss string) string {
	return `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>` + pageTitle + `</title>
<style>
	body{
			--pixel-width: 2px; /* width of a sqaure pixel in the image */
	}
	.pR{ /* .pixelImg_row */
			display: flex;
			align-items: center;
	}
	.pR > *{ /* .pixelImg_row */
			display: block;
			min-width: var(--pixel-width);
			height: var(--pixel-width);
	}
	/* something like this to extend pixels, would like to make these programatically and only for the widths that get used */
    ` + additionalCss + `

	/* Color map of custom elements, avoiding all established 1 letter element tags: https://www.javatpoint.com/html-tags */
	r{background-color:#f00;}
	n{background-color:#0f0;}
	l{background-color:#00f;}
	y{background-color:#ff0;}
	c{background-color:#0ff;}
	m{background-color:#f0f;}
	k{background-color:#000;}
	w{background-color:#fff;}
	/* Contemplated 2 letter tags but that results in ~30% more text (filesize!) overall */
</style>
</head>
<body>
<div class="pixelImg_container">
` + htmlBodyToWrap + `
</div>
</body>
</html>
`
}
