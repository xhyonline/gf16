// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/xhyonline/gf16.

package gins

import (
	"fmt"
	"github.com/xhyonline/gf16/errors/gerror"
	"github.com/xhyonline/gf16/internal/intlog"
	"github.com/xhyonline/gf16/text/gstr"
	"github.com/xhyonline/gf16/util/gutil"

	"github.com/xhyonline/gf16/database/gdb"
	"github.com/xhyonline/gf16/text/gregex"
	"github.com/xhyonline/gf16/util/gconv"
)

const (
	frameCoreComponentNameDatabase = "gf.core.component.database"
	configNodeNameDatabase         = "database"
)

// Database returns an instance of database ORM object
// with specified configuration group name.
func Database(name ...string) gdb.DB {
	group := gdb.DefaultGroupName
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	instanceKey := fmt.Sprintf("%s.%s", frameCoreComponentNameDatabase, group)
	db := instances.GetOrSetFuncLock(instanceKey, func() interface{} {
		var (
			configMap     map[string]interface{}
			configNodeKey string
		)
		// It firstly searches the configuration of the instance name.
		if Config().Available() {
			configNodeKey, _ = gutil.MapPossibleItemByKey(
				Config().GetMap("."),
				configNodeNameDatabase,
			)
			if configNodeKey == "" {
				configNodeKey = configNodeNameDatabase
			}
			configMap = Config().GetMap(configNodeKey)
		}
		if len(configMap) == 0 && !gdb.IsConfigured() {
			configFilePath, err := Config().GetFilePath()
			if configFilePath == "" {
				exampleFileName := "config.example.toml"
				if exampleConfigFilePath, _ := Config().GetFilePath(exampleFileName); exampleConfigFilePath != "" {
					panic(gerror.Wrapf(
						err,
						`configuration file "%s" not found, but found "%s", did you miss renaming the example configuration file?`,
						Config().GetFileName(),
						exampleFileName,
					))
				} else {
					panic(gerror.Wrapf(
						err,
						`configuration file "%s" not found, did you miss the configuration file or the misspell the configuration file name?`,
						Config().GetFileName(),
					))
				}
			}
			panic(gerror.Wrapf(
				err,
				`database initialization failed: "%s" node not found, is configuration file or configuration node missing?`,
				configNodeNameDatabase,
			))
		}
		if len(configMap) == 0 {
			configMap = make(map[string]interface{})
		}
		// Parse <m> as map-slice and adds it to gdb's global configurations.
		for g, groupConfig := range configMap {
			cg := gdb.ConfigGroup{}
			switch value := groupConfig.(type) {
			case []interface{}:
				for _, v := range value {
					if node := parseDBConfigNode(v); node != nil {
						cg = append(cg, *node)
					}
				}
			case map[string]interface{}:
				if node := parseDBConfigNode(value); node != nil {
					cg = append(cg, *node)
				}
			}
			if len(cg) > 0 {
				if gdb.GetConfig(group) == nil {
					intlog.Printf("add configuration for group: %s, %#v", g, cg)
					gdb.SetConfigGroup(g, cg)
				} else {
					intlog.Printf("ignore configuration as it already exists for group: %s, %#v", g, cg)
					intlog.Printf("%s, %#v", g, cg)
				}
			}
		}
		// Parse <m> as a single node configuration,
		// which is the default group configuration.
		if node := parseDBConfigNode(configMap); node != nil {
			cg := gdb.ConfigGroup{}
			if node.Link != "" || node.Host != "" {
				cg = append(cg, *node)
			}

			if len(cg) > 0 {
				if gdb.GetConfig(group) == nil {
					intlog.Printf("add configuration for group: %s, %#v", gdb.DefaultGroupName, cg)
					gdb.SetConfigGroup(gdb.DefaultGroupName, cg)
				} else {
					intlog.Printf("ignore configuration as it already exists for group: %s, %#v", gdb.DefaultGroupName, cg)
					intlog.Printf("%s, %#v", gdb.DefaultGroupName, cg)
				}
			}
		}
		// Create a new ORM object with given configurations.
		if db, err := gdb.New(name...); err == nil {
			if Config().Available() {
				// Initialize logger for ORM.
				var loggerConfigMap map[string]interface{}
				loggerConfigMap = Config().GetMap(fmt.Sprintf("%s.%s", configNodeKey, configNodeNameLogger))
				if len(loggerConfigMap) == 0 {
					loggerConfigMap = Config().GetMap(configNodeKey)
				}
				if len(loggerConfigMap) > 0 {
					if err := db.GetLogger().SetConfigWithMap(loggerConfigMap); err != nil {
						panic(err)
					}
				}
			}
			return db
		} else {
			// It panics often because it dose not find its configuration for given group.
			panic(err)
		}
		return nil
	})
	if db != nil {
		return db.(gdb.DB)
	}
	return nil
}

func parseDBConfigNode(value interface{}) *gdb.ConfigNode {
	nodeMap, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	node := &gdb.ConfigNode{}
	err := gconv.Struct(nodeMap, node)
	if err != nil {
		panic(err)
	}
	// To be compatible with old version.
	if _, v := gutil.MapPossibleItemByKey(nodeMap, "LinkInfo"); v != nil {
		node.Link = gconv.String(v)
	}
	if _, v := gutil.MapPossibleItemByKey(nodeMap, "Link"); v != nil {
		node.Link = gconv.String(v)
	}
	// Parse link syntax.
	if node.Link != "" && node.Type == "" {
		match, _ := gregex.MatchString(`([a-z]+):(.+)`, node.Link)
		if len(match) == 3 {
			node.Type = gstr.Trim(match[1])
			node.Link = gstr.Trim(match[2])
		}
	}
	return node
}
