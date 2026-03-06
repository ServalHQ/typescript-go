// Package core re-exports compiler options and core types from internal/core.
// This package is part of the public API surface for the ServalHQ fork of typescript-go.
package core

import "github.com/microsoft/typescript-go/internal/core"

// CompilerOptions represents the full set of TypeScript compiler options.
type CompilerOptions = core.CompilerOptions

// ParsedOptions holds parsed compiler options along with file names and project references.
type ParsedOptions = core.ParsedOptions

// Tristate represents a three-valued boolean: TSUnknown, TSFalse, TSTrue.
type Tristate = core.Tristate

const (
	TSUnknown = core.TSUnknown
	TSFalse   = core.TSFalse
	TSTrue    = core.TSTrue
)

var BoolToTristate = core.BoolToTristate

// ScriptKind identifies the language variant of a source file.
type ScriptKind = core.ScriptKind

const (
	ScriptKindUnknown  = core.ScriptKindUnknown
	ScriptKindJS       = core.ScriptKindJS
	ScriptKindJSX      = core.ScriptKindJSX
	ScriptKindTS       = core.ScriptKindTS
	ScriptKindTSX      = core.ScriptKindTSX
	ScriptKindExternal = core.ScriptKindExternal
	ScriptKindJSON     = core.ScriptKindJSON
	ScriptKindDeferred = core.ScriptKindDeferred
)

var GetScriptKindFromFileName = core.GetScriptKindFromFileName

// ScriptTarget represents the ECMAScript version target.
type ScriptTarget = core.ScriptTarget

const (
	ScriptTargetNone           = core.ScriptTargetNone
	ScriptTargetES5            = core.ScriptTargetES5
	ScriptTargetES2015         = core.ScriptTargetES2015
	ScriptTargetES2016         = core.ScriptTargetES2016
	ScriptTargetES2017         = core.ScriptTargetES2017
	ScriptTargetES2018         = core.ScriptTargetES2018
	ScriptTargetES2019         = core.ScriptTargetES2019
	ScriptTargetES2020         = core.ScriptTargetES2020
	ScriptTargetES2021         = core.ScriptTargetES2021
	ScriptTargetES2022         = core.ScriptTargetES2022
	ScriptTargetES2023         = core.ScriptTargetES2023
	ScriptTargetES2024         = core.ScriptTargetES2024
	ScriptTargetES2025         = core.ScriptTargetES2025
	ScriptTargetESNext         = core.ScriptTargetESNext
	ScriptTargetJSON           = core.ScriptTargetJSON
	ScriptTargetLatest         = core.ScriptTargetLatest
	ScriptTargetLatestStandard = core.ScriptTargetLatestStandard
)

// ModuleKind represents the module system to use.
type ModuleKind = core.ModuleKind

const (
	ModuleKindNone     = core.ModuleKindNone
	ModuleKindCommonJS = core.ModuleKindCommonJS
	ModuleKindAMD      = core.ModuleKindAMD
	ModuleKindUMD      = core.ModuleKindUMD
	ModuleKindSystem   = core.ModuleKindSystem
	ModuleKindES2015   = core.ModuleKindES2015
	ModuleKindES2020   = core.ModuleKindES2020
	ModuleKindES2022   = core.ModuleKindES2022
	ModuleKindESNext   = core.ModuleKindESNext
	ModuleKindNode16   = core.ModuleKindNode16
	ModuleKindNode18   = core.ModuleKindNode18
	ModuleKindNode20   = core.ModuleKindNode20
	ModuleKindNodeNext = core.ModuleKindNodeNext
	ModuleKindPreserve = core.ModuleKindPreserve
)

// ResolutionMode is an alias for ModuleKind used in module resolution contexts.
type ResolutionMode = core.ResolutionMode

const (
	ResolutionModeNone     = core.ResolutionModeNone
	ResolutionModeCommonJS = core.ResolutionModeCommonJS
	ResolutionModeESM      = core.ResolutionModeESM
)

// ModuleResolutionKind specifies the module resolution strategy.
type ModuleResolutionKind = core.ModuleResolutionKind

const (
	ModuleResolutionKindUnknown  = core.ModuleResolutionKindUnknown
	ModuleResolutionKindClassic  = core.ModuleResolutionKindClassic
	ModuleResolutionKindNode10   = core.ModuleResolutionKindNode10
	ModuleResolutionKindNode16   = core.ModuleResolutionKindNode16
	ModuleResolutionKindNodeNext = core.ModuleResolutionKindNodeNext
	ModuleResolutionKindBundler  = core.ModuleResolutionKindBundler
)

// ModuleDetectionKind specifies how module-ness is detected.
type ModuleDetectionKind = core.ModuleDetectionKind

const (
	ModuleDetectionKindNone   = core.ModuleDetectionKindNone
	ModuleDetectionKindAuto   = core.ModuleDetectionKindAuto
	ModuleDetectionKindLegacy = core.ModuleDetectionKindLegacy
	ModuleDetectionKindForce  = core.ModuleDetectionKindForce
)

// JsxEmit specifies how JSX is emitted.
type JsxEmit = core.JsxEmit

const (
	JsxEmitNone        = core.JsxEmitNone
	JsxEmitPreserve    = core.JsxEmitPreserve
	JsxEmitReactNative = core.JsxEmitReactNative
	JsxEmitReact       = core.JsxEmitReact
	JsxEmitReactJSX    = core.JsxEmitReactJSX
	JsxEmitReactJSXDev = core.JsxEmitReactJSXDev
)

// NewLineKind specifies the line ending style.
type NewLineKind = core.NewLineKind

const (
	NewLineKindNone = core.NewLineKindNone
	NewLineKindCRLF = core.NewLineKindCRLF
	NewLineKindLF   = core.NewLineKindLF
)
