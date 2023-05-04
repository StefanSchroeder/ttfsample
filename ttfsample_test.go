// Copyright (c) 2014 Stefan Schroeder, NY, 2014-03-10
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package main

import (
	"os"
	"testing"
)

var test_outdir = "test-output" + string(os.PathSeparator)

func Test_Example00(t *testing.T) {
	strArr := []string{"Abc", "Xyz", "Mno"}
	outdir = &test_outdir
	Printjabber("fonts/FreeSansBold.ttf", strArr)
}
