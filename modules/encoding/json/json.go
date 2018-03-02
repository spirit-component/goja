package json

import (
	original_json "encoding/json"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("json")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"Compact":       original_json.Compact,
			"HTMLEscape":    original_json.HTMLEscape,
			"Indent":        original_json.Indent,
			"Marshal":       original_json.Marshal,
			"MarshalIndent": original_json.MarshalIndent,
			"NewDecoder":    original_json.NewDecoder,
			"NewEncoder":    original_json.NewEncoder,
			"Unmarshal":     original_json.Unmarshal,
			"Valid":         original_json.Valid,

			// Var and consts

			// Types (value type)
			"Decoder":               func() original_json.Decoder { return original_json.Decoder{} },
			"Encoder":               func() original_json.Encoder { return original_json.Encoder{} },
			"InvalidUTF8Error":      func() original_json.InvalidUTF8Error { return original_json.InvalidUTF8Error{} },
			"InvalidUnmarshalError": func() original_json.InvalidUnmarshalError { return original_json.InvalidUnmarshalError{} },
			"MarshalerError":        func() original_json.MarshalerError { return original_json.MarshalerError{} },
			"SyntaxError":           func() original_json.SyntaxError { return original_json.SyntaxError{} },
			"UnmarshalFieldError":   func() original_json.UnmarshalFieldError { return original_json.UnmarshalFieldError{} },
			"UnmarshalTypeError":    func() original_json.UnmarshalTypeError { return original_json.UnmarshalTypeError{} },
			"UnsupportedTypeError":  func() original_json.UnsupportedTypeError { return original_json.UnsupportedTypeError{} },
			"UnsupportedValueError": func() original_json.UnsupportedValueError { return original_json.UnsupportedValueError{} },

			// Types (pointer type)
			"NewInvalidUTF8Error":      func() *original_json.InvalidUTF8Error { return &original_json.InvalidUTF8Error{} },
			"NewInvalidUnmarshalError": func() *original_json.InvalidUnmarshalError { return &original_json.InvalidUnmarshalError{} },
			"NewMarshalerError":        func() *original_json.MarshalerError { return &original_json.MarshalerError{} },
			"NewSyntaxError":           func() *original_json.SyntaxError { return &original_json.SyntaxError{} },
			"NewUnmarshalFieldError":   func() *original_json.UnmarshalFieldError { return &original_json.UnmarshalFieldError{} },
			"NewUnmarshalTypeError":    func() *original_json.UnmarshalTypeError { return &original_json.UnmarshalTypeError{} },
			"NewUnsupportedTypeError":  func() *original_json.UnsupportedTypeError { return &original_json.UnsupportedTypeError{} },
			"NewUnsupportedValueError": func() *original_json.UnsupportedValueError { return &original_json.UnsupportedValueError{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
