package uuid

import (
	original_uuid "github.com/pborman/uuid"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("uuid")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"ClockSequence":    original_uuid.ClockSequence,
			"Equal":            original_uuid.Equal,
			"GetTime":          original_uuid.GetTime,
			"New":              original_uuid.New,
			"NewDCEGroup":      original_uuid.NewDCEGroup,
			"NewDCEPerson":     original_uuid.NewDCEPerson,
			"NewDCESecurity":   original_uuid.NewDCESecurity,
			"NewHash":          original_uuid.NewHash,
			"NewMD5":           original_uuid.NewMD5,
			"NewRandom":        original_uuid.NewRandom,
			"NewSHA1":          original_uuid.NewSHA1,
			"NewUUID":          original_uuid.NewUUID,
			"NodeID":           original_uuid.NodeID,
			"NodeInterface":    original_uuid.NodeInterface,
			"Parse":            original_uuid.Parse,
			"SetClockSequence": original_uuid.SetClockSequence,
			"SetNodeID":        original_uuid.SetNodeID,
			"SetNodeInterface": original_uuid.SetNodeInterface,
			"SetRand":          original_uuid.SetRand,

			// Var and consts
			"Future":         original_uuid.Future,
			"Group":          original_uuid.Group,
			"Invalid":        original_uuid.Invalid,
			"Microsoft":      original_uuid.Microsoft,
			"NIL":            original_uuid.NIL,
			"NameSpace_DNS":  original_uuid.NameSpace_DNS,
			"NameSpace_OID":  original_uuid.NameSpace_OID,
			"NameSpace_URL":  original_uuid.NameSpace_URL,
			"NameSpace_X500": original_uuid.NameSpace_X500,
			"Org":            original_uuid.Org,
			"Person":         original_uuid.Person,
			"RFC4122":        original_uuid.RFC4122,
			"Reserved":       original_uuid.Reserved,

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
