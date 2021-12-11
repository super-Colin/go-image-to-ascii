package htmlcreator

import (
	"fmt"
	"log"
	"os"
)

func WrapWithBoilerplateHTML(htmlBodyToWrap string, pageTitle string) string {
	return `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>` + pageTitle + `</title>
<style>
	body{
			--pixel-width: 12px; /* width of a sqaure pixel in the image */
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
	#w2{width: calc(2 * var(--pixel-width));}
	#w3{width: calc(3 * var(--pixel-width));}
	#w4{width: calc(4 * var(--pixel-width));}
	#w5{width: calc(5 * var(--pixel-width));}
	#w6{width: calc(6 * var(--pixel-width));}
	#w7{width: calc(7 * var(--pixel-width));}
	#w8{width: calc(8 * var(--pixel-width));}
	#w9{width: calc(9 * var(--pixel-width));}
	#wa{width: calc(10 * var(--pixel-width));}
	#wb{width: calc(11 * var(--pixel-width));}
	#wc{width: calc(12 * var(--pixel-width));}
	#wd{width: calc(13 * var(--pixel-width));}
	#we{width: calc(14 * var(--pixel-width));}
	#wf{width: calc(15 * var(--pixel-width));}
	/* ........ */
	#wx{width: calc(33 * var(--pixel-width));}
	#wy{width: calc(34 * var(--pixel-width));}
	#wz{width: calc(35 * var(--pixel-width));}

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
` + htmlBodyToWrap + `
</body>
</html>
`
}

func WriteToHtmlFile(htmlBodyToWrite string, pageTitle string, fileName string) {
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
	fmt.Fprintln(file, WrapWithBoilerplateHTML(htmlBodyToWrite, pageTitle))
}
