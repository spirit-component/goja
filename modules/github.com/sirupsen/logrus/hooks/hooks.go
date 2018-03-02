package hooks

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

import (
	"github.com/spirit-component/goja/modules/github.com/sirupsen/logrus/hooks/syslog"
	"github.com/spirit-component/goja/modules/github.com/sirupsen/logrus/hooks/test"
)

type hooks struct {
	syslog *goja.Object
	test   *goja.Object
}

func init() {
	require.RegisterNativeModule("hooks", Require)
}

func Require(runtime *goja.Runtime, module *goja.Object) {

	pkg := &hooks{
		syslog: require.Require(runtime, "syslog").(*goja.Object),
		test:   require.Require(runtime, "test").(*goja.Object),
	}

	o := module.Get("exports").(*goja.Object)

	o.Set("syslog", pkg.syslog)
	o.Set("test", pkg.test)
}

func Enable(runtime *goja.Runtime) {
	syslog.Enable(runtime)
	test.Enable(runtime)

	runtime.Set("hooks", require.Require(runtime, "hooks"))
}
