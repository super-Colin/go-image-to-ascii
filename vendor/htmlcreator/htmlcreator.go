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
