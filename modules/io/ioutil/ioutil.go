package ioutil

import (
	original_ioutil "io/ioutil"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("ioutil")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"NopCloser": original_ioutil.NopCloser,
			"ReadAll":   original_ioutil.ReadAll,
			"ReadDir":   original_ioutil.ReadDir,
			"ReadFile":  original_ioutil.ReadFile,
			"TempDir":   original_ioutil.TempDir,
			"TempFile":  original_ioutil.TempFile,
			"WriteFile": original_ioutil.WriteFile,

			// Var and consts
			"Discard": original_ioutil.Discard,

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
