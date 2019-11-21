package worklog

import (
	"github.com/chonla/homedir"
)

// App is worklog application
type App struct {
	home homedir.HomeWrapper
}

// NewApp to create new worklog application and initialize file storage to given path
func NewApp(homeDir homedir.HomeWrapper) (*App, error) {
	return &App{
		home: homeDir,
	}, nil
}

// StoragePath returns path that stores app content
func (a *App) StoragePath() string {
	return a.home.Path()
}

// CreateWorklog to create a new worklog instance
func (a *App) CreateWorklog() (*Worklog, error) {
	return NewWorklog(a.home)
}
