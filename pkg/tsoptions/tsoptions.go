// Package tsoptions re-exports TypeScript config parsing types from internal/tsoptions.
package tsoptions

import (
	"github.com/microsoft/typescript-go/internal/tsoptions"
	"github.com/microsoft/typescript-go/pkg/core"
	"github.com/microsoft/typescript-go/pkg/tspath"
)

// ParsedCommandLine represents a fully parsed tsconfig.json or command-line options.
type ParsedCommandLine = tsoptions.ParsedCommandLine

// ExtendedConfigCache is a cache for extended tsconfig resolution.
type ExtendedConfigCache = tsoptions.ExtendedConfigCache

// TsConfigSourceFile represents a parsed tsconfig.json file.
type TsConfigSourceFile = tsoptions.TsConfigSourceFile

// FileExtensionInfo describes a file extension and its script kind.
type FileExtensionInfo = tsoptions.FileExtensionInfo

// NewParsedCommandLine creates a ParsedCommandLine from compiler options and file names.
func NewParsedCommandLine(
	compilerOptions *core.CompilerOptions,
	rootFileNames []string,
	comparePathsOptions tspath.ComparePathsOptions,
) *ParsedCommandLine {
	return tsoptions.NewParsedCommandLine(compilerOptions, rootFileNames, comparePathsOptions)
}
