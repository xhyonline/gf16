// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/xhyonline/gf16.

// go test *.go -bench=".*"

package gstr_test

import (
	"testing"

	"github.com/xhyonline/gf16/test/gtest"
	"github.com/xhyonline/gf16/text/gstr"
)

func Test_OctStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.OctStr(`\346\200\241`), "怡")
	})
}
