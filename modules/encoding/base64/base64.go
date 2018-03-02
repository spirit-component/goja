package base64

import (
	original_base64 "encoding/base64"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("base64")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"NewDecoder":  original_base64.NewDecoder,
			"NewEncoder":  original_base64.NewEncoder,
			"NewEncoding": original_base64.NewEncoding,

			// Var and consts
			"NoPadding":      original_base64.NoPadding,
			"RawStdEncoding": original_base64.RawStdEncoding,
			"RawURLEncoding": original_base64.RawURLEncoding,
			"StdEncoding":    original_base64.StdEncoding,
			"StdPadding":     original_base64.StdPadding,
			"URLEncoding":    original_base64.URLEncoding,

			// Types (value type)
			"Encoding": func() original_base64.Encoding { return original_base64.Encoding{} },

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
