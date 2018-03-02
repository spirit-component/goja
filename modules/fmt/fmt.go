package fmt

import (
	original_fmt "fmt"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("fmt")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"Errorf":   original_fmt.Errorf,
			"Fprint":   original_fmt.Fprint,
			"Fprintf":  original_fmt.Fprintf,
			"Fprintln": original_fmt.Fprintln,
			"Fscan":    original_fmt.Fscan,
			"Fscanf":   original_fmt.Fscanf,
			"Fscanln":  original_fmt.Fscanln,
			"Print":    original_fmt.Print,
			"Printf":   original_fmt.Printf,
			"Println":  original_fmt.Println,
			"Scan":     original_fmt.Scan,
			"Scanf":    original_fmt.Scanf,
			"Scanln":   original_fmt.Scanln,
			"Sprint":   original_fmt.Sprint,
			"Sprintf":  original_fmt.Sprintf,
			"Sprintln": original_fmt.Sprintln,
			"Sscan":    original_fmt.Sscan,
			"Sscanf":   original_fmt.Sscanf,
			"Sscanln":  original_fmt.Sscanln,

			// Var and consts

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
