package gitwrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPath(t *testing.T) {

	path, err := GetCurrentDirectory()

	assert.Nil(t, err, "is nill")
	assert.Contains(t, path, "git-phlow", "project folder is a part of the path")
}

func TestInit(t *testing.T) {
	var path = "git-phlow/"

	repo := IsRepository(path)

	assert.True(t, repo, true, "Repository is initialized")

}
