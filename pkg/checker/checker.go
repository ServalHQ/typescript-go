// Package checker re-exports the TypeScript type checker and type system from internal/checker.
// This package is part of the public API surface for the ServalHQ fork of typescript-go.
package checker

import internal "github.com/microsoft/typescript-go/internal/checker"

// Checker is the TypeScript type checker. Obtain one via compiler.Program.GetTypeChecker.
type Checker = internal.Checker

// Type represents a TypeScript type in the type system.
type Type = internal.Type

// TypeId is the unique identifier for a Type.
type TypeId = internal.TypeId

// Signature represents a function, method, or constructor signature.
type Signature = internal.Signature

// IndexInfo represents an index signature (e.g., [key: string]: value).
type IndexInfo = internal.IndexInfo

// TypePredicate represents a type predicate (e.g., "x is string").
type TypePredicate = internal.TypePredicate

// TypeAlias represents a type alias reference within the type system.
type TypeAlias = internal.TypeAlias

// SignatureKind distinguishes call signatures from construct signatures.
type SignatureKind = internal.SignatureKind

const (
	SignatureKindCall      = internal.SignatureKindCall
	SignatureKindConstruct = internal.SignatureKindConstruct
)

// TypeFormatFlags control how types are formatted to strings.
type TypeFormatFlags = internal.TypeFormatFlags

const (
	TypeFormatFlagsAddUndefined                        = internal.TypeFormatFlagsAddUndefined
	TypeFormatFlagsAllowUniqueESSymbolType             = internal.TypeFormatFlagsAllowUniqueESSymbolType
	TypeFormatFlagsGenerateNamesForShadowedTypeParams  = internal.TypeFormatFlagsGenerateNamesForShadowedTypeParams
	TypeFormatFlagsInArrayType                         = internal.TypeFormatFlagsInArrayType
	TypeFormatFlagsInElementType                       = internal.TypeFormatFlagsInElementType
	TypeFormatFlagsInFirstTypeArgument                 = internal.TypeFormatFlagsInFirstTypeArgument
	TypeFormatFlagsInTypeAlias                         = internal.TypeFormatFlagsInTypeAlias
	TypeFormatFlagsMultilineObjectLiterals             = internal.TypeFormatFlagsMultilineObjectLiterals
	TypeFormatFlagsNone                                = internal.TypeFormatFlagsNone
	TypeFormatFlagsNoTruncation                        = internal.TypeFormatFlagsNoTruncation
	TypeFormatFlagsNoTypeReduction                     = internal.TypeFormatFlagsNoTypeReduction
	TypeFormatFlagsOmitParameterModifiers              = internal.TypeFormatFlagsOmitParameterModifiers
	TypeFormatFlagsOmitThisParameter                   = internal.TypeFormatFlagsOmitThisParameter
	TypeFormatFlagsSuppressAnyReturnType               = internal.TypeFormatFlagsSuppressAnyReturnType
	TypeFormatFlagsUseAliasDefinedOutsideCurrentScope  = internal.TypeFormatFlagsUseAliasDefinedOutsideCurrentScope
	TypeFormatFlagsUseFullyQualifiedType               = internal.TypeFormatFlagsUseFullyQualifiedType
	TypeFormatFlagsUseSingleQuotesForStringLiteralType = internal.TypeFormatFlagsUseSingleQuotesForStringLiteralType
	TypeFormatFlagsUseStructuralFallback               = internal.TypeFormatFlagsUseStructuralFallback
	TypeFormatFlagsUseTypeOfFunction                   = internal.TypeFormatFlagsUseTypeOfFunction
	TypeFormatFlagsWriteArrayAsGenericType             = internal.TypeFormatFlagsWriteArrayAsGenericType
	TypeFormatFlagsWriteArrowStyleSignature            = internal.TypeFormatFlagsWriteArrowStyleSignature
	TypeFormatFlagsWriteCallStyleSignature             = internal.TypeFormatFlagsWriteCallStyleSignature
	TypeFormatFlagsWriteClassExpressionAsTypeLiteral   = internal.TypeFormatFlagsWriteClassExpressionAsTypeLiteral
	TypeFormatFlagsWriteTypeArgumentsOfSignature       = internal.TypeFormatFlagsWriteTypeArgumentsOfSignature
)

