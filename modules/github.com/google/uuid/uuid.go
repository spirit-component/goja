package uuid

import (
	original_uuid "github.com/google/uuid"
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
			"ClockSequence":       original_uuid.ClockSequence,
			"FromBytes":           original_uuid.FromBytes,
			"GetTime":             original_uuid.GetTime,
			"Must":                original_uuid.Must,
			"MustParse":           original_uuid.MustParse,
			"New":                 original_uuid.New,
			"NewDCEGroup":         original_uuid.NewDCEGroup,
			"NewDCEPerson":        original_uuid.NewDCEPerson,
			"NewDCESecurity":      original_uuid.NewDCESecurity,
			"NewHash":             original_uuid.NewHash,
			"NewMD5":              original_uuid.NewMD5,
			"NewRandom":           original_uuid.NewRandom,
			"NewRandomFromReader": original_uuid.NewRandomFromReader,
			"NewSHA1":             original_uuid.NewSHA1,
			"NewUUID":             original_uuid.NewUUID,
			"NodeID":              original_uuid.NodeID,
			"NodeInterface":       original_uuid.NodeInterface,
			"Parse":               original_uuid.Parse,
			"ParseBytes":          original_uuid.ParseBytes,
			"SetClockSequence":    original_uuid.SetClockSequence,
			"SetNodeID":           original_uuid.SetNodeID,
			"SetNodeInterface":    original_uuid.SetNodeInterface,
			"SetRand":             original_uuid.SetRand,

			// Var and consts
			"Future":        original_uuid.Future,
			"Group":         original_uuid.Group,
			"Invalid":       original_uuid.Invalid,
			"Microsoft":     original_uuid.Microsoft,
			"NameSpaceDNS":  original_uuid.NameSpaceDNS,
			"NameSpaceOID":  original_uuid.NameSpaceOID,
			"NameSpaceURL":  original_uuid.NameSpaceURL,
			"NameSpaceX500": original_uuid.NameSpaceX500,
			"Nil":           original_uuid.Nil,
			"Org":           original_uuid.Org,
			"Person":        original_uuid.Person,
			"RFC4122":       original_uuid.RFC4122,
			"Reserved":      original_uuid.Reserved,

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
