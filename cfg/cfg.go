package cfg

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Cfg struct {
	Slack CfgSlack `yaml:"slack"`
	Toggl CfgToggl `yaml:"toggl"`
}

type CfgSlack struct {
	Token   string `yaml:"token"`
	Channel string `yaml:"channel"`
}

type CfgToggl struct {
	Token     string `yaml:"token"`
	Workspace string `yaml:"workspace"`
	Client    string `yaml:"client"`
	Project   string `yaml:"project"`
}

func New(filename string) (t *Cfg, err error) {
	if !fileExists(filename) {
		return nil, errors.New("error: configfile not exists")
	}

	t = &Cfg{}
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buf, t)
	if err != nil {
		return nil, err
	}
	log.Printf("debug: [Read] %s", filename)
	return t, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
