// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/xhyonline/gf16.

package gins

import (
	"fmt"
	"github.com/xhyonline/gf16/net/ghttp"
	"github.com/xhyonline/gf16/util/gutil"
)

const (
	frameCoreComponentNameServer = "gf.core.component.server"
	configNodeNameServer         = "server"
)

// Server returns an instance of http server with specified name.
func Server(name ...interface{}) *ghttp.Server {
	instanceKey := fmt.Sprintf("%s.%v", frameCoreComponentNameServer, name)
	return instances.GetOrSetFuncLock(instanceKey, func() interface{} {
		s := ghttp.GetServer(name...)
		// To avoid file no found error while it's not necessary.
		if Config().Available() {
			var m map[string]interface{}
			nodeKey, _ := gutil.MapPossibleItemByKey(Config().GetMap("."), configNodeNameServer)
			if nodeKey == "" {
				nodeKey = configNodeNameServer
			}
			m = Config().GetMap(fmt.Sprintf(`%s.%s`, nodeKey, s.GetName()))
			if len(m) == 0 {
				m = Config().GetMap(nodeKey)
			}
			if len(m) > 0 {
				if err := s.SetConfigWithMap(m); err != nil {
					panic(err)
				}
			}
			// As it might use template feature,
			// it initialize the view instance as well.
			_ = getViewInstance()
		}
		return s
	}).(*ghttp.Server)
}
