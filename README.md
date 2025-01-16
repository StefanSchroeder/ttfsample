[![GoDoc](https://godoc.org/github.com/StefanSchroeder/ttfsample?status.png)](https://godoc.org/github.com/StefanSchroeder/ttfsample)
[![Go Report Card](https://goreportcard.com/badge/github.com/StefanSchroeder/ttfsample)](https://goreportcard.com/report/github.com/StefanSchroeder/ttfsample)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/StefanSchroeder/ttfsample/badge)](https://scorecard.dev/viewer/?uri=github.com/StefanSchroeder/ttfsample)
[![Go Build](https://github.com/StefanSchroeder/ttfsample/actions/workflows/go.yml/badge.svg)](https://github.com/StefanSchroeder/ttfsample/actions/workflows/go.yml)
[![Coverage](https://github.com/StefanSchroeder/ttfsample/actions/workflows/codecov.yml/badge.svg)](https://github.com/StefanSchroeder/ttfsample/actions/workflows/codecov.yml)

# ttfsample

*ttfsample* is a small utility to create a sample image of a Truetype TTF font or Opentype OTF font.

For the License see [LICENSE](LICENSE)

The program comes with a GNU Free Sans and Serif Bold True Type font which 
are under the [GNU Free Font license](https://www.gnu.org/software/freefont/license.html).

There are a couple of options, primary being, that you can supply the text to be
printed as an argument. But there is also a sensible default (see image).

When run with the font Arial Narrow, the result will look like this:

![Sample](https://raw.githubusercontent.com/StefanSchroeder/ttfsample/master/sample/sample.png)

The name of the font will always be included, printed with a
boring font, GNU FreeSansBold, that is always
readable even if the font has only symbols.

Author: Stefan Schröder, 2019 - 2024

# Changelog

0.4.0: we have changed the options. Now the font files are to
be used as parameters, whereas the desired output string is the
new option *-wanted*. This feels more natural.

# Install and quickstart

	go install github.com/StefanSchroeder/ttfsample/ttfsample@latest

	ttfsample somefont.ttf someotherfont.ttf

will create a PNG image in the newly created directory *png/*.

# Build

	go build . 

will do the trick if your [Go development environment is setup properly](https://go.dev/doc/install).

# Parameters

	ttfsample somefont.ttf someotherfont.ttf

This command will create two sample files, one for each font in
the *png* directory, which will be created for you.

# Options

    -wanted "Hello font"
    -wanted "First line\nSecond line"

Print the text _Hello font_ on the canvas instead of the default
alphabet. You can use *\n* to insert a newline.

	-hinting <none|full>

Set *hinting* to *none* to disable hinting. Default is *full*.

	-dpi INTEGER

The default value for *dpi* is 72. Set dots per inch.

	-outdir STRING

The *outdir* option defines the output directory where the image will be stored.

	-size INTEGER

*-size* sets the font size in points. Default is 100. Use responsibly.

	-spacing FLOAT

*-spacing* set the distance between two lines. Defaults to 1.5.

	-walk PATH

The *-walk* option recursively searches the directory tree for fonts to print
starting from PATH and not following symlinks.

    -width INTEGER

The *-width* option defines the width of the generated PNG image.
The default is 2000. If you make the width too small, the image
will be simply cut off. Setting this value to 0 is illegal.
The absolute value will taken, so you can use negative numbers, 
but why would you do that?

# Testing

Tested on Windows and Linux.

Run

	go test

