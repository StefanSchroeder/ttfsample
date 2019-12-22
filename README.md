# ttfsample
Creates a sample image of a Truetype font

ttfsample will take a TTF-file as an input and create a PNG-image 
with a sample of the font. 

For the License see LICENSE.

The program comes with a GNU Free Sans Bold True Type font which 
is under GNU Free Font license.

https://www.gnu.org/software/freefont/license.html

Usage:

    go run ttfsample.go -fontfile FreeSansBold.ttf

When run with the font CABALETT, the result will look like this:

![Sample](https://raw.githubusercontent.com/StefanSchroeder/ttfsample/master/sample/CABALETT.TTF.png)

This will create a file "FreeSansBold.ttf.png" in the png/ directory which must exist.

The name of the font will always be included printed with a boring 
font, by default GNU FreeSansBold, but changable, that is always 
readable, even if the font has only symbols.

Have a look at the Makefile for more examples.

Author: Stefan Schröder, 2019

# Build

    go build ttfsample.go 

will do the trick if your Go development environment is setup properly.

Tested on Windows and Linux.





