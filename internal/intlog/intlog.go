// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package intlog provides internal logging for GoFrame development usage only.
package intlog

import (
	"fmt"
	"github.com/gogf/gf/debug/gdebug"
	"github.com/gogf/gf/internal/utils"
	"path/filepath"
	"time"
)

const (
	stackFilterKey = "/internal/intlog"
)

var (
	// isGFDebug marks whether printing GoFrame debug information.
	isGFDebug = false
)

func init() {
	isGFDebug = utils.IsDebugEnabled()
}

// SetEnabled enables/disables the internal logging manually.
// Note that this function is not concurrent safe, be aware of the DATA RACE.
func SetEnabled(enabled bool) {
	// If they're the same, it does not write the `isGFDebug` but only reading operation.
	if isGFDebug != enabled {
		isGFDebug = enabled
	}
}

// Print prints `v` with newline using fmt.Println.
// The parameter `v` can be multiple variables.
func Print(v ...interface{}) {
	if !isGFDebug {
		return
	}
	fmt.Println(append([]interface{}{now(), "[INTE]", file()}, v...)...)
}

// Printf prints `v` with format `format` using fmt.Printf.
// The parameter `v` can be multiple variables.
func Printf(format string, v ...interface{}) {
	if !isGFDebug {
		return
	}
	fmt.Printf(now()+" [INTE] "+file()+" "+format+"\n", v...)
}

// Error prints `v` with newline using fmt.Println.
// The parameter `v` can be multiple variables.
func Error(v ...interface{}) {
	if !isGFDebug {
		return
	}
	array := append([]interface{}{now(), "[INTE]", file()}, v...)
	array = append(array, "\n"+gdebug.StackWithFilter(stackFilterKey))
	fmt.Println(array...)
}

// Errorf prints `v` with format `format` using fmt.Printf.
func Errorf(format string, v ...interface{}) {
	if !isGFDebug {
		return
	}
	fmt.Printf(
		now()+" [INTE] "+file()+" "+format+"\n%s\n",
		append(v, gdebug.StackWithFilter(stackFilterKey))...,
	)
}

// now returns current time string.
func now() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

// file returns caller file name along with its line number.
func file() string {
	_, p, l := gdebug.CallerWithFilter(stackFilterKey)
	return fmt.Sprintf(`%s:%d`, filepath.Base(p), l)
}
