package worklog

import "github.com/chonla/homedir"

// Worklog is a worklog sheet
type Worklog struct {
	home  homedir.HomeWrapper
	sites []*Site
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

// SiteList returns site stored in local storage
func (w *Worklog) SiteList() []*Site {
	return w.sites
}

// AddSite add a given site to site list
func (w *Worklog) AddSite(s *Site) error {
	w.sites = append(w.sites, s)
	return nil
}
