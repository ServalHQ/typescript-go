# ServalHQ typescript-go Public API

This fork exposes key TypeScript compiler functionality as importable Go packages
via the `pkg/` directory. The core `internal/` packages remain untouched to
make upstream syncs straightforward.

## Quick Start

### 1. Add the dependency

In your consuming project's `go.mod`:

```
require github.com/microsoft/typescript-go v0.0.0-<timestamp>-<hash>
replace github.com/microsoft/typescript-go => github.com/ServalHQ/typescript-go v0.0.0-<timestamp>-<hash>
```

Or use a branch/tag reference:

```bash
go get github.com/ServalHQ/typescript-go@main
```

Then add the replace directive (required because the module path inside `go.mod`
is still `github.com/microsoft/typescript-go`):

```bash
go mod edit -replace github.com/microsoft/typescript-go=github.com/ServalHQ/typescript-go@main
```

### 2. Import the packages

```go
import (
    "github.com/microsoft/typescript-go/pkg/ast"
    "github.com/microsoft/typescript-go/pkg/bundled"
    "github.com/microsoft/typescript-go/pkg/checker"
    "github.com/microsoft/typescript-go/pkg/compiler"
    "github.com/microsoft/typescript-go/pkg/core"
    "github.com/microsoft/typescript-go/pkg/parser"
    "github.com/microsoft/typescript-go/pkg/tsoptions"
    "github.com/microsoft/typescript-go/pkg/tspath"
    "github.com/microsoft/typescript-go/pkg/vfs"
    "github.com/microsoft/typescript-go/pkg/vfs/osvfs"
)
```

## Available Packages

| Package | Description |
|---------|-------------|
| `pkg/ast` | AST node types, `Kind` constants, `SymbolFlags`, `ModifierFlags` |
| `pkg/bundled` | Access to bundled `lib.d.ts` files |
| `pkg/checker` | Type checker, `Type`, `Signature`, `TypeFlags`, `ObjectFlags` |
| `pkg/compiler` | `Program` creation, `CompilerHost` |
| `pkg/core` | `CompilerOptions`, `ScriptKind`, `Tristate`, `ModuleKind`, etc. |
| `pkg/parser` | `ParseSourceFile` for standalone parsing |
| `pkg/tsoptions` | `ParsedCommandLine` and config parsing |
| `pkg/tspath` | `Path` type and path utilities |
| `pkg/vfs` | `FS` interface (virtual file system) |
| `pkg/vfs/osvfs` | OS-backed `FS` implementation |

## Example: Type Check a Project

```go
package main

import (
    "context"
    "fmt"

    "github.com/microsoft/typescript-go/pkg/bundled"
    "github.com/microsoft/typescript-go/pkg/compiler"
    "github.com/microsoft/typescript-go/pkg/core"
    "github.com/microsoft/typescript-go/pkg/tsoptions"
    "github.com/microsoft/typescript-go/pkg/tspath"
    "github.com/microsoft/typescript-go/pkg/vfs/osvfs"
)

func main() {
    cwd := "/path/to/project"
    fs := bundled.WrapFS(osvfs.FS())
    libPath := bundled.LibPath()

    host := compiler.NewCompilerHost(cwd, fs, libPath, nil, nil)
    config := tsoptions.NewParsedCommandLine(
        &core.CompilerOptions{
            Strict: core.TSTrue,
            Target: core.ScriptTargetESNext,
            Module: core.ModuleKindESNext,
        },
        []string{cwd + "/src/index.ts"},
        tspath.ComparePathsOptions{
            UseCaseSensitiveFileNames: true,
            CurrentDirectory:          cwd,
        },
    )

    program := compiler.NewProgram(compiler.ProgramOptions{
        Host:   host,
        Config: config,
    })

    ctx := context.Background()
    diagnostics := program.GetSemanticDiagnostics(ctx, nil)
    for _, d := range diagnostics {
        fmt.Printf("%s: %s\n", d.File().FileName(), d.Message())
    }
    fmt.Printf("Found %d diagnostic(s)\n", len(diagnostics))
}
```

## Example: Extract Types from a Source File

