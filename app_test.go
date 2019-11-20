package worklog_test

import (
	"os"
	"testing"

	"github.com/chonla/worklog"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewAppWithExistingPath(t *testing.T) {
	worklog.PathStat = func(path string) (os.FileInfo, error) {
		return nil, nil
	}

	app, e := worklog.NewApp("/some/path")

	assert.NotNil(t, app)
	assert.Nil(t, e)
}

func TestCreateNewAppWithNonExistingPath(t *testing.T) {
	worklog.PathStat = func(path string) (os.FileInfo, error) {
		return nil, &os.PathError{Op: "stat", Path: "/some/nonexisting/path", Err: os.ErrNotExist}
	}

	app, e := worklog.NewApp("/some/nonexisting/path")

	assert.NotNil(t, e)
	assert.Equal(t, os.ErrNotExist, e.(*os.PathError).Err)
	assert.Nil(t, app)
}
