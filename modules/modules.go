package modules

import (
	"fmt"

	"github.com/dop251/goja"
)

import (
	"github.com/spirit-component/goja/modules/crypto/md5"
	"github.com/spirit-component/goja/modules/encoding/base64"
	"github.com/spirit-component/goja/modules/encoding/json"
	jsfmt "github.com/spirit-component/goja/modules/fmt"
	"github.com/spirit-component/goja/modules/io/ioutil"
	"github.com/spirit-component/goja/modules/os"
	"github.com/spirit-component/goja/modules/os/exec"
	"github.com/spirit-component/goja/modules/time"

	"github.com/spirit-component/goja/modules/utils"

	"github.com/spirit-component/goja/modules/github.com/go-redis/redis"
	"github.com/spirit-component/goja/modules/github.com/google/uuid"
	"github.com/spirit-component/goja/modules/github.com/sirupsen/logrus"
)

var (
	golibs = map[string]func(runtime *goja.Runtime){
		"crypto/md5":      md5.Enable,
		"encoding/base64": base64.Enable,
		"encoding/json":   json.Enable,
		"fmt":             jsfmt.Enable,
		"time":            time.Enable,
		"uuid":            uuid.Enable,
		"io/ioutil":       ioutil.Enable,
		"utils":           utils.Enable,
		"os":              os.Enable,
		"os/exec":         exec.Enable,
		"redis":           redis.Enable,
		"log":             logrus.Enable,
	}
)

func RegisterNativeModule(importPath string, enableFunc func(runtime *goja.Runtime)) (err error) {
	if len(importPath) == 0 {
		err = fmt.Errorf("import path could not be empty")
		return
	}

	if enableFunc == nil {
		err = fmt.Errorf("register EnableFunc could not be nil, path: %s", importPath)
		return
	}

	golibs[importPath] = enableFunc

	return
}

type GoLib struct {
	vm *goja.Runtime
}

func NewGoLib(vm *goja.Runtime) *GoLib {
	return &GoLib{vm}
}

func (p *GoLib) Import(modules ...string) {
	if len(modules) == 0 {
		return
	}

	var loadFns []func(runtime *goja.Runtime)

	for _, module := range modules {
		fn, exist := golibs[module]
		if !exist {
			panic(fmt.Errorf("module of '%s' dost not exist", module))
		}

		loadFns = append(loadFns, fn)
	}

	for i := 0; i < len(loadFns); i++ {
		loadFns[i](p.vm)
	}
}
