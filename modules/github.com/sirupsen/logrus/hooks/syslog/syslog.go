package syslog

import (
	original_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("syslog")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"NewSyslogHook": original_syslog.NewSyslogHook,

			// Var and consts

			// Types (value type)
			"SyslogHook": func() original_syslog.SyslogHook { return original_syslog.SyslogHook{} },

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
