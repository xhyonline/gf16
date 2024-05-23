// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"database/sql"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
	"reflect"
)

// ScanList converts `r` to struct slice which contains other complex struct attributes.
// Note that the parameter `listPointer` should be type of *[]struct/*[]*struct.
// Usage example:
//
// type Entity struct {
// 	   User       *EntityUser
// 	   UserDetail *EntityUserDetail
//	   UserScores []*EntityUserScores
// }
// var users []*Entity
// or
// var users []Entity
//
// ScanList(&users, "User")
// ScanList(&users, "UserDetail", "User", "uid:Uid")
// ScanList(&users, "UserScores", "User", "uid:Uid")
//
// The parameters "User/UserDetail/UserScores" in the example codes specify the target attribute struct
// that current result will be bound to.
//
// The "uid" in the example codes is the table field name of the result, and the "Uid" is the relational
// struct attribute name - not the attribute name of the bound to target. In the example codes, it's attribute
// name "Uid" of "User" of entity "Entity". It automatically calculates the HasOne/HasMany relationship with
// given `relation` parameter.
//
// See the example or unit testing cases for clear understanding for this function.
func (r Result) ScanList(listPointer interface{}, bindToAttrName string, relationKV ...string) (err error) {
	if r.IsEmpty() {
		return nil
	}
	// Necessary checks for parameters.
	if bindToAttrName == "" {
		return gerror.New(`bindToAttrName should not be empty`)
	}
	//if len(relation) > 0 {
	//	if len(relation) < 2 {
	//		return gerror.New(`relation name and key should are both necessary`)
	//	}
	//	if relation[0] == "" || relation[1] == "" {
	//		return gerror.New(`relation name and key should not be empty`)
	//	}
	//}

	var (
		reflectValue = reflect.ValueOf(listPointer)
		reflectKind  = reflectValue.Kind()
	)
	if reflectKind == reflect.Interface {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	if reflectKind != reflect.Ptr {
		return gerror.Newf("parameter should be type of *[]struct/*[]*struct, but got: %v", reflectKind)
	}
	reflectValue = reflectValue.Elem()
	reflectKind = reflectValue.Kind()
	if reflectKind != reflect.Slice && reflectKind != reflect.Array {
		return gerror.Newf("parameter should be type of *[]struct/*[]*struct, but got: %v", reflectKind)
	}
	length := len(r)
	if length == 0 {
		// The pointed slice is not empty.
		if reflectValue.Len() > 0 {
			// It here checks if it has struct item, which is already initialized.
			// It then returns error to warn the developer its empty and no conversion.
			if v := reflectValue.Index(0); v.Kind() != reflect.Ptr {
				return sql.ErrNoRows
			}
		}
		// Do nothing for empty struct slice.
		return nil
	}
	var (
		arrayValue    reflect.Value // Like: []*Entity
		arrayItemType reflect.Type  // Like: *Entity
		reflectType   = reflect.TypeOf(listPointer)
	)
	if reflectValue.Len() > 0 {
		arrayValue = reflectValue
	} else {
		arrayValue = reflect.MakeSlice(reflectType.Elem(), length, length)
	}

	// Slice element item.
	arrayItemType = arrayValue.Index(0).Type()

	// Relation variables.
	var (
		relationKVStr             string
		relationDataMap           map[string]Value
		relationFromAttrName      string // Eg: relationKV: User, uid:Uid -> User
		relationResultFieldName   string // Eg: relationKV: uid:Uid       -> uid
		relationBindToSubAttrName string // Eg: relationKV: uid:Uid       -> Uid
	)
	if len(relationKV) > 0 {
		if len(relationKV) == 1 {
			relationKVStr = relationKV[0]
		} else {
			relationFromAttrName = relationKV[0]
			relationKVStr = relationKV[1]
		}
		// The relation key string of table filed name and attribute name
		// can be joined with char '=' or ':'.
		array := gstr.SplitAndTrim(relationKVStr, "=")
		if len(array) == 1 {
			// Compatible with old splitting char ':'.
			array = gstr.SplitAndTrim(relationKVStr, ":")
		}
		if len(array) == 1 {
			// The relation names are the same.
			array = []string{relationKVStr, relationKVStr}
		}
		if len(array) == 2 {
			// Defined table field to relation attribute name.
			// Like:
			// uid:Uid
			// uid:UserId
			relationResultFieldName = array[0]
			relationBindToSubAttrName = array[1]
			if key, _ := gutil.MapPossibleItemByKey(r[0].Map(), relationResultFieldName); key == "" {
				return gerror.Newf(
					`cannot find possible related table field name "%s" from given relation key "%s"`,
					relationResultFieldName,
					relationKVStr,
				)
			} else {
				relationResultFieldName = key
			}
		} else {
			return gerror.New(`parameter relationKV should be format of "ResultFieldName:BindToAttrName"`)
		}
		if relationResultFieldName != "" {
			relationDataMap = r.MapKeyValue(relationResultFieldName)
		}
		if len(relationDataMap) == 0 {
			return gerror.Newf(`cannot find the relation data map, maybe invalid relation given "%v"`, relationKV)
		}
	}
	// Bind to target attribute.
	var (
		ok              bool
		bindToAttrValue reflect.Value
		bindToAttrKind  reflect.Kind
		bindToAttrType  reflect.Type
		bindToAttrField reflect.StructField
	)
	if arrayItemType.Kind() == reflect.Ptr {
		if bindToAttrField, ok = arrayItemType.Elem().FieldByName(bindToAttrName); !ok {
			return gerror.Newf(`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`, bindToAttrName)
		}
	} else {
		if bindToAttrField, ok = arrayItemType.FieldByName(bindToAttrName); !ok {
			return gerror.Newf(`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`, bindToAttrName)
		}
	}
	bindToAttrType = bindToAttrField.Type
	bindToAttrKind = bindToAttrType.Kind()

	// Bind to relation conditions.
	var (
		relationFromAttrValue            reflect.Value
		relationFromAttrField            reflect.Value
		relationBindToSubAttrNameChecked bool
	)
	for i := 0; i < arrayValue.Len(); i++ {
		arrayElemValue := arrayValue.Index(i)
		// The FieldByName should be called on non-pointer reflect.Value.
		if arrayElemValue.Kind() == reflect.Ptr {
			// Like: []*Entity
			arrayElemValue = arrayElemValue.Elem()
			if !arrayElemValue.IsValid() {
				// The element is nil, then create one and set it to the slice.
				// The "reflect.New(itemType.Elem())" creates a new element and returns the address of it.
				// For example:
				// reflect.New(itemType.Elem())        => *Entity
				// reflect.New(itemType.Elem()).Elem() => Entity
				arrayElemValue = reflect.New(arrayItemType.Elem()).Elem()
				arrayValue.Index(i).Set(arrayElemValue.Addr())
			}
		} else {
			// Like: []Entity
		}
		bindToAttrValue = arrayElemValue.FieldByName(bindToAttrName)
		if relationFromAttrName != "" {
			// Attribute value of current slice element.
			relationFromAttrValue = arrayElemValue.FieldByName(relationFromAttrName)
			if relationFromAttrValue.Kind() == reflect.Ptr {
				relationFromAttrValue = relationFromAttrValue.Elem()
			}
		} else {
			// Current slice element.
			relationFromAttrValue = arrayElemValue
		}
		if len(relationDataMap) > 0 && !relationFromAttrValue.IsValid() {
			return gerror.Newf(`invalid relation specified: "%v"`, relationKV)
		}
		// Check and find possible bind to attribute name.
		if relationKVStr != "" && !relationBindToSubAttrNameChecked {
			relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToSubAttrName)
			if !relationFromAttrField.IsValid() {
				var (
					relationFromAttrType = relationFromAttrValue.Type()
					filedMap             = make(map[string]interface{})
				)
				for i := 0; i < relationFromAttrType.NumField(); i++ {
					filedMap[relationFromAttrType.Field(i).Name] = struct{}{}
				}
				if key, _ := gutil.MapPossibleItemByKey(filedMap, relationBindToSubAttrName); key == "" {
					return gerror.Newf(
						`cannot find possible related attribute name "%s" from given relation key "%s"`,
						relationBindToSubAttrName,
						relationKVStr,
					)
				} else {
					relationBindToSubAttrName = key
				}
			}
			relationBindToSubAttrNameChecked = true
		}
		switch bindToAttrKind {
		case reflect.Array, reflect.Slice:
			if len(relationDataMap) > 0 {
				relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToSubAttrName)
				if relationFromAttrField.IsValid() {
					if err = gconv.Structs(
						relationDataMap[gconv.String(relationFromAttrField.Interface())],
						bindToAttrValue.Addr(),
					); err != nil {
						return err
					}
				} else {
					// May be the attribute does not exist yet.
					return gerror.Newf(`invalid relation specified: "%v"`, relationKV)
				}
			} else {
				return gerror.Newf(`relationKey should not be empty as field "%s" is slice`, bindToAttrName)
			}

		case reflect.Ptr:
			var element reflect.Value
			if bindToAttrValue.IsNil() {
				element = reflect.New(bindToAttrType.Elem()).Elem()
			} else {
				element = bindToAttrValue.Elem()
			}
			if len(relationDataMap) > 0 {
				relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToSubAttrName)
				if relationFromAttrField.IsValid() {
					v := relationDataMap[gconv.String(relationFromAttrField.Interface())]
					if v == nil {
						// There's no relational data.
						continue
					}
					if err = gconv.Struct(v, element); err != nil {
						return err
					}
				} else {
					// May be the attribute does not exist yet.
					return gerror.Newf(`invalid relation specified: "%v"`, relationKV)
				}
			} else {
				if i >= len(r) {
					// There's no relational data.
					continue
				}
				v := r[i]
				if v == nil {
					// There's no relational data.
					continue
				}
				if err = gconv.Struct(v, element); err != nil {
					return err
				}
			}
			bindToAttrValue.Set(element.Addr())

		case reflect.Struct:
			if len(relationDataMap) > 0 {
				relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToSubAttrName)
				if relationFromAttrField.IsValid() {
					relationDataItem := relationDataMap[gconv.String(relationFromAttrField.Interface())]
					if relationDataItem == nil {
						// There's no relational data.
						continue
					}
					if err = gconv.Struct(relationDataItem, bindToAttrValue); err != nil {
						return err
					}
				} else {
					// May be the attribute does not exist yet.
					return gerror.Newf(`invalid relation specified: "%v"`, relationKV)
				}
			} else {
				if i >= len(r) {
					// There's no relational data.
					continue
				}
				relationDataItem := r[i]
				if relationDataItem == nil {
					// There's no relational data.
					continue
				}
				if err = gconv.Struct(relationDataItem, bindToAttrValue); err != nil {
					return err
				}
			}

		default:
			return gerror.Newf(`unsupported attribute type: %s`, bindToAttrKind.String())
		}
	}
	reflect.ValueOf(listPointer).Elem().Set(arrayValue)
	return nil
}
