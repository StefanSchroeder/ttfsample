/*
Take a TTF or OTF font file and write a PNG image
that contains a sample of that font.

SPDX: MIT

Written by Stefan Schröder. 2019, 2023
*/
package main

import (
	"golang.org/x/image/font/opentype"

	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/fs"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

const (
	imgW = 2000
	imgH = 800
)

// https://www.microsoft.com/typography/otspec/name.htm
var otfNameFields = map[int]string{
	0:  "Copyright",
	1:  "Family",
	2:  "Subfamily",
	3:  "UniqueIdentifier",
	4:  "Full",
	5:  "Version",
	6:  "PostScript",
	7:  "Trademark",
	8:  "Manufacturer",
	9:  "Designer",
	10: "Description",
	11: "VendorURL",
	12: "DesignerURL",
	13: "License",
	14: "LicenseURL",
	16: "TypographicFamily",
	17: "TypographicSubfamily",
	18: "CompatibleFull",
	19: "SampleText",
	20: "PostScriptCID",
	21: "WWSFamily",
	22: "WWSSubfamily",
	23: "LightBackgroundPalette",
	24: "DarkBackgroundPalette",
	25: "VariationsPostScriptPrefix",
}

var title = "Default Title"

var defaultJabberText = []string{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"!?%&1234567890üöäÜÖÄßéèáà@",
}

//go:embed fonts/FreeSansBold.ttf
var freesansbold []byte

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "", "filename of the ttf/otf font")
	hinting  = flag.String("hinting", "none", "none | full")
	outdir   = flag.String("outdir", "png", "Output directory")
	size     = flag.Float64("size", 100, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	walk     = flag.String("walk", "", "recursively look for fonts.")
)

func walkDirectories(s string, sampleText []string) {
	if fi, err := os.Stat(s); err == nil {
		switch {
		case fi.IsDir():
			err = filepath.Walk(*walk, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("Error accessing path %q: %v\n", path, err)
					return err
				}
				if strings.HasSuffix(path, ".ttf") || strings.HasSuffix(path, ".otf") {
					Printjabber(path, sampleText)
				}
				return nil
			})
			_ = err
		default:
			log.Printf("Walk arg is not a directory.")
		}
	} else {
		log.Printf("Walk arg is not a directory.")
	}
}

func main() {
	flag.Parse()

	wantedText := defaultJabberText
	if flag.NArg() != 0 {
		wantedText = flag.Args()
	}
	if *walk != "" {
		walkDirectories(*walk, wantedText)
		return
	}

	if _, err := os.Stat(*fontfile); err != nil {
		log.Printf("Missing file.\n")
		return
	}
	Printjabber(*fontfile, wantedText)
}

// Writefile writes the png file to disk.
func Writefile(outputName string, i *image.RGBA) {
	outFile, err := os.Create(outputName)
	if err != nil {
		log.Println(err)
		return
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	log.Printf("Written to \"" + outputName + "\"")
	err = png.Encode(b, i)
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

// Printjabber prints the string to an Image.
func Printjabber(ffile string, textToJabber []string) {
	// Draw the background and the guidelines.
	fg := image.Black
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}

	rgba := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	draw.Draw(rgba, rgba.Bounds(), image.White, image.Point{}, draw.Src)
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Hinting
	h := font.HintingNone
	switch *hinting {
	case "full":
		h = font.HintingFull
	}

	// Fetch the font
	fontBytes, err := ioutil.ReadFile(ffile)
	basename := filepath.Base(ffile)
	log.Println("Reading \"" + basename + "\"")
	if err != nil {
		log.Println(err)
		return
	}
	fontsize := *size

	fontObject, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	var b sfnt.Buffer
	fullfontname, _ := fontObject.Name(&b, 4)
	title = fullfontname

	// Print the meta-data
	for i := 0; i < 20; i++ {
		j, _ := fontObject.Name(&b, sfnt.NameID(i))
		fmt.Printf("    %v: <%v>\n", otfNameFields[i], j)
	}

	fontface, _ := opentype.NewFace(fontObject, &opentype.FaceOptions{
		Size:    fontsize,
		DPI:     *dpi,
		Hinting: h,
	})

	d := font.Drawer{
		Dst:  rgba,
		Src:  fg,
		Face: fontface,
	}

	// We could use d.MeasureString to get the width,
	// but it's not worth it.

	y := 10 + int(math.Ceil(fontsize**dpi/72))
	dy := int(math.Ceil(fontsize * *spacing * *dpi / 72))
	y += dy
	d.Dot = fixed.Point26_6{
		X: (fixed.I(imgW) - d.MeasureString(title)) / 2,
		Y: fixed.I(y),
	}

	// Draw name of font using that font
	d.DrawString(title)
	y += dy
	for _, s := range textToJabber {
		d.Dot = fixed.P(10, y)
		d.DrawString(s)
		y += dy
	}

	// ************************************
	// Print the header in the boring font
	// ************************************

	fBoringFont, err := truetype.Parse(freesansbold)
	if err != nil {
		log.Println(err)
		return
	}
	boringFace := truetype.NewFace(fBoringFont, &truetype.Options{
		Size:    fontsize,
		DPI:     *dpi,
		Hinting: h,
	})

	// Draw the title in the boring font
	drawerBoring := &font.Drawer{
		Dst:  rgba,
		Src:  fg,
		Face: boringFace,
	}

	y = 10 + int(math.Ceil(fontsize**dpi/72))
	drawerBoring.Dot = fixed.Point26_6{
		X: 100,
		Y: fixed.I(y),
	}
	drawerBoring.DrawString(title)

	// **********************************
	// Done writing to canvas.
	// **********************************

	err = os.MkdirAll(*outdir, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	Writefile(*outdir+"/"+basename+".png", rgba)
}
