package worklog_test

import (
	"testing"

	"github.com/chonla/worklog"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewWorklog(t *testing.T) {
	hd := new(MockedHomeDir)
	hd.On("Path").Return("/some/path")

	w, e := worklog.NewWorklog(hd)

	assert.Nil(t, e)
	assert.NotNil(t, w)
	assert.Equal(t, "/some/path", w.StoragePath())
}
