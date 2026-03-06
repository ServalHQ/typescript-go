// Package compiler re-exports the TypeScript compiler program from internal/compiler.
// This package is part of the public API surface for the ServalHQ fork of typescript-go.
package compiler

import (
	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/diagnostics"
	"github.com/microsoft/typescript-go/pkg/tsoptions"
	"github.com/microsoft/typescript-go/pkg/vfs"
)

// Program represents a complete TypeScript program (a set of source files and compiler options).
// Create one with [NewProgram], then use it to obtain diagnostics and a type checker.
type Program = compiler.Program

// ProgramOptions configures how a Program is created.
type ProgramOptions = compiler.ProgramOptions

// CompilerHost provides the file system and source loading interface the compiler needs.
type CompilerHost = compiler.CompilerHost

// CheckerPool is the interface for managing type checker instances.
type CheckerPool = compiler.CheckerPool

// EmitResult contains the results of emitting output files.
type EmitResult = compiler.EmitResult

// DiagnosticMessage is a compiler diagnostic message definition.
type DiagnosticMessage = diagnostics.Message

// NewProgram creates a new TypeScript program from the given options.
//
// Example:
//
//	host := compiler.NewCompilerHost(cwd, fs, libPath, nil, nil)
//	config := tsoptions.NewParsedCommandLine(compilerOptions, fileNames, compareOpts)
//	program := compiler.NewProgram(compiler.ProgramOptions{Host: host, Config: config})
func NewProgram(opts ProgramOptions) *Program {
	return compiler.NewProgram(opts)
}

// NewCompilerHost creates a CompilerHost backed by the given file system.
// The trace function is optional (pass nil to disable tracing).
func NewCompilerHost(
	currentDirectory string,
	fs vfs.FS,
	defaultLibraryPath string,
	extendedConfigCache tsoptions.ExtendedConfigCache,
	trace func(msg *DiagnosticMessage, args ...any),
) CompilerHost {
	return compiler.NewCompilerHost(currentDirectory, fs, defaultLibraryPath, extendedConfigCache, trace)
}

// NewCachedFSCompilerHost is like [NewCompilerHost] but wraps the FS in a read cache
// for better performance.
func NewCachedFSCompilerHost(
	currentDirectory string,
	fs vfs.FS,
	defaultLibraryPath string,
	extendedConfigCache tsoptions.ExtendedConfigCache,
	trace func(msg *DiagnosticMessage, args ...any),
) CompilerHost {
	return compiler.NewCachedFSCompilerHost(currentDirectory, fs, defaultLibraryPath, extendedConfigCache, trace)
}
