package project_root_directory

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_getExecutable(t *testing.T) {

	//t.Log(getExecutable())

	t.Log(os.Getwd())

}

func TestGetRootDirectory(t *testing.T) {
	directory, err := GetRootDirectory()
	assert.Nil(t, err)
	t.Log(directory)
}

func TestGetRootFilePath(t *testing.T) {
	path, err := GetRootFilePath("go.mod")
	assert.Nil(t, err)
	t.Log(path)
	// Output:
	// D:\workspace\go-project-root-directory\go.mod
}
