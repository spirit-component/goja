package os

import (
	original_os "os"
)

import (
	"github.com/dop251/goja"
	"github.com/gogap/gojs-tool/gojs"
)

var (
	module = gojs.NewGojaModule("os")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"Chdir":           original_os.Chdir,
			"Chmod":           original_os.Chmod,
			"Chown":           original_os.Chown,
			"Chtimes":         original_os.Chtimes,
			"Clearenv":        original_os.Clearenv,
			"Create":          original_os.Create,
			"Environ":         original_os.Environ,
			"Executable":      original_os.Executable,
			"Exit":            original_os.Exit,
			"Expand":          original_os.Expand,
			"ExpandEnv":       original_os.ExpandEnv,
			"FindProcess":     original_os.FindProcess,
			"Getegid":         original_os.Getegid,
			"Getenv":          original_os.Getenv,
			"Geteuid":         original_os.Geteuid,
			"Getgid":          original_os.Getgid,
			"Getgroups":       original_os.Getgroups,
			"Getpagesize":     original_os.Getpagesize,
			"Getpid":          original_os.Getpid,
			"Getppid":         original_os.Getppid,
			"Getuid":          original_os.Getuid,
			"Getwd":           original_os.Getwd,
			"Hostname":        original_os.Hostname,
			"IsExist":         original_os.IsExist,
			"IsNotExist":      original_os.IsNotExist,
			"IsPathSeparator": original_os.IsPathSeparator,
			"IsPermission":    original_os.IsPermission,
			"Lchown":          original_os.Lchown,
			"Link":            original_os.Link,
			"LookupEnv":       original_os.LookupEnv,
			"Lstat":           original_os.Lstat,
			"Mkdir":           original_os.Mkdir,
			"MkdirAll":        original_os.MkdirAll,
			"NewFile":         original_os.NewFile,
			"NewSyscallError": original_os.NewSyscallError,
			"Open":            original_os.Open,
			"OpenFile":        original_os.OpenFile,
			"Pipe":            original_os.Pipe,
			"Readlink":        original_os.Readlink,
			"Remove":          original_os.Remove,
			"RemoveAll":       original_os.RemoveAll,
			"Rename":          original_os.Rename,
			"SameFile":        original_os.SameFile,
			"Setenv":          original_os.Setenv,
			"StartProcess":    original_os.StartProcess,
			"Stat":            original_os.Stat,
			"Symlink":         original_os.Symlink,
			"TempDir":         original_os.TempDir,
			"Truncate":        original_os.Truncate,
			"Unsetenv":        original_os.Unsetenv,

			// Var and consts
			"Args":              original_os.Args,
			"DevNull":           original_os.DevNull,
			"ErrClosed":         original_os.ErrClosed,
			"ErrExist":          original_os.ErrExist,
			"ErrInvalid":        original_os.ErrInvalid,
			"ErrNotExist":       original_os.ErrNotExist,
			"ErrPermission":     original_os.ErrPermission,
			"Interrupt":         original_os.Interrupt,
			"Kill":              original_os.Kill,
			"ModeAppend":        original_os.ModeAppend,
			"ModeCharDevice":    original_os.ModeCharDevice,
			"ModeDevice":        original_os.ModeDevice,
			"ModeDir":           original_os.ModeDir,
			"ModeExclusive":     original_os.ModeExclusive,
			"ModeNamedPipe":     original_os.ModeNamedPipe,
			"ModePerm":          original_os.ModePerm,
			"ModeSetgid":        original_os.ModeSetgid,
			"ModeSetuid":        original_os.ModeSetuid,
			"ModeSocket":        original_os.ModeSocket,
			"ModeSticky":        original_os.ModeSticky,
			"ModeSymlink":       original_os.ModeSymlink,
			"ModeTemporary":     original_os.ModeTemporary,
			"ModeType":          original_os.ModeType,
			"O_APPEND":          original_os.O_APPEND,
			"O_CREATE":          original_os.O_CREATE,
			"O_EXCL":            original_os.O_EXCL,
			"O_RDONLY":          original_os.O_RDONLY,
			"O_RDWR":            original_os.O_RDWR,
			"O_SYNC":            original_os.O_SYNC,
			"O_TRUNC":           original_os.O_TRUNC,
			"O_WRONLY":          original_os.O_WRONLY,
			"PathListSeparator": original_os.PathListSeparator,
			"PathSeparator":     original_os.PathSeparator,
			"SEEK_CUR":          original_os.SEEK_CUR,
			"SEEK_END":          original_os.SEEK_END,
			"SEEK_SET":          original_os.SEEK_SET,
			"Stderr":            original_os.Stderr,
			"Stdin":             original_os.Stdin,
			"Stdout":            original_os.Stdout,

			// Types (value type)
			"File":         func() original_os.File { return original_os.File{} },
			"LinkError":    func() original_os.LinkError { return original_os.LinkError{} },
			"PathError":    func() original_os.PathError { return original_os.PathError{} },
			"ProcAttr":     func() original_os.ProcAttr { return original_os.ProcAttr{} },
			"Process":      func() original_os.Process { return original_os.Process{} },
			"ProcessState": func() original_os.ProcessState { return original_os.ProcessState{} },
			"SyscallError": func() original_os.SyscallError { return original_os.SyscallError{} },

			// Types (pointer type)
			"NewLinkError":    func() *original_os.LinkError { return &original_os.LinkError{} },
			"NewPathError":    func() *original_os.PathError { return &original_os.PathError{} },
			"NewProcAttr":     func() *original_os.ProcAttr { return &original_os.ProcAttr{} },
			"NewProcess":      func() *original_os.Process { return &original_os.Process{} },
			"NewProcessState": func() *original_os.ProcessState { return &original_os.ProcessState{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
