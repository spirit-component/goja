package time

import (
	original_time "time"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("time")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"After":           original_time.After,
			"AfterFunc":       original_time.AfterFunc,
			"Date":            original_time.Date,
			"FixedZone":       original_time.FixedZone,
			"LoadLocation":    original_time.LoadLocation,
			"NewTicker":       original_time.NewTicker,
			"NewTimer":        original_time.NewTimer,
			"Now":             original_time.Now,
			"Parse":           original_time.Parse,
			"ParseDuration":   original_time.ParseDuration,
			"ParseInLocation": original_time.ParseInLocation,
			"Since":           original_time.Since,
			"Sleep":           original_time.Sleep,
			"Tick":            original_time.Tick,
			"Unix":            original_time.Unix,
			"Until":           original_time.Until,

			// Var and consts
			"ANSIC":       original_time.ANSIC,
			"April":       original_time.April,
			"August":      original_time.August,
			"December":    original_time.December,
			"February":    original_time.February,
			"Friday":      original_time.Friday,
			"Hour":        original_time.Hour,
			"January":     original_time.January,
			"July":        original_time.July,
			"June":        original_time.June,
			"Kitchen":     original_time.Kitchen,
			"Local":       original_time.Local,
			"March":       original_time.March,
			"May":         original_time.May,
			"Microsecond": original_time.Microsecond,
			"Millisecond": original_time.Millisecond,
			"Minute":      original_time.Minute,
			"Monday":      original_time.Monday,
			"Nanosecond":  original_time.Nanosecond,
			"November":    original_time.November,
			"October":     original_time.October,
			"RFC1123":     original_time.RFC1123,
			"RFC1123Z":    original_time.RFC1123Z,
			"RFC3339":     original_time.RFC3339,
			"RFC3339Nano": original_time.RFC3339Nano,
			"RFC822":      original_time.RFC822,
			"RFC822Z":     original_time.RFC822Z,
			"RFC850":      original_time.RFC850,
			"RubyDate":    original_time.RubyDate,
			"Saturday":    original_time.Saturday,
			"Second":      original_time.Second,
			"September":   original_time.September,
			"Stamp":       original_time.Stamp,
			"StampMicro":  original_time.StampMicro,
			"StampMilli":  original_time.StampMilli,
			"StampNano":   original_time.StampNano,
			"Sunday":      original_time.Sunday,
			"Thursday":    original_time.Thursday,
			"Tuesday":     original_time.Tuesday,
			"UTC":         original_time.UTC,
			"UnixDate":    original_time.UnixDate,
			"Wednesday":   original_time.Wednesday,

			// Types (value type)
			"Location":   func() original_time.Location { return original_time.Location{} },
			"ParseError": func() original_time.ParseError { return original_time.ParseError{} },
			"Ticker":     func() original_time.Ticker { return original_time.Ticker{} },
			"Time":       func() original_time.Time { return original_time.Time{} },
			"Timer":      func() original_time.Timer { return original_time.Timer{} },

			// Types (pointer type)
			"NewLocation":   func() *original_time.Location { return &original_time.Location{} },
			"NewParseError": func() *original_time.ParseError { return &original_time.ParseError{} },
			"NewTime":       func() *original_time.Time { return &original_time.Time{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
