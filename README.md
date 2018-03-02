goja
----

**component of goja**


#### build-config

```bash
> spirit-builder pull --config build-goja.conf
> spirit-builder build --config build-goja.conf
```

`build-goja.conf`

> working with post-api

```hocon
# project
goja  {

	# import packages
	packages = ["github.com/spirit-component/goja", "github.com/spirit-component/postapi"]

	build-args = {
		go-get = ["-v"]
		go-build = ["-v"]
	}

	fetchers {
		git {
			gopath = ${GOPATH}
		}
		goget {
			gopath = ${GOPATH}
		}
	}


	# the dependencies
	repos = {
		goja {
			fetcher = goget
			args = ["-v"]
			url = "github.com/spirit-component/goja"
			revision = master
		}

		postapi {
			fetcher = goget
			args = ["-v"]
			url = "github.com/spirit-component/postapi"
			revision = master
		}
	}
}
```


#### run config

```bash
> ./goja run --config goja.conf
```


`goja.conf`

```hocon

# goja config
components.goja.api-mock {
	dir = "scripts/"
	timeout = 3s
}

# post-api graph config
components.post-api.external.grapher.default = {

	todo-task-new {
		name  = "todo.task.new"
		graph = {
			errors {
				response {
					seq = 1
					url = "spirit://actors/fbp/post-api/external?action=callback"
				}
			}
			normal {
				to-todo {
					seq = 1
					url = "spirit://actors/fbp/goja/api-mock?action=todo.task.new"
				}

				response {
					seq = 2
					url = "spirit://actors/fbp/post-api/external?action=callback"
				}
			}
		}
	}
}

```


#### javascript

`scripts/todo.task.new.js`

```javascript
go.Import("uuid")
go.Import("fmt")
go.Import("time")

time.Sleep(time.Second*10)

id = uuid.New()
fbp.SetBody({id: id})
```



### javascript golibs

name |import path | repositry
:--|:--|:--
md5 | `crypto/md5`|
base64 | `encoding/base64`|
json | `encoding/json`|
fmt | `fmt`|
uuid | `uuid`| `github.com/pborman/uuid`
ioutil|`io/ioutil`|
os|`os`|
exec|`os/exec`|
time|`time`|
redis|`redis`| `github.com/go-redis/redis`
log|`log`| `github.com/sirupsen/logrus`

#### import lib in javascript

```javascript
go.Import("fmt")
go.Import("encoding/base64")

fmt.Println(base64.StdEncoding.EncodeToString("hello"))
```


#### generate your own native go module to javascript module

my another project is: [gojs-tool](https://github.com/gogap/gojs-tool), use this command to generate the module bridge

- generate moudle

```bash
# Default GOPATH
gojs-tool gen --template goja -r -n github.com/go-redis/redis

# GO Internal LIBs
gojs-tool gen --gopath $(go env GOROOT) --template goja -r -n encoding/json
```

- update generated code, add import `github.com/spirit-component/goja/modules` and add `modules.RegisterNativeModule` at `func init()`

```
package your_package_name

import (
	// ....
	"github.com/spirit-component/goja/modules"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
   // the package name should be better equals to the "import/path/at/js" 's last slash name
   // here is js
	module = gojs.NewGojaModule("you_package_name")
)

func init() {
	// .........
	modules.RegisterNativeModule("import/path/at/js", Enable)
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
```

- update `build-goja.conf`

```
goja  {
	# append the generated module path
	packages = [..., "your_native_package_path"]
}
```
