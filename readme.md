
Desired results:
A CLI tool that can take and image and create an HTML representation of it

C:/xampp> imgtoahtml ( pathToImage, scaleToMaxWidth(# of "pixels"), nameOfFile )

Since we're outputting html that has to be interpretted by the browser and potentially sent online, size of the output is one of the top considerations.

resources:
- https://golangbyexample.com/print-output-text-color-console/
- https://golangdocs.com/golang-image-processing

- https://pkg.go.dev/image/color
- https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/image/color/color.go
- https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
- https://en.wikipedia.org/wiki/Block_Elements
