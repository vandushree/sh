// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package sh

import (
	"bytes"
	"io/ioutil"

	"github.com/mvdan/sh/syntax"
)

func Fuzz(data []byte) int {
	prog, err := syntax.Parse(bytes.NewReader(data), "", syntax.ParseComments)
	if err != nil {
		return 0
	}
	syntax.Fprint(ioutil.Discard, prog)
	return 1
}
