package worklog_test

import (
	"testing"

	worklog "github.com/ODDS-TEAM/go-worklog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	hd.On("With", mock.Anything).Return("")

	app, _ := worklog.NewApp(hd)

	w, e := app.CreateWorklog()

	assert.Nil(t, e)
	assert.NotNil(t, w)
}
