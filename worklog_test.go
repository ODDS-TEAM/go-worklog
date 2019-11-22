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

func TestGettingSiteList(t *testing.T) {
	hd := new(MockedHomeDir)
	hd.On("Path").Return("/some/path")

	w, _ := worklog.NewWorklog(hd)

	l := w.SiteList()

	assert.Len(t, l, 0)
}

func TestAddingNewSiteToListShouldBeSuccess(t *testing.T) {
	hd := new(MockedHomeDir)
	hd.On("Path").Return("/some/path")

	w, _ := worklog.NewWorklog(hd)

	e := w.AddSite(worklog.NewSite("somesite"))
	l := w.SiteList()

	assert.Nil(t, e)
	assert.Len(t, l, 1)
	assert.Equal(t, l[0].Key, "somesite")
}
