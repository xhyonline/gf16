// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gvalid_test

import (
	"context"
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

func ExampleCheckMap() {
	params := map[string]interface{}{
		"passport":  "",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := []string{
		"passport@required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
		"password2@required|length:6,16#",
	}
	if e := gvalid.CheckMap(context.TODO(), params, rules); e != nil {
		fmt.Println(e.Map())
		fmt.Println(e.FirstItem())
		fmt.Println(e.FirstString())
	}
	// May Output:
	// map[required:账号不能为空 length:账号长度应当在6到16之间]
	// passport map[required:账号不能为空 length:账号长度应当在6到16之间]
	// 账号不能为空
}

func ExampleCheckMap2() {
	params := map[string]interface{}{
		"passport":  "",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := []string{
		"passport@length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
		"password2@required|length:6,16#",
	}
	if e := gvalid.CheckMap(context.TODO(), params, rules); e != nil {
		fmt.Println(e.Map())
		fmt.Println(e.FirstItem())
		fmt.Println(e.FirstString())
	}
	// Output:
	// map[same:两次密码输入不相等]
	// password map[same:两次密码输入不相等]
	// 两次密码输入不相等
}

// Empty string attribute.
func ExampleCheckStruct() {
	type Params struct {
		Page      int    `v:"required|min:1         # page is required"`
		Size      int    `v:"required|between:1,100 # size is required"`
		ProjectId string `v:"between:1,10000        # project id must between :min, :max"`
	}
	obj := &Params{
		Page: 1,
		Size: 10,
	}
	err := gvalid.CheckStruct(context.TODO(), obj, nil)
	fmt.Println(err == nil)
	// Output:
	// true
}

// Empty pointer attribute.
func ExampleCheckStruct2() {
	type Params struct {
		Page      int       `v:"required|min:1         # page is required"`
		Size      int       `v:"required|between:1,100 # size is required"`
		ProjectId *gvar.Var `v:"between:1,10000        # project id must between :min, :max"`
	}
	obj := &Params{
		Page: 1,
		Size: 10,
	}
	err := gvalid.CheckStruct(context.TODO(), obj, nil)
	fmt.Println(err == nil)
	// Output:
	// true
}

// Empty integer attribute.
func ExampleCheckStruct3() {
	type Params struct {
		Page      int `v:"required|min:1         # page is required"`
		Size      int `v:"required|between:1,100 # size is required"`
		ProjectId int `v:"between:1,10000        # project id must between :min, :max"`
	}
	obj := &Params{
		Page: 1,
		Size: 10,
	}
	err := gvalid.CheckStruct(context.TODO(), obj, nil)
	fmt.Println(err)
	// Output:
	// project id must between 1, 10000
}

func ExampleRegisterRule() {
	type User struct {
		Id   int
		Name string `v:"required|unique-name # 请输入用户名称|用户名称已被占用"`
		Pass string `v:"required|length:6,18"`
	}
	user := &User{
		Id:   1,
		Name: "john",
		Pass: "123456",
	}

	rule := "unique-name"
	gvalid.RegisterRule(rule, func(ctx context.Context, rule string, value interface{}, message string, data interface{}) error {
		var (
			id   = data.(*User).Id
			name = gconv.String(value)
		)
		n, err := g.Model("user").Where("id != ? and name = ?", id, name).Count()
		if err != nil {
			return err
		}
		if n > 0 {
			return errors.New(message)
		}
		return nil
	})
	err := gvalid.CheckStruct(context.TODO(), user, nil)
	fmt.Println(err.Error())
	// May Output:
	// 用户名称已被占用
}

func ExampleRegisterRule_OverwriteRequired() {
	rule := "required"
	gvalid.RegisterRule(rule, func(ctx context.Context, rule string, value interface{}, message string, data interface{}) error {
		reflectValue := reflect.ValueOf(value)
		if reflectValue.Kind() == reflect.Ptr {
			reflectValue = reflectValue.Elem()
		}
		isEmpty := false
		switch reflectValue.Kind() {
		case reflect.Bool:
			isEmpty = !reflectValue.Bool()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			isEmpty = reflectValue.Int() == 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			isEmpty = reflectValue.Uint() == 0
		case reflect.Float32, reflect.Float64:
			isEmpty = math.Float64bits(reflectValue.Float()) == 0
		case reflect.Complex64, reflect.Complex128:
			c := reflectValue.Complex()
			isEmpty = math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
		case reflect.String, reflect.Map, reflect.Array, reflect.Slice:
			isEmpty = reflectValue.Len() == 0
		}
		if isEmpty {
			return errors.New(message)
		}
		return nil
	})
	fmt.Println(gvalid.CheckValue(context.TODO(), "", "required", "It's required"))
	fmt.Println(gvalid.CheckValue(context.TODO(), 0, "required", "It's required"))
	fmt.Println(gvalid.CheckValue(context.TODO(), false, "required", "It's required"))
	gvalid.DeleteRule(rule)
	fmt.Println("rule deleted")
	fmt.Println(gvalid.CheckValue(context.TODO(), "", "required", "It's required"))
	fmt.Println(gvalid.CheckValue(context.TODO(), 0, "required", "It's required"))
	fmt.Println(gvalid.CheckValue(context.TODO(), false, "required", "It's required"))
	// Output:
	// It's required
	// It's required
	// It's required
	// rule deleted
	// It's required
	// <nil>
	// <nil>
}

func ExampleValidator_Rules() {
	data := g.Map{
		"password": "123",
	}
	err := g.Validator().Data(data).Rules("required-with:password").Messages("请输入确认密码").CheckValue("")
	fmt.Println(err.String())

	// Output:
	// 请输入确认密码
}

func ExampleValidator_CheckValue() {
	err := g.Validator().Rules("min:18").Messages("未成年人不允许注册哟").CheckValue(16)
	fmt.Println(err.String())

	// Output:
	// 未成年人不允许注册哟
}

func ExampleValidator_CheckMap() {
	params := map[string]interface{}{
		"passport":  "",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := map[string]string{
		"passport":  "required|length:6,16",
		"password":  "required|length:6,16|same:password2",
		"password2": "required|length:6,16",
	}
	messages := map[string]interface{}{
		"passport": "账号不能为空|账号长度应当在:min到:max之间",
		"password": map[string]string{
			"required": "密码不能为空",
			"same":     "两次密码输入不相等",
		},
	}
	err := g.Validator().Messages(messages).Rules(rules).CheckMap(params)
	if err != nil {
		g.Dump(err.Maps())
	}

	// May Output:
	//{
	//	"passport": {
	//	"length": "账号长度应当在6到16之间",
	//		"required": "账号不能为空"
	//},
	//	"password": {
	//	"same": "两次密码输入不相等"
	//}
	//}
}

func ExampleValidator_CheckStruct() {
	type User struct {
		Name string `v:"required#请输入用户姓名"`
		Type int    `v:"required#请选择用户类型"`
	}
	data := g.Map{
		"name": "john",
	}
	user := User{}
	if err := gconv.Scan(data, &user); err != nil {
		panic(err)
	}
	err := g.Validator().Data(data).CheckStruct(user)
	if err != nil {
		fmt.Println(err.Items())
	}

	// Output:
	// [map[Type:map[required:请选择用户类型]]]
}
