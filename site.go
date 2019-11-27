package worklog

// Site is working site
type Site struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

// NewSite creates a simple dead site object
func NewSite(url, username, token string) *Site {
	return &Site{
		URL:      url,
		Username: username,
		Token:    token,
	}
}
