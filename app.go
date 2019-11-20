package worklog

import "os"

// PathStat is statting file function alias
var PathStat = os.Stat

// ErrorFileNotExist is error checking function alias
var ErrorFileNotExist = os.IsNotExist

// App is worklog application
type App struct {
	storage string
}

// NewApp to create new worklog application and initialize file storage to given path
func NewApp(path string) (*App, error) {
	ok, e := verifyPath(path)
	if ok {
		return &App{
			storage: path,
		}, nil
	}
	return nil, e
}

func verifyPath(path string) (bool, error) {
	if _, e := PathStat(path); ErrorFileNotExist(e) {
		return false, e
	}
	return true, nil
}

// NewWorklog to create a new worklog instance
func (a *App) NewWorklog() (*Worklog, error) {
	return &Worklog{}, nil
}
