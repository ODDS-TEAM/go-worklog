package worklog

// Worklog is a worklog sheet
type Worklog struct {
	storage string
}

// NewWorklog to create a new worklog instance by giving worklog storage path
func NewWorklog(path string) (*Worklog, error) {
	return &Worklog{
		storage: path,
	}, nil
}

// StoragePath returns path that stores worklog content
func (w *Worklog) StoragePath() string {
	return w.storage
}
