// Package tspath re-exports path types and utilities from internal/tspath.
package tspath

import "github.com/microsoft/typescript-go/internal/tspath"

// Path is the canonical representation of a file path used throughout the compiler.
type Path = tspath.Path

// ComparePathsOptions configures path comparison behavior.
type ComparePathsOptions = tspath.ComparePathsOptions

var (
	NormalizePath           = tspath.NormalizePath
	NormalizeSlashes        = tspath.NormalizeSlashes
	GetDirectoryPath        = tspath.GetDirectoryPath
	CombinePaths            = tspath.CombinePaths
	ResolvePath             = tspath.ResolvePath
	GetBaseFileName         = tspath.GetBaseFileName
	HasExtension            = tspath.HasExtension
	GetAnyExtensionFromPath = tspath.GetAnyExtensionFromPath
)
