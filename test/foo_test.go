package test

import (
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_foo(t *testing.T) {
	directory, err := project_root_directory.GetRootDirectory()
	assert.Nil(t, err)
	t.Log(directory)
}
