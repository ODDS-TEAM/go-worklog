package worklog

import "github.com/chonla/homedir"

// Worklog is a worklog sheet
type Worklog struct {
	home homedir.HomeWrapper
}

// NewWorklog to create a new worklog instance by giving worklog storage path
func NewWorklog(homeDir homedir.HomeWrapper) (*Worklog, error) {
	return &Worklog{
		home: homeDir,
	}, nil
}

// StoragePath returns path that stores worklog content
func (w *Worklog) StoragePath() string {
	return w.home.Path()
}
