package internal

import (
	"fmt"
	"io"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const (
	dirName string = ".ccoio"
	fileFmt string = "ccoio%s.yml"
)

type AccessSettings struct {
	ClientID     string   `yaml:"clientID"`
	ClientSecret string   `yaml:"clientSecret"`
	Scopes       []string `yaml:"scopes"`
	TokenURL     string   `yaml:"tokenURL"`
}

func LoadAccessSettings(envName string) (*AccessSettings, error) {
	// determine path
	h, _ := os.UserHomeDir()
	basePath := path.Join(h, dirName)

	// determine file name
	fn := fmt.Sprintf(fileFmt, "")
	if envName != "" && envName != "production" {
		fn = fmt.Sprintf(fileFmt, fmt.Sprintf(".%s", envName))
	}

	fp := path.Join(basePath, fn)
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	yml, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var as AccessSettings
	if err := yaml.Unmarshal(yml, &as); err != nil {
		return nil, err
	}

	return &as, nil
}
