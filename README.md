[![GoDoc](https://godoc.org/github.com/StefanSchroeder/ttfsample?status.png)](https://godoc.org/github.com/StefanSchroeder/ttfsample)
[![Go Report Card](https://goreportcard.com/badge/github.com/StefanSchroeder/ttfsample)](https://goreportcard.com/report/github.com/StefanSchroeder/ttfsample)
 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# ttfsample
Creates a sample image of a Truetype font

ttfsample will take a TTF-file as an input and create a PNG-image 
with a sample of the font. 

For the License see LICENSE.

The program comes with a GNU Free Sans Bold True Type font which 
is under GNU Free Font license.

https://www.gnu.org/software/freefont/license.html

Usage:

    go run *.go -fontfile <path-to-your-ttf-font>
    ./ttfsample.exe -fontfile <path-to-your-ttf-font> (Windows)
    ./ttfsample -fontfile <path-to-your-ttf-font> (Linux)

There are a couple of options, primary being, that you can supply the text to be
printed as an argument. But there is also a sensible default (see image).

When run with the font Arial Narrow, the result will look like this:

![Sample](https://raw.githubusercontent.com/StefanSchroeder/ttfsample/master/sample/sample.png)

This will create an image file in the png/ directory which will 
created if it does not exist.

The name of the font will always be included printed with a
boring font, by default GNU FreeSansBold, that is always
readable even if the font has only symbols.

Author: Stefan Schröder, 2019

# Install

	go install github.com/StefanSchroeder/ttfsample/ttfsample@latest

# Build

	go build *.go 

will do the trick if your Go development environment is setup properly.

Tested on Windows and Linux.

# TODO

- OK Create output dir if missing.
- OK Do not look for Hack-font.
- OK Embed TTF properly.
- Support OTF.

