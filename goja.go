package goja

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"

	"github.com/go-spirit/spirit/component"
	"github.com/go-spirit/spirit/mail"
	"github.com/go-spirit/spirit/worker"
	"github.com/gogap/config"
	"github.com/spirit-component/goja/modules"
)

type Goja struct {
	alias string
	opts  component.Options
	conf  config.Configuration

	jsDir   string
	timeout time.Duration
}

func init() {
	component.RegisterComponent("goja", NewGoja)
}

func NewGoja(alias string, opts ...component.Option) (comp component.Component, err error) {

	compOpts := component.Options{}

	for _, o := range opts {
		o(&compOpts)
	}

	dir := compOpts.Config.GetString("dir")
	if len(dir) == 0 {
		err = errors.New("javascript dir is empty")
		return
	}

	fi, err := os.Stat(dir)
	if err != nil {
		return
	}

	if !fi.IsDir() {
		err = fmt.Errorf("path '%s' should be a folder", dir)
		return
	}

	g := &Goja{
		alias:   alias,
		opts:    compOpts,
		conf:    compOpts.Config,
		jsDir:   dir,
		timeout: compOpts.Config.GetTimeDuration("timeout", time.Second*30),
	}

	return g, nil
}

func (p *Goja) Start() error {
	return nil
}

func (p *Goja) Stop() error {
	return nil
}

func (p *Goja) Alias() string {
	if p == nil {
		return ""
	}

	return p.alias
}

func (p *Goja) Route(session mail.Session) worker.HandlerFunc {
	return p.runJS
}

func (p *Goja) runJS(session mail.Session) (err error) {

	action := session.Query("action")
	dir := session.Query("dir")
	strTimeout := session.Query("timeout")
	version := session.Query("version")

	if len(dir) == 0 {
		dir = p.jsDir
	}

	timeout := p.timeout
	if len(strTimeout) > 0 {
		if dur, e := time.ParseDuration(strTimeout); e == nil {
			timeout = dur
		}
	}

	var apiJsFile string

	if len(version) == 0 {
		apiJsFile = filepath.Join(dir, action+".js")
	} else {
		apiJsFile = filepath.Join(dir, action+"_"+version+".js")
	}

	jsData, err := ioutil.ReadFile(apiJsFile)
	if err != nil {
		return
	}

	vm := newVM()
	golib := modules.NewGoLib(vm)

	golib.Import("utils")
	golib.Import("log")

	vm.Set("session", session)
	vm.Set("caches", p.opts.Caches)
	vm.Set("config", p.opts.Config)
	vm.Set("go", golib)
	vm.Set("fbp", &fbp{session})

	prg, err := goja.Compile(apiJsFile, string(jsData), false)
	if err != nil {
		return
	}

	time.AfterFunc(timeout, func() {
		vm.Interrupt("goja runtime execute timeout")

	})

	_, vmErr := vm.RunProgram(prg)

	if vmErr != nil {
		switch jsErr := vmErr.(type) {
		case *goja.Exception:
			err = fmt.Errorf("%v", jsErr.Value())
		case *goja.InterruptedError:
			err = fmt.Errorf("execute action: '%s' interrupted, %v", action, jsErr.Value())
		default:
			err = vmErr
		}
	}

	return
}

func newVM() *goja.Runtime {
	vm := goja.New()
	vm.SetRandSource(newRandSource())
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	return vm
}

func newRandSource() goja.RandSource {
	var seed int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		panic(fmt.Errorf("Could not read random bytes: %v", err))
	}
	return rand.New(rand.NewSource(seed)).Float64
}