// SymbolFormatFlags control how symbols are formatted to strings.
type SymbolFormatFlags = internal.SymbolFormatFlags

const (
	SymbolFormatFlagsAllowAnyNodeKind                   = internal.SymbolFormatFlagsAllowAnyNodeKind
	SymbolFormatFlagsDoNotIncludeSymbolChain            = internal.SymbolFormatFlagsDoNotIncludeSymbolChain
	SymbolFormatFlagsNone                               = internal.SymbolFormatFlagsNone
	SymbolFormatFlagsUseAliasDefinedOutsideCurrentScope = internal.SymbolFormatFlagsUseAliasDefinedOutsideCurrentScope
	SymbolFormatFlagsUseOnlyExternalAliasing            = internal.SymbolFormatFlagsUseOnlyExternalAliasing
	SymbolFormatFlagsWriteComputedProps                 = internal.SymbolFormatFlagsWriteComputedProps
	SymbolFormatFlagsWriteTypeParametersOrArguments     = internal.SymbolFormatFlagsWriteTypeParametersOrArguments
)

// ContextFlags control contextual type resolution behavior.
type ContextFlags = internal.ContextFlags

const (
	ContextFlagsIgnoreNodeInferences = internal.ContextFlagsIgnoreNodeInferences
	ContextFlagsNoConstraints        = internal.ContextFlagsNoConstraints
	ContextFlagsNone                 = internal.ContextFlagsNone
	ContextFlagsSignature            = internal.ContextFlagsSignature
	ContextFlagsSkipBindingPatterns  = internal.ContextFlagsSkipBindingPatterns
)

// TypeFlags classify the kind of a type.
type TypeFlags = internal.TypeFlags