```go
package main

import (
    "context"
    "fmt"

    "github.com/microsoft/typescript-go/pkg/ast"
    "github.com/microsoft/typescript-go/pkg/bundled"
    "github.com/microsoft/typescript-go/pkg/compiler"
    "github.com/microsoft/typescript-go/pkg/core"
    "github.com/microsoft/typescript-go/pkg/tsoptions"
    "github.com/microsoft/typescript-go/pkg/tspath"
    "github.com/microsoft/typescript-go/pkg/vfs/osvfs"
)

func main() {
    cwd := "/path/to/project"
    fs := bundled.WrapFS(osvfs.FS())
    libPath := bundled.LibPath()

    host := compiler.NewCompilerHost(cwd, fs, libPath, nil, nil)
    config := tsoptions.NewParsedCommandLine(
        &core.CompilerOptions{
            Strict: core.TSTrue,
            Target: core.ScriptTargetESNext,
            Module: core.ModuleKindESNext,
        },
        []string{cwd + "/src/index.ts"},
        tspath.ComparePathsOptions{
            UseCaseSensitiveFileNames: true,
            CurrentDirectory:          cwd,
        },
    )

    program := compiler.NewProgram(compiler.ProgramOptions{
        Host:   host,
        Config: config,
    })

    ctx := context.Background()
    tc, done := program.GetTypeChecker(ctx)
    defer done()

    for _, sf := range program.SourceFiles() {
        fmt.Printf("File: %s\n", sf.FileName())

        // Walk top-level statements
        for _, stmt := range sf.Statements.Nodes {
            switch stmt.Kind {
            case ast.KindVariableStatement:
                // Get types of declared variables
                declList := stmt.AsVariableStatement().DeclarationList
                for _, decl := range declList.AsVariableDeclarationList().Declarations.Nodes {
                    name := decl.Name()
                    if name != nil {
                        t := tc.GetTypeAtLocation(name)
                        if t != nil {
                            fmt.Printf("  %s: %s\n", name.Text(), tc.TypeToString(t))
                        }
                    }
                }

            case ast.KindFunctionDeclaration:
                fn := stmt.AsFunctionDeclaration()
                if fn.Name() != nil {
                    t := tc.GetTypeAtLocation(fn.Name())
                    if t != nil {
                        fmt.Printf("  function %s: %s\n", fn.Name().Text(), tc.TypeToString(t))
                    }
                }

            case ast.KindInterfaceDeclaration:
                iface := stmt.AsInterfaceDeclaration()
                symbol := tc.GetResolvedSymbol(iface.Name())
                if symbol != nil {
                    t := tc.GetDeclaredTypeOfSymbol(symbol)
                    if t != nil {
                        fmt.Printf("  interface %s: %s\n", iface.Name().Text(), tc.TypeToString(t))
                        for _, prop := range tc.GetPropertiesOfType(t) {
                            propType := tc.GetTypeOfSymbol(prop)
                            fmt.Printf("    .%s: %s\n", prop.Name, tc.TypeToString(propType))
                        }
                    }
                }
            }
        }
    }
}
```

## Example: Parse a Single File (No Type Checking)

```go
package main

import (
    "fmt"

    "github.com/microsoft/typescript-go/pkg/ast"
    "github.com/microsoft/typescript-go/pkg/core"
    "github.com/microsoft/typescript-go/pkg/parser"
    "github.com/microsoft/typescript-go/pkg/tspath"
)

func main() {
    source := `
        interface User {
            name: string;
            age: number;
        }
        const greet = (user: User): string => user.name;
    `

    file := parser.ParseSourceFile(
        ast.SourceFileParseOptions{
            FileName: "/main.ts",
            Path:     tspath.Path("/main.ts"),
        },
        source,
        core.ScriptKindTS,
    )

    fmt.Printf("Parsed %d statements\n", len(file.Statements.Nodes))
    for _, d := range file.Diagnostics() {
        fmt.Printf("Parse error: %s\n", d.Message())
    }
}
```

## Architecture

All `pkg/` packages are thin wrappers using Go type aliases (`type X = internal.X`)
pointing to the corresponding `internal/` packages. This means:

- **Zero overhead**: Type aliases are resolved at compile time, no runtime cost.
- **Full compatibility**: Types from `pkg/` and `internal/` are interchangeable.
- **Minimal maintenance**: Adding new exports is a one-line type alias.

The `internal/` packages remain completely untouched, making upstream merges clean.

## Upstream Sync Strategy

Since all changes are purely additive (new files in `pkg/`), syncing from
upstream is straightforward:

```bash
# Fetch latest from Microsoft's upstream
git fetch upstream

# Merge upstream changes into your branch
git merge upstream/main

# The merge should be clean since pkg/ files don't exist upstream.
# If there are conflicts in go.mod/go.sum, resolve them by accepting
# upstream changes and re-running: go mod tidy

# After merge, verify the pkg/ packages still compile
go build ./pkg/...
go vet ./pkg/...
```

If upstream adds new exported symbols you want to expose, add new type aliases
to the corresponding `pkg/` package.

If upstream renames or removes exported symbols, the build will fail after
the merge, making it easy to identify what needs updating in `pkg/`.
