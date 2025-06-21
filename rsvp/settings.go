package rsvp

import "time"

type Settings struct {
	// STATIC FILES
	STATIC_DIR            string
	STATIC_URL            string
	STATIC_COMPRESS       bool
	STATIC_BYTE_RANGE     bool
	STATIC_BROWSE         bool
	STATIC_DOWNLOAD       bool
	STATIC_INDEX          string
	STATIC_CACHE_DURATION time.Duration
	STATIC_MAX_AGE        int

	// TEMPLATES
	TEMPLATE_DIR       string
	TEMPLATE_EXTENSION string
}

func (s *Settings) BuildConf() *Settings {
	// STATIC DEFAULT CONF
	s.STATIC_DIR = "./static"
	s.STATIC_URL = "/static"
	s.STATIC_COMPRESS = true
	s.STATIC_BYTE_RANGE = true
	s.STATIC_BROWSE = false
	s.STATIC_DOWNLOAD = false
	s.STATIC_INDEX = "img/default.jpg"
	s.STATIC_CACHE_DURATION = 3600 * time.Second
	s.STATIC_MAX_AGE = 3600

	// TEMPLATES DEFAULT CONF
	s.TEMPLATE_DIR = "./templates"
	s.TEMPLATE_EXTENSION = ".html"

	return s
}
