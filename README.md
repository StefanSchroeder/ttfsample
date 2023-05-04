[![GoDoc](https://godoc.org/github.com/StefanSchroeder/ttfsample?status.png)](https://godoc.org/github.com/StefanSchroeder/ttfsample)
[![Go Report Card](https://goreportcard.com/badge/github.com/StefanSchroeder/ttfsample)](https://goreportcard.com/report/github.com/StefanSchroeder/ttfsample)
 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# ttfsample

Creates a sample image of a Truetype TTF font or Opentype OTF font.

ttfsample will take a font file as an input and create a PNG-image 
with a sample of the font. 

For the License see LICENSE.

The program comes with a GNU Free Sans and Serif Bold True Type font which 
is under GNU Free Font license.

https://www.gnu.org/software/freefont/license.html

There are a couple of options, primary being, that you can supply the text to be
printed as an argument. But there is also a sensible default (see image).

When run with the font Arial Narrow, the result will look like this:

![Sample](https://raw.githubusercontent.com/StefanSchroeder/ttfsample/master/sample/sample.png)

This will create an image file in the png/ directory which will 
be created if it does not exist.

The name of the font will always be included printed with a
boring font, by default GNU FreeSansBold, that is always
readable even if the font has only symbols.

Author: Stefan Schröder, 2019, 2023

# Install and quickstart

	go install github.com/StefanSchroeder/ttfsample/ttfsample@latest

	ttfsample -fontfile somefont.ttf

will create a PNG image in the newly created directory *png/*.

# Build

	go build . 

will do the trick if your Go development environment is setup properly.

# Options

	-fontfile path/to/font.ttf

This is a mandatory option. Provide the path to the font, TTF or
OTF. You can only process one font.

	-hinting <none|full>

Set *hinting* to *none* to disable hinting. Default is *full*.

	-dpi INTEGER

Default is 72. Set dots per inch.

	-outdir

The output directory where the image will be stored.

	-size INTEGER

Font size in points. Default is 100. Use responsibly.

	-boringfont path/to/font

The font that is used to print the title of the font. This
must be a TTF font.

	-spacing FLOAT

Defaults to 1.5. Distance between two lines. 

# Testing

Tested on Windows and Linux.

Run

	go test

