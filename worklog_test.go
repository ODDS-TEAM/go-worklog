package worklog_test

import (
	"os"
	"testing"

	worklog "github.com/ODDS-TEAM/go-worklog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewWorklog(t *testing.T) {
	hd := new(MockedHomeDir)
	hd.On("Path").Return("/some/path")
	hd.On("With", mock.Anything).Return("")

	w, e := worklog.NewWorklog(hd)

	assert.Nil(t, e)
	assert.NotNil(t, w)
	assert.Equal(t, "/some/path", w.StoragePath())
}

func TestGettingSiteList(t *testing.T) {
	hd := new(MockedHomeDir)
	hd.On("Path").Return("/some/path")
	hd.On("With", mock.Anything).Return(".sitelist")

	worklog.FileRead = func(filename string) ([]byte, error) {
		return []byte{}, nil
	}

	w, _ := worklog.NewWorklog(hd)

	l := w.SiteList()

	assert.Len(t, l, 0)
}

func TestAddingNewSiteToListShouldBeSuccess(t *testing.T) {
	temp := []byte{}

	hd := new(MockedHomeDir)
	hd.On("Path").Return("/some/path")
	hd.On("With", "sites").Return("./sites")

	worklog.FileRead = func(filename string) ([]byte, error) {
		return temp, nil
	}
	worklog.FileWrite = func(filename string, data []byte, perm os.FileMode) error {
		temp = data
		return nil
	}

	w, _ := worklog.NewWorklog(hd)

	e := w.AddSite(worklog.NewSite("somesite", "user", "tok"))

	w2, _ := worklog.NewWorklog(hd)

	l := w2.SiteList()

	assert.Nil(t, e)
	assert.Len(t, l, 1)
	assert.Equal(t, l[0].URL, "somesite")
	assert.Equal(t, l[0].Username, "user")
	assert.Equal(t, l[0].Token, "tok")
}
