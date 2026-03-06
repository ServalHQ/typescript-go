// Package osvfs re-exports the OS-backed virtual file system from internal/vfs/osvfs.
package osvfs

import (
	"github.com/microsoft/typescript-go/internal/vfs/osvfs"
	"github.com/microsoft/typescript-go/pkg/vfs"
)

// FS returns a [vfs.FS] backed by the real OS file system.
func FS() vfs.FS {
	return osvfs.FS()
}

// GetGlobalTypingsCacheLocation returns the platform-specific cache location for TypeScript typings.
func GetGlobalTypingsCacheLocation() string {
	return osvfs.GetGlobalTypingsCacheLocation()
}
