package worklog

import (
	"github.com/chonla/homedir"
)

// Worklog is a worklog sheet
type Worklog struct {
	home  homedir.HomeWrapper
	sites []*Site
}

// NewWorklog to create a new worklog instance by giving worklog storage path
func NewWorklog(homeDir homedir.HomeWrapper) (*Worklog, error) {
	w := &Worklog{
		home:  homeDir,
		sites: []*Site{},
	}

	w.ReloadSites()

	return w, nil
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
	e := w.saveSite()
	return e
}

// ReloadSites read and store sites into internal sites
func (w *Worklog) ReloadSites() error {
	e := ReadFile(w.home.With("sites"), &w.sites)
	return e
}

// saveSite is to make sites persistent
func (w *Worklog) saveSite() error {
	e := WriteFile(w.home.With("sites"), w.sites)
	return e
}
