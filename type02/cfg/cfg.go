package cfg

import (
	"errors"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Cfg struct {
	TogglWorkspace  CfgTogglWorkspace `yaml:"toggl_workspace"`
	RestProjectName string            `yaml:"rest_project_name"`
}

type CfgTogglWorkspace struct {
	Token         string `yaml:"token"`
	WorkspaceName string `yaml:"workspace_name"`
}

func New(filename string) (t *Cfg, err error) {
	if !fileExists(filename) {
		return nil, errors.New("error: configfile not exists")
	}

	t = &Cfg{}
	buf, err := os.ReadFile(filename)
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
