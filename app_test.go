package worklog_test

import (
	"testing"

	"github.com/chonla/worklog"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewAppWithCustomAppName(t *testing.T) {
	hd := new(MockedHomeDir)

	app, e := worklog.NewApp(hd)

	assert.NotNil(t, app)
	assert.Nil(t, e)
}

func TestGettingHomeDirShouldReturnPathFromHomeDir(t *testing.T) {
	hd := new(MockedHomeDir)
	hd.On("Path").Return("/some/path")

	app, _ := worklog.NewApp(hd)

	assert.Equal(t, "/some/path", app.StoragePath())

	hd.AssertExpectations(t)
}

func TestCreateNewWorklogFromApp(t *testing.T) {
	hd := new(MockedHomeDir)
	hd.On("Path").Return("/some/path")

	app, _ := worklog.NewApp(hd)

	w, e := app.CreateWorklog()

	assert.Nil(t, e)
	assert.NotNil(t, w)
}