const (
	TypeFlagsAny                             = internal.TypeFlagsAny
	TypeFlagsAnyOrUnknown                    = internal.TypeFlagsAnyOrUnknown
	TypeFlagsBigInt                          = internal.TypeFlagsBigInt
	TypeFlagsBigIntLike                      = internal.TypeFlagsBigIntLike
	TypeFlagsBigIntLiteral                   = internal.TypeFlagsBigIntLiteral
	TypeFlagsBoolean                         = internal.TypeFlagsBoolean
	TypeFlagsBooleanLike                     = internal.TypeFlagsBooleanLike
	TypeFlagsBooleanLiteral                  = internal.TypeFlagsBooleanLiteral
	TypeFlagsConditional                     = internal.TypeFlagsConditional
	TypeFlagsDefinitelyFalsy                 = internal.TypeFlagsDefinitelyFalsy
	TypeFlagsDefinitelyNonNullable           = internal.TypeFlagsDefinitelyNonNullable
	TypeFlagsDisjointDomains                 = internal.TypeFlagsDisjointDomains
	TypeFlagsEnum                            = internal.TypeFlagsEnum
	TypeFlagsEnumLike                        = internal.TypeFlagsEnumLike
	TypeFlagsEnumLiteral                     = internal.TypeFlagsEnumLiteral
	TypeFlagsESSymbol                        = internal.TypeFlagsESSymbol
	TypeFlagsESSymbolLike                    = internal.TypeFlagsESSymbolLike
	TypeFlagsFreshable                       = internal.TypeFlagsFreshable
	TypeFlagsIncludesConstrainedTypeVariable = internal.TypeFlagsIncludesConstrainedTypeVariable
	TypeFlagsIncludesEmptyObject             = internal.TypeFlagsIncludesEmptyObject
	TypeFlagsIncludesError                   = internal.TypeFlagsIncludesError
	TypeFlagsIncludesInstantiable            = internal.TypeFlagsIncludesInstantiable
	TypeFlagsIncludesMask                    = internal.TypeFlagsIncludesMask
	TypeFlagsIncludesMissingType             = internal.TypeFlagsIncludesMissingType
	TypeFlagsIncludesNonWideningType         = internal.TypeFlagsIncludesNonWideningType
	TypeFlagsIncludesWildcard                = internal.TypeFlagsIncludesWildcard
	TypeFlagsIndex                           = internal.TypeFlagsIndex
	TypeFlagsIndexedAccess                   = internal.TypeFlagsIndexedAccess
	TypeFlagsInstantiable                    = internal.TypeFlagsInstantiable
	TypeFlagsInstantiableNonPrimitive        = internal.TypeFlagsInstantiableNonPrimitive
	TypeFlagsInstantiablePrimitive           = internal.TypeFlagsInstantiablePrimitive
	TypeFlagsIntersection                    = internal.TypeFlagsIntersection
	TypeFlagsIntrinsic                       = internal.TypeFlagsIntrinsic
	TypeFlagsLiteral                         = internal.TypeFlagsLiteral
	TypeFlagsNarrowable                      = internal.TypeFlagsNarrowable
	TypeFlagsNever                           = internal.TypeFlagsNever
	TypeFlagsNone                            = internal.TypeFlagsNone
	TypeFlagsNonPrimitive                    = internal.TypeFlagsNonPrimitive
	TypeFlagsNotPrimitiveUnion               = internal.TypeFlagsNotPrimitiveUnion
	TypeFlagsNull                            = internal.TypeFlagsNull
	TypeFlagsNullable                        = internal.TypeFlagsNullable
	TypeFlagsNumber                          = internal.TypeFlagsNumber
	TypeFlagsNumberLike                      = internal.TypeFlagsNumberLike
	TypeFlagsNumberLiteral                   = internal.TypeFlagsNumberLiteral
	TypeFlagsObject                          = internal.TypeFlagsObject
	TypeFlagsObjectFlagsType                 = internal.TypeFlagsObjectFlagsType
	TypeFlagsPossiblyFalsy                   = internal.TypeFlagsPossiblyFalsy
	TypeFlagsPrimitive                       = internal.TypeFlagsPrimitive
	TypeFlagsReserved1                       = internal.TypeFlagsReserved1
	TypeFlagsReserved2                       = internal.TypeFlagsReserved2
	TypeFlagsReserved3                       = internal.TypeFlagsReserved3
	TypeFlagsSimplifiable                    = internal.TypeFlagsSimplifiable
	TypeFlagsSingleton                       = internal.TypeFlagsSingleton
	TypeFlagsString                          = internal.TypeFlagsString
	TypeFlagsStringLike                      = internal.TypeFlagsStringLike
	TypeFlagsStringLiteral                   = internal.TypeFlagsStringLiteral
	TypeFlagsStringMapping                   = internal.TypeFlagsStringMapping
	TypeFlagsStringOrNumberLiteral           = internal.TypeFlagsStringOrNumberLiteral
	TypeFlagsStringOrNumberLiteralOrUnique   = internal.TypeFlagsStringOrNumberLiteralOrUnique
	TypeFlagsStructuredOrInstantiable        = internal.TypeFlagsStructuredOrInstantiable
	TypeFlagsStructuredType                  = internal.TypeFlagsStructuredType
	TypeFlagsSubstitution                    = internal.TypeFlagsSubstitution
	TypeFlagsTemplateLiteral                 = internal.TypeFlagsTemplateLiteral
	TypeFlagsTypeParameter                   = internal.TypeFlagsTypeParameter
	TypeFlagsTypeVariable                    = internal.TypeFlagsTypeVariable
	TypeFlagsUndefined                       = internal.TypeFlagsUndefined
	TypeFlagsUnion                           = internal.TypeFlagsUnion
	TypeFlagsUnionOrIntersection             = internal.TypeFlagsUnionOrIntersection
	TypeFlagsUniqueESSymbol                  = internal.TypeFlagsUniqueESSymbol
	TypeFlagsUnit                            = internal.TypeFlagsUnit
	TypeFlagsUnknown                         = internal.TypeFlagsUnknown
	TypeFlagsVoid                            = internal.TypeFlagsVoid
	TypeFlagsVoidLike                        = internal.TypeFlagsVoidLike
)

// ObjectFlags provide additional classification for object types.
type ObjectFlags = internal.ObjectFlags

