package python

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asgardehs/ratatoskr/internal"
)

func TestExternalPython(t *testing.T) {
	ep := NewPython()

	pexe, _ := ep.GetExePath()
	assert.True(t, internal.Exists(pexe))

	cmd, _ := ep.PythonCmd("-c", "print('test test')")
	stdout, err := cmd.StdoutPipe()
	assert.NoError(t, err)
	defer stdout.Close()

	err = cmd.Start()
	assert.NoError(t, err)

	stdoutStr, err := io.ReadAll(stdout)
	assert.NoError(t, err)

	err = cmd.Wait()
	assert.NoError(t, err)

	stdoutStr = bytes.TrimSpace(stdoutStr)
	assert.Equal(t, "test test", string(stdoutStr))
}
