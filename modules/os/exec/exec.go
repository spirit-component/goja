package exec

import (
	original_exec "os/exec"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("exec")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"Command":        original_exec.Command,
			"CommandContext": original_exec.CommandContext,
			"LookPath":       original_exec.LookPath,

			// Var and consts
			"ErrNotFound": original_exec.ErrNotFound,

			// Types (value type)
			"Cmd":       func() original_exec.Cmd { return original_exec.Cmd{} },
			"Error":     func() original_exec.Error { return original_exec.Error{} },
			"ExitError": func() original_exec.ExitError { return original_exec.ExitError{} },

			// Types (pointer type)
			"NewCmd":       func() *original_exec.Cmd { return &original_exec.Cmd{} },
			"NewError":     func() *original_exec.Error { return &original_exec.Error{} },
			"NewExitError": func() *original_exec.ExitError { return &original_exec.ExitError{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