const (
	ObjectFlagsAnonymous                                  = internal.ObjectFlagsAnonymous
	ObjectFlagsArrayLiteral                               = internal.ObjectFlagsArrayLiteral
	ObjectFlagsClass                                      = internal.ObjectFlagsClass
	ObjectFlagsClassOrInterface                           = internal.ObjectFlagsClassOrInterface
	ObjectFlagsContainsIntersections                      = internal.ObjectFlagsContainsIntersections
	ObjectFlagsContainsObjectOrArrayLiteral               = internal.ObjectFlagsContainsObjectOrArrayLiteral
	ObjectFlagsContainsSpread                             = internal.ObjectFlagsContainsSpread
	ObjectFlagsContainsWideningType                       = internal.ObjectFlagsContainsWideningType
	ObjectFlagsCouldContainTypeVariables                  = internal.ObjectFlagsCouldContainTypeVariables
	ObjectFlagsCouldContainTypeVariablesComputed          = internal.ObjectFlagsCouldContainTypeVariablesComputed
	ObjectFlagsEvolvingArray                              = internal.ObjectFlagsEvolvingArray
	ObjectFlagsFreshLiteral                               = internal.ObjectFlagsFreshLiteral
	ObjectFlagsIdenticalBaseTypeCalculated                = internal.ObjectFlagsIdenticalBaseTypeCalculated
	ObjectFlagsIdenticalBaseTypeExists                    = internal.ObjectFlagsIdenticalBaseTypeExists
	ObjectFlagsInstantiated                               = internal.ObjectFlagsInstantiated
	ObjectFlagsInstantiatedMapped                         = internal.ObjectFlagsInstantiatedMapped
	ObjectFlagsInstantiationExpressionType                = internal.ObjectFlagsInstantiationExpressionType
	ObjectFlagsInterface                                  = internal.ObjectFlagsInterface
	ObjectFlagsIsClassInstanceClone                       = internal.ObjectFlagsIsClassInstanceClone
	ObjectFlagsIsConstrainedTypeVariable                  = internal.ObjectFlagsIsConstrainedTypeVariable
	ObjectFlagsIsGenericIndexType                         = internal.ObjectFlagsIsGenericIndexType
	ObjectFlagsIsGenericObjectType                        = internal.ObjectFlagsIsGenericObjectType
	ObjectFlagsIsGenericType                              = internal.ObjectFlagsIsGenericType
	ObjectFlagsIsGenericTypeComputed                      = internal.ObjectFlagsIsGenericTypeComputed
	ObjectFlagsIsNeverIntersection                        = internal.ObjectFlagsIsNeverIntersection
	ObjectFlagsIsNeverIntersectionComputed                = internal.ObjectFlagsIsNeverIntersectionComputed
	ObjectFlagsIsUnknownLikeUnion                         = internal.ObjectFlagsIsUnknownLikeUnion
	ObjectFlagsIsUnknownLikeUnionComputed                 = internal.ObjectFlagsIsUnknownLikeUnionComputed
	ObjectFlagsJSLiteral                                  = internal.ObjectFlagsJSLiteral
	ObjectFlagsJsxAttributes                              = internal.ObjectFlagsJsxAttributes
	ObjectFlagsMapped                                     = internal.ObjectFlagsMapped
	ObjectFlagsMembersResolved                            = internal.ObjectFlagsMembersResolved
	ObjectFlagsNone                                       = internal.ObjectFlagsNone
	ObjectFlagsNonInferrableType                          = internal.ObjectFlagsNonInferrableType
	ObjectFlagsObjectLiteral                              = internal.ObjectFlagsObjectLiteral
	ObjectFlagsObjectLiteralPatternWithComputedProperties = internal.ObjectFlagsObjectLiteralPatternWithComputedProperties
	ObjectFlagsObjectRestType                             = internal.ObjectFlagsObjectRestType
	ObjectFlagsObjectTypeKindMask                         = internal.ObjectFlagsObjectTypeKindMask
	ObjectFlagsPrimitiveUnion                             = internal.ObjectFlagsPrimitiveUnion
	ObjectFlagsPropagatingFlags                           = internal.ObjectFlagsPropagatingFlags
	ObjectFlagsReference                                  = internal.ObjectFlagsReference
	ObjectFlagsRequiresWidening                           = internal.ObjectFlagsRequiresWidening
	ObjectFlagsReverseMapped                              = internal.ObjectFlagsReverseMapped
	ObjectFlagsSingleSignatureType                        = internal.ObjectFlagsSingleSignatureType
	ObjectFlagsTuple                                      = internal.ObjectFlagsTuple
)

// CheckMode controls how expressions are checked.
type CheckMode = internal.CheckMode

// IsTupleType checks if a type is a tuple type.
var IsTupleType = internal.IsTupleType

// GetDeclarationModifierFlagsFromSymbol returns the modifier flags of a symbol's declaration.
var GetDeclarationModifierFlagsFromSymbol = internal.GetDeclarationModifierFlagsFromSymbol
