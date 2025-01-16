// Copyright (c) 2014 Stefan Schroeder, NY, 2014-03-10
//
// The license that can be found in the LICENSE file

package main

import (
	"os"
	"testing"
)

var testOutdir = "test-output" + string(os.PathSeparator)

func Test_Example00(t *testing.T) {
	strArr := []string{"Abc", "Xyz", "Mno"}
	outdir = &testOutdir
	Printjabber("fonts/FreeSansBold.ttf", strArr, 2000)
}

func Test_Example01(t *testing.T) {
	strArr := []string{"Abc", "Xyz", "Mno"}
	outdir = &testOutdir
	Printjabber("fonts/FreeSerifBold.otf", strArr, 2001)
}
