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
    #w10{width: calc(10 * var(--pixel-width));}
    #w11{width: calc(11 * var(--pixel-width));}
    #w12{width: calc(12 * var(--pixel-width));}
    #w13{width: calc(13 * var(--pixel-width));}
    #w14{width: calc(14 * var(--pixel-width));}
    #w15{width: calc(15 * var(--pixel-width));}
    #w16{width: calc(16 * var(--pixel-width));}
    #w17{width: calc(17 * var(--pixel-width));}
    #w18{width: calc(18 * var(--pixel-width));}
    #w19{width: calc(19 * var(--pixel-width));}
    #w20{width: calc(20 * var(--pixel-width));}
    #w21{width: calc(21 * var(--pixel-width));}
    #w22{width: calc(22 * var(--pixel-width));}
    #w23{width: calc(23 * var(--pixel-width));}
    #w24{width: calc(24 * var(--pixel-width));}
    #w25{width: calc(25 * var(--pixel-width));}
    #w26{width: calc(26 * var(--pixel-width));}
    #w27{width: calc(27 * var(--pixel-width));}
    #w28{width: calc(28 * var(--pixel-width));}
    #w29{width: calc(29 * var(--pixel-width));}
    #w30{width: calc(30 * var(--pixel-width));}
    #w31{width: calc(31 * var(--pixel-width));}
    #w32{width: calc(32 * var(--pixel-width));}
    #w33{width: calc(33 * var(--pixel-width));}
    #w34{width: calc(34 * var(--pixel-width));}
    #w35{width: calc(35 * var(--pixel-width));}
    #w36{width: calc(36 * var(--pixel-width));}
    #w37{width: calc(37 * var(--pixel-width));}
    #w38{width: calc(38 * var(--pixel-width));}
    #w39{width: calc(39 * var(--pixel-width));}
    #w40{width: calc(40 * var(--pixel-width));}
    #w41{width: calc(41 * var(--pixel-width));}
    #w42{width: calc(42 * var(--pixel-width));}
    #w43{width: calc(43 * var(--pixel-width));}
    #w44{width: calc(44 * var(--pixel-width));}
    #w45{width: calc(45 * var(--pixel-width));}
    #w46{width: calc(46 * var(--pixel-width));}
    #w47{width: calc(47 * var(--pixel-width));}
    #w48{width: calc(48 * var(--pixel-width));}
    #w49{width: calc(49 * var(--pixel-width));}
    #w50{width: calc(50 * var(--pixel-width));}
    #w51{width: calc(51 * var(--pixel-width));}
    #w52{width: calc(52 * var(--pixel-width));}
    #w53{width: calc(53 * var(--pixel-width));}
    #w54{width: calc(54 * var(--pixel-width));}
    #w55{width: calc(55 * var(--pixel-width));}
    #w56{width: calc(56 * var(--pixel-width));}
    #w57{width: calc(57 * var(--pixel-width));}
    #w58{width: calc(58 * var(--pixel-width));} 
    #w59{width: calc(59 * var(--pixel-width));}
    #w60{width: calc(60 * var(--pixel-width));}
    #w61{width: calc(61 * var(--pixel-width));}
    #w62{width: calc(62 * var(--pixel-width));}
    #w63{width: calc(63 * var(--pixel-width));}
    #w64{width: calc(64 * var(--pixel-width));}
    #w65{width: calc(65 * var(--pixel-width));}
    #w66{width: calc(66 * var(--pixel-width));}
    #w67{width: calc(67 * var(--pixel-width));}

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
