package python

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/asgardehs/ratatoskr/embed_util"
	"github.com/asgardehs/ratatoskr/python/internal/data"
)

type EmbeddedPython struct {
	e *embed_util.EmbeddedFiles
	Python
}

// NewEmbeddedPython creates a new EmbeddedPython instance. The embedded source code and python binaries are
// extracted on demand using the given name as the base for the temporary directory. You should ensure that the chosen
// name does collide with other consumers of this library.
func NewEmbeddedPython(name string) (*EmbeddedPython, error) {
	e, err := embed_util.NewEmbeddedFiles(data.Data, fmt.Sprintf("python-%s", name))
	if err != nil {
		return nil, err
	}
	return &EmbeddedPython{
		e:      e,
		Python: NewPython(WithPythonHome(e.GetExtractedPath())),
	}, nil
}

// NewEmbeddedPythonInCacheDir extracts the embedded Python distribution into
// the user's OS cache directory (XDG_CACHE_HOME on Linux, ~/Library/Caches on
// macOS, %LOCALAPPDATA% on Windows) under a subdirectory named after appName.
// Unlike NewEmbeddedPython, the extracted distribution survives reboots and
// /tmp cleanup policies, which matters for long-lived desktop apps.
func NewEmbeddedPythonInCacheDir(appName string) (*EmbeddedPython, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, fmt.Errorf("ratatoskr: %w", err)
	}
	return NewEmbeddedPythonWithTmpDir(filepath.Join(cacheDir, appName, "python"), true)
}

func NewEmbeddedPythonWithTmpDir(tmpDir string, withHashInDir bool) (*EmbeddedPython, error) {
	e, err := embed_util.NewEmbeddedFilesWithTmpDir(data.Data, tmpDir, withHashInDir)
	if err != nil {
		return nil, err
	}
	return &EmbeddedPython{
		e:      e,
		Python: NewPython(WithPythonHome(e.GetExtractedPath())),
	}, nil
}

func (ep *EmbeddedPython) Cleanup() error {
	return ep.e.Cleanup()
}

func (ep *EmbeddedPython) GetExtractedPath() string {
	return ep.e.GetExtractedPath()
}
