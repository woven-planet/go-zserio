//go:build !windows
// +build !windows

package generator

import "syscall"

// FileOpenFlags are the flags that are used when opening a file on non-Windows
// systems
const FileOpenFlags = syscall.O_RDWR | syscall.O_CREAT | syscall.O_NOFOLLOW | syscall.O_TRUNC
