package logrus

import (
	original_logrus "github.com/sirupsen/logrus"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("log")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"AddHook":             original_logrus.AddHook,
			"Debug":               original_logrus.Debug,
			"Debugf":              original_logrus.Debugf,
			"Debugln":             original_logrus.Debugln,
			"Error":               original_logrus.Error,
			"Errorf":              original_logrus.Errorf,
			"Errorln":             original_logrus.Errorln,
			"Exit":                original_logrus.Exit,
			"Fatal":               original_logrus.Fatal,
			"Fatalf":              original_logrus.Fatalf,
			"Fatalln":             original_logrus.Fatalln,
			"GetLevel":            original_logrus.GetLevel,
			"Info":                original_logrus.Info,
			"Infof":               original_logrus.Infof,
			"Infoln":              original_logrus.Infoln,
			"New":                 original_logrus.New,
			"NewEntry":            original_logrus.NewEntry,
			"Panic":               original_logrus.Panic,
			"Panicf":              original_logrus.Panicf,
			"Panicln":             original_logrus.Panicln,
			"ParseLevel":          original_logrus.ParseLevel,
			"Print":               original_logrus.Print,
			"Printf":              original_logrus.Printf,
			"Println":             original_logrus.Println,
			"RegisterExitHandler": original_logrus.RegisterExitHandler,
			"SetFormatter":        original_logrus.SetFormatter,
			"SetLevel":            original_logrus.SetLevel,
			"SetOutput":           original_logrus.SetOutput,
			"StandardLogger":      original_logrus.StandardLogger,
			"Warn":                original_logrus.Warn,
			"Warnf":               original_logrus.Warnf,
			"Warning":             original_logrus.Warning,
			"Warningf":            original_logrus.Warningf,
			"Warningln":           original_logrus.Warningln,
			"Warnln":              original_logrus.Warnln,
			"WithError":           original_logrus.WithError,
			"WithField":           original_logrus.WithField,
			"WithFields":          original_logrus.WithFields,

			// Var and consts
			"AllLevels":     original_logrus.AllLevels,
			"DebugLevel":    original_logrus.DebugLevel,
			"ErrorKey":      original_logrus.ErrorKey,
			"ErrorLevel":    original_logrus.ErrorLevel,
			"FatalLevel":    original_logrus.FatalLevel,
			"FieldKeyLevel": original_logrus.FieldKeyLevel,
			"FieldKeyMsg":   original_logrus.FieldKeyMsg,
			"FieldKeyTime":  original_logrus.FieldKeyTime,
			"InfoLevel":     original_logrus.InfoLevel,
			"PanicLevel":    original_logrus.PanicLevel,
			"WarnLevel":     original_logrus.WarnLevel,

			// Types (value type)
			"Entry":         func() original_logrus.Entry { return original_logrus.Entry{} },
			"JSONFormatter": func() original_logrus.JSONFormatter { return original_logrus.JSONFormatter{} },
			"Logger":        func() original_logrus.Logger { return original_logrus.Logger{} },
			"MutexWrap":     func() original_logrus.MutexWrap { return original_logrus.MutexWrap{} },
			"TextFormatter": func() original_logrus.TextFormatter { return original_logrus.TextFormatter{} },

			// Types (pointer type)
			"NewJSONFormatter": func() *original_logrus.JSONFormatter { return &original_logrus.JSONFormatter{} },
			"NewLogger":        func() *original_logrus.Logger { return &original_logrus.Logger{} },
			"NewMutexWrap":     func() *original_logrus.MutexWrap { return &original_logrus.MutexWrap{} },
			"NewTextFormatter": func() *original_logrus.TextFormatter { return &original_logrus.TextFormatter{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
