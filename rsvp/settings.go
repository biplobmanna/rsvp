package rsvp

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

// GLOBAL "SETTINGS" variable
// holds configured settings after app has intialized
var SETTINGS Settings

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

	// SECRETS
	SECRET_KEY  string
	ADMIN_TOKEN string
}

func (s *Settings) BuildConf() *Settings {
	// Read from dotEnv
	dotEnvMap := ParseDotEnv()

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

	// SECRETS
	s.SECRET_KEY = dotEnvMap["SECRET_KEY"]
	s.ADMIN_TOKEN = dotEnvMap["ADMIN_TOKEN"]

	return s
}

func ParseDotEnv() map[string]string {
	m := make(map[string]string)

	// open the dotEnv file
	f, err := os.Open("./.env")
	if err != nil {
		log.Fatal(err)
		return m
	}

	// close when all reading is done
	defer f.Close()

	// scanner to read line-by-line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// ignore comments with #
		// ignore blank lines or lines with only spaces
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		// split on `=` into key, value pairs
		list := strings.SplitN(line, "=", 2)
		// trim leading, and trailing spaces and add to the hashmap
		m[strings.TrimSpace(list[0])] = strings.TrimSpace(list[1])
	}

	// if any error encountered while reading dotEnv
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return m
}
