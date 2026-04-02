# ServalHQ typescript-go Fork

Fork of [microsoft/typescript-go](https://github.com/microsoft/typescript-go) with a thin public wrapper layer.

## Why

Upstream typescript-go keeps all Go packages in `internal/` with no public API planned. This fork adds `pkg/serval/` — thin wrappers that re-export internal types and functions for use in the Serval backend.

## Rules

1. **Never modify files in `internal/`** — keep upstream mergeable
2. **Only add files in `pkg/serval/`** — our additions
3. **Sync upstream regularly:**
   ```bash
   git fetch upstream
   git merge upstream/main
   ```
   No conflicts expected since we never touch upstream files.

## Usage

```go
import "github.com/microsoft/typescript-go/pkg/serval"

sf := serval.ParseSourceFile("example.ts", `
  export function hello(): string {
    return "world";
  }
`)

serval.WalkAST(sf.AsNode(), func(node *serval.Node) bool {
    // visit each AST node
    return false
})
```

## Structure

```
pkg/serval/
├── parser.go   — ParseSourceFile, ParseTSX, ParseJS
└── walk.go     — WalkAST, FindNodes, re-exported type checks
```

## Requires

- Go 1.26+
