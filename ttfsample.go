/*Take a TTF font file as an argument and write a PNG image
  that contains a sample of that font.

   * Use of this source code is governed by a BSD-style
   * license that can be found in the LICENSE file
  Written by Stefan Schröder. 2019 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
)

// https://developer.apple.com/fonts/TrueType-Reference-Manual/RM06/Chap6name.html
var ttfNameFields = map[int]string{
	0: "Copyright notice",
	1: "Font Family",
	2: "Font Subfamily",
	3: "Unique Subfamily identification",
	4: "Full name of the font",
	5: "Version of the name table",
}
var title = "Default Title"

const imgW, imgH = 1640, 800

var jabber = []string{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"!?%&1234567890üöäÜÖÄßéèáà@",
}

var (
	boringfont = flag.String("boringfont", "FreeSansBold.ttf", "The path to the boring font")
	verbose    = flag.Bool("verbose", false, "Print more info")
	dpi        = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile   = flag.String("fontfile", "../fnts/usr/share/fonts/truetype/hack/Hack-Bold.ttf", "filename of the ttf font")
	hinting    = flag.String("hinting", "none", "none | full")
	outdir     = flag.String("outdir", "png", "Output directory")
	size       = flag.Float64("size", 100, "font size in points")
	spacing    = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	text       = string("ABab1")
)

// Info is a wrapper around print to control verbosity.
func Info(format string, args ...interface{}) {
	if *verbose {
		msg := fmt.Sprintf(format, args...)
		fmt.Print(msg)
	}
}

func main() {
	flag.Parse()
	Printjabber(*fontfile)
}

// Printjabber does all the work.
func Printjabber(ffile string) {

	fontBytes, err := ioutil.ReadFile(ffile)
	basename := filepath.Base(ffile)
	log.Println("Reading \"" + basename + "\"")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	fBoringFont, err := truetype.Parse(getFreeSansBold())
	if err != nil {
		log.Println(err)
		return
	}

	fontname := f.Name(truetype.NameID(1))
	fontnam2 := f.Name(truetype.NameID(2))
	title = fontname + "/" + fontnam2
	for i := 0; i < 5; i++ {
		fmt.Printf("    %v: <%v>\n", ttfNameFields[i], f.Name(truetype.NameID(i)))
	}

	// Draw the background and the guidelines.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}

	rgba := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	h := font.HintingNone
	switch *hinting {
	case "full":
		h = font.HintingFull
	}
	resize := 1.0
	d := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    *size,
			DPI:     *dpi,
			Hinting: h,
		}),
	}

	lx := float64((d.MeasureString(title).Round()))
	if lx > 1640 { // The font+text is too wide. Resize!
		resize = 1640.0 / lx
		d = &font.Drawer{
			Dst: rgba,
			Src: fg,
			Face: truetype.NewFace(f, &truetype.Options{
				Size:    *size * resize,
				DPI:     *dpi,
				Hinting: h,
			}),
		}
	}
	y := 10 + int(math.Ceil(*size**dpi/72))
	dy := int(math.Ceil(*size * *spacing * *dpi / 72))
	y += dy
	d.Dot = fixed.Point26_6{
		X: (fixed.I(imgW) - d.MeasureString(title)) / 2,
		Y: fixed.I(y),
	}
	d.DrawString(title)
	y += dy
	for _, s := range jabber {
		d.Dot = fixed.P(10, y)
		d.DrawString(s)
		y += dy
	}

	// Draw the title in the standard name
	d2 := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(fBoringFont, &truetype.Options{
			Size:    *size * resize,
			DPI:     *dpi,
			Hinting: h,
		}),
	}
	y = 10 + int(math.Ceil(*size**dpi/72))
	d2.Dot = fixed.Point26_6{
		X: 100,
		Y: fixed.I(y),
	}
	d2.DrawString(title)

	// Write file
	outputName := *outdir + "/" + basename + ".png"
	outFile, err := os.Create(outputName)
	if err != nil {
		log.Println(err)
		return
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	log.Printf("Written to \"" + outputName + "\"")
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		return
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}
