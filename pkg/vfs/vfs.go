// Package vfs re-exports the virtual file system interface from internal/vfs.
// This package is part of the public API surface for the ServalHQ fork of typescript-go.
package vfs

import "github.com/microsoft/typescript-go/internal/vfs"

// FS is the file system abstraction used throughout the TypeScript compiler.
type FS = vfs.FS

// Entries represents the files and directories in a directory.
type Entries = vfs.Entries

// DirEntry is [fs.DirEntry].
type DirEntry = vfs.DirEntry

// FileInfo is [fs.FileInfo].
type FileInfo = vfs.FileInfo

// WalkDirFunc is [fs.WalkDirFunc].
type WalkDirFunc = vfs.WalkDirFunc

var (
	ErrInvalid    = vfs.ErrInvalid
	ErrPermission = vfs.ErrPermission
	ErrExist      = vfs.ErrExist
	ErrNotExist   = vfs.ErrNotExist
	ErrClosed     = vfs.ErrClosed
	SkipAll       = vfs.SkipAll
	SkipDir       = vfs.SkipDir
)
