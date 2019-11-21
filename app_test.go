package worklog_test

import (
	"testing"

	"github.com/chonla/worklog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedHomeDir struct {
	mock.Mock
}

func (m *MockedHomeDir) Path() string {
	args := m.Called()
	return args.String(0)
}

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

// func TestCreateNewWorklogFromApp(t *testing.T) {
// 	worklog.PathStat = func(path string) (os.FileInfo, error) {
// 		return nil, nil
// 	}

// 	app, _ := worklog.NewApp("/some/path")

// 	w, e := app.CreateWorklog()

// 	assert.Nil(t, e)
// 	assert.NotNil(t, w)
// }
