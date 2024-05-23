// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/xhyonline/gf16.

package gi18n_test

import (
	"context"
	"testing"

	"github.com/xhyonline/gf16/os/gres"

	"github.com/xhyonline/gf16/os/gtime"
	"github.com/xhyonline/gf16/util/gconv"

	"github.com/xhyonline/gf16/frame/g"

	"github.com/xhyonline/gf16/i18n/gi18n"

	"github.com/xhyonline/gf16/debug/gdebug"
	"github.com/xhyonline/gf16/os/gfile"

	"github.com/xhyonline/gf16/test/gtest"

	_ "github.com/xhyonline/gf16/os/gres/testdata/data"
)

func Test_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: gdebug.TestDataPath("i18n"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
		i18n.SetDelimiters("{$", "}")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")
		t.Assert(i18n.T(context.Background(), "{$hello}{$world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: gdebug.TestDataPath("i18n-file"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n-dir",
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})
}

func Test_TranslateFormat(t *testing.T) {
	// Tf
	gtest.C(t, func(t *gtest.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: gdebug.TestDataPath("i18n"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.Tf(context.Background(), "{#hello}{#world} %d", 2020), "{#hello}{#world} 2020")

		i18n.SetLanguage("ja")
		t.Assert(i18n.Tf(context.Background(), "{#hello}{#world} %d", 2020), "こんにちは世界 2020")
	})
}

func Test_DefaultManager(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := gi18n.SetPath(gdebug.TestDataPath("i18n"))
		t.Assert(err, nil)

		gi18n.SetLanguage("none")
		t.Assert(gi18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		gi18n.SetLanguage("ja")
		t.Assert(gi18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		gi18n.SetLanguage("zh-CN")
		t.Assert(gi18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gi18n.SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n-dir")
		t.Assert(err, nil)

		gi18n.SetLanguage("none")
		t.Assert(gi18n.Translate(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		gi18n.SetLanguage("ja")
		t.Assert(gi18n.Translate(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		gi18n.SetLanguage("zh-CN")
		t.Assert(gi18n.Translate(context.Background(), "{#hello}{#world}"), "你好世界")
	})
}

func Test_Instance(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		m := gi18n.Instance()
		err := m.SetPath("i18n-dir")
		t.Assert(err, nil)
		m.SetLanguage("zh-CN")
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		m := gi18n.Instance()
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})

	gtest.C(t, func(t *gtest.T) {
		t.Assert(g.I18n().T(context.Background(), "{#hello}{#world}"), "你好世界")
	})
	// Default language is: en
	gtest.C(t, func(t *gtest.T) {
		m := gi18n.Instance(gconv.String(gtime.TimestampNano()))
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "HelloWorld")
	})
}

func Test_Resource(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.I18n("resource")
		err := m.SetPath("i18n-dir")
		t.Assert(err, nil)

		m.SetLanguage("none")
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		m.SetLanguage("ja")
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		m.SetLanguage("zh-CN")
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})
}
