// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/xhyonline/gf16.

package gdb

import (
	"database/sql"
	"github.com/xhyonline/gf16/container/gmap"
	"github.com/xhyonline/gf16/encoding/gparser"
	"github.com/xhyonline/gf16/internal/empty"
	"github.com/xhyonline/gf16/util/gconv"
)

// Json converts `r` to JSON format content.
func (r Record) Json() string {
	content, _ := gparser.VarToJson(r.Map())
	return gconv.UnsafeBytesToStr(content)
}

// Xml converts `r` to XML format content.
func (r Record) Xml(rootTag ...string) string {
	content, _ := gparser.VarToXml(r.Map(), rootTag...)
	return gconv.UnsafeBytesToStr(content)
}

// Map converts `r` to map[string]interface{}.
func (r Record) Map() Map {
	m := make(map[string]interface{})
	for k, v := range r {
		m[k] = v.Val()
	}
	return m
}

// GMap converts `r` to a gmap.
func (r Record) GMap() *gmap.StrAnyMap {
	return gmap.NewStrAnyMapFrom(r.Map())
}

// Struct converts `r` to a struct.
// Note that the parameter `pointer` should be type of *struct/**struct.
//
// Note that it returns sql.ErrNoRows if `r` is empty.
func (r Record) Struct(pointer interface{}) error {
	// If the record is empty, it returns error.
	if r.IsEmpty() {
		if !empty.IsNil(pointer, true) {
			return sql.ErrNoRows
		}
		return nil
	}
	return gconv.StructTag(r.Map(), pointer, OrmTagForStruct)
}

// IsEmpty checks and returns whether `r` is empty.
func (r Record) IsEmpty() bool {
	return len(r) == 0
}
