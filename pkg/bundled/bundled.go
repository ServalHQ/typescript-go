// Package bundled re-exports access to files bundled with TypeScript from internal/bundled.
package bundled

import (
	"github.com/microsoft/typescript-go/internal/bundled"
	"github.com/microsoft/typescript-go/pkg/vfs"
)

// WrapFS returns an FS that redirects embedded TypeScript lib paths to the embedded file system.
// If embedding is not available, it returns the original FS unchanged.
func WrapFS(fs vfs.FS) vfs.FS {
	return bundled.WrapFS(fs)
}

// LibPath returns the path to the directory containing the bundled lib.d.ts files.
func LibPath() string {
	return bundled.LibPath()
}

// Embedded is true if the bundled files are implemented through an embedded FS.
const Embedded = bundled.Embedded
