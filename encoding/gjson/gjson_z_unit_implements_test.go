// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/xhyonline/gf16.

package gjson_test

import (
	"github.com/xhyonline/gf16/internal/json"
	"github.com/xhyonline/gf16/util/gconv"
	"testing"

	"github.com/xhyonline/gf16/encoding/gjson"
	"github.com/xhyonline/gf16/frame/g"
	"github.com/xhyonline/gf16/test/gtest"
)

func TestJson_UnmarshalJSON(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		j := gjson.New(nil)
		err := json.UnmarshalUseNumber(data, j)
		t.Assert(err, nil)
		t.Assert(j.Get("n"), "123456789")
		t.Assert(j.Get("m"), g.Map{"k": "v"})
		t.Assert(j.Get("m.k"), "v")
		t.Assert(j.Get("a"), g.Slice{1, 2, 3})
		t.Assert(j.Get("a.1"), 2)
	})
}

func TestJson_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Json *gjson.Json
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"json": []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`),
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.Get("n"), "123456789")
		t.Assert(v.Json.Get("m"), g.Map{"k": "v"})
		t.Assert(v.Json.Get("m.k"), "v")
		t.Assert(v.Json.Get("a"), g.Slice{1, 2, 3})
		t.Assert(v.Json.Get("a.1"), 2)
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"json": g.Map{
				"n": 123456789,
				"m": g.Map{"k": "v"},
				"a": g.Slice{1, 2, 3},
			},
		}, &v)
		t.Assert(err, nil)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.Get("n"), "123456789")
		t.Assert(v.Json.Get("m"), g.Map{"k": "v"})
		t.Assert(v.Json.Get("m.k"), "v")
		t.Assert(v.Json.Get("a"), g.Slice{1, 2, 3})
		t.Assert(v.Json.Get("a.1"), 2)
	})
}
