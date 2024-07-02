package toolversions

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteToolVersions(t *testing.T) {
	// Use Afero's in-memory file system for testing
	fs := afero.NewMemMapFs()
	filePath := ".tool-versions"
	afero.WriteFile(fs, filePath, []byte(""), 0777)

	entries := []Entry{
		{Key: "nodejs", Value: "20.12.0"},
		{Key: "go", Value: "1.20"},
	}
	t.Run("test writeToolVersions", func(t *testing.T) {
		newFile, err := fs.Create(filePath)
		assert.NoError(t, err)

		err = writeToolVersions(fs, newFile, entries)
		assert.NoError(t, err)

		content, err := afero.ReadFile(fs, filePath)
		assert.NoError(t, err)

		expected := "nodejs 20.12.0\ngo 1.20\n"
		assert.Equal(t, expected, string(content))
	})
}