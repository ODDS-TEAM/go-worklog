package worklog_test

import (
	"testing"

	"github.com/chonla/worklog"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewWorklog(t *testing.T) {
	w, e := worklog.NewWorklog("/some/path")

	assert.Nil(t, e)
	assert.NotNil(t, w)
	assert.Equal(t, "/some/path", w.StoragePath())
}
