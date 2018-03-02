package md5

import (
	original_md5 "crypto/md5"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("md5")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"New": original_md5.New,
			"Sum": original_md5.Sum,

			// Var and consts
			"BlockSize": original_md5.BlockSize,
			"Size":      original_md5.Size,
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
