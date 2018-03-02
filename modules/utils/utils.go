package utils

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

func isNil(v interface{}) bool {
	return v == nil
}

var (
	module = gojs.NewGojaModule("utils")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"IsNil": isNil,
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
