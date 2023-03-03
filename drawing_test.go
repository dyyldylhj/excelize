// Copyright 2016 - 2023 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.16 or later.

package excelize

import (
	"encoding/xml"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawingParser(t *testing.T) {
	f := File{
		Drawings: sync.Map{},
		Pkg:      sync.Map{},
	}
	f.Pkg.Store("charset", MacintoshCyrillicCharset)
	f.Pkg.Store("wsDr", []byte(xml.Header+`<xdr:wsDr xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"><xdr:oneCellAnchor><xdr:graphicFrame/></xdr:oneCellAnchor></xdr:wsDr>`))
	// Test with one cell anchor
	_, _, err := f.drawingParser("wsDr")
	assert.NoError(t, err)
	// Test with unsupported charset
	_, _, err = f.drawingParser("charset")
	assert.EqualError(t, err, "XML syntax error on line 1: invalid UTF-8")
	// Test with alternate content
	f.Drawings = sync.Map{}
	f.Pkg.Store("wsDr", []byte(xml.Header+`<xdr:wsDr xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"><mc:AlternateContent xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"><mc:Choice xmlns:a14="http://schemas.microsoft.com/office/drawing/2010/main" Requires="a14"><xdr:twoCellAnchor editAs="oneCell"></xdr:twoCellAnchor></mc:Choice><mc:Fallback/></mc:AlternateContent></xdr:wsDr>`))
	_, _, err = f.drawingParser("wsDr")
	assert.NoError(t, err)
}