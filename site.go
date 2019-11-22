package worklog

// Site is working site
type Site struct {
	Key string
}

// NewSite creates a simple dead site object
func NewSite(key string) *Site {
	return &Site{
		Key: key,
	}
}
