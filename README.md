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
# default graph driver
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


```
# if you just want mapping api to js, you can use templer grapher
# the template file is json format, and it will render by text/template to replace some vars.

{
    "errors": {
        "ports": [{
            "seq": 1,
            "url": "spirit://actors/fbp/post-api/external?action=callback"
        }]
    },
    "normal": {
        "ports": [{
            "seq": 1,
            "url": "spirit://actors/fbp/goja/api-mock?action={{.api}}"
        },{
            "seq": 2,
            "url": "spirit://actors/fbp/post-api/external?action=callback"
        }]
    }
}

```


#### javascript

`scripts/todo.task.new.js`

```javascript
go.Import("uuid")
go.Import("fmt")
go.Import("time", "encoding/base64")


log.Infoln("hello I am the logger by logrus")

id = uuid.New()
obj = fbp.Object()
cache.Set(id, {id: id, name: obj.name})
fbp.SetBody({id: id})
```

`scripts/todo.task.get.js`

```javascript
obj = fbp.Object()

if (obj) {
	ret = cache.Get(obj.id)
	if (ret[1]) {
		fbp.SetBody(ret[0])
	} else {
		fbp.SetError("SPIRIT-GOJA",404,"task not exist")
	}
} else {
	fbp.SetError("SPIRIT-GOJA",400,"bad request")
}
```



### javascript

#### vars

name|type
:--|:--
session|github.com/go-spirit/spirit/mail.Session
cache|github.com/go-spirit/spirit/cache.Cache
config|github.com/gogap/config.Configuration
go|github.com/spirit-component/goja/modules.GoLib
fbp|github.com/spirit-component/goja.fbp


#### golibs

name |import path | repositry | default imported
:--|:--|:--|:--
md5 | `crypto/md5`||false
base64 | `encoding/base64`||false
json | `encoding/json`||false
fmt | `fmt`||false
uuid | `uuid`| `github.com/pborman/uuid`|false
ioutil|`io/ioutil`||false
os|`os`||false
exec|`os/exec`||false
time|`time`||false
redis|`redis`| `github.com/go-redis/redis`|false
log|`log`| `github.com/sirupsen/logrus`|true
utils|`utils`|`internal`|true




#### import lib in javascript

```javascript
go.Import("uuid")
go.Import("fmt")
go.Import("time", "encoding/base64")


log.Infoln("hello I am the logger by logrus")

id = uuid.New()
fbp.SetBody({id: id})

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

```go
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
