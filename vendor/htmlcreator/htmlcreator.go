package htmlcreator

import (
	"fmt"
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

func WriteToHtmlFile(fileName string, htmlBodyToWrite string, pageTitle string) {
	// set a name for the file
	fName := ""
	if fileName == "" {
		fName = "index.html"
	}
	// create / open the file
	file, err := os.Create("./htmlCreatorOutput/" + fName + ".html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// print to the file
	fmt.Fprintln(file, WrapWithBoilerplateHTML(htmlBodyToWrite, pageTitle))
}
