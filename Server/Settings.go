package main

import (
	"fmt"
	"io/ioutil"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type settings struct {
	Mutex          sync.RWMutex `yaml:"-" json:"-,omitempty"`
	File           string       `yaml:"-" json:"-,omitempty"`
	GUID           string       `yaml:"-" json:"-,omitempty"`
	UpdateLocation string       `yaml:"update_location"`
	UpdateInterval int32        `yaml:"update_interval"`
	Listener       string       `yaml:"listener" json:"listener,omitempty"`
	BaseURL        string       `yaml:"base_url"`
	Production     bool         `yaml:"production" json:"production,omitempty"`
	DBHost         string       `yaml:"db_host"`
	DBUser         string       `yaml:"db_user"`
	DBPassword     string       `yaml:"db_password"`
	DBBucket       string       `yaml:"db_bucket"`
}

func (s *settings) getSettings(location string) error {

	// General
	buf, err := ioutil.ReadFile(location)

	if err != nil {
		return fmt.Errorf("unable to read settings file: %v", err)
	}
	s.Mutex.Lock()
	s.File = location
	err = yaml.Unmarshal(buf, &s)
	s.Mutex.Unlock()

	if err != nil {
		return fmt.Errorf("unable to parse settings file: %v", err)
	}

	return nil

}
