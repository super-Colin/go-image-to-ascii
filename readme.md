# 2024 Updatingzs!!!
## Documentation! Verbosity! Error Handling!
##### All soon to come (... hopefully..)



Flags:
- "-img"
- - The path to the image to modify
- "-v" 
- - verbose feedback as the process runs
- - bool : default=false
- "-cdr"
- - Color Distance Requirement - The distance requirement in hue between colors for them to be distinct
- - int: 0 - 255 : default : 40
- "-ccd" 
- - Color Close to Default - The distance requirement between colors for them to be close







scmg ( pathToImage, scaleToMaxWidth(# of "pixels"), nameOfFile )

Since we're outputting html that has to be interpretted by the browser and potentially sent online, size of the output is one of the top considerations.

Big thanks to all these resources! :
- https://golangbyexample.com/print-output-text-color-console/
- https://golangdocs.com/golang-image-processing

- https://pkg.go.dev/image/color
- https://cs.opensource.google/go/go/+/refs/tags/go1.17.3:src/image/color/color.go
- https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
- https://en.wikipedia.org/wiki/Block_Elements

