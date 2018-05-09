package config

import "github.com/callum-ramage/jsonconfig"

var (
	Config jsonconfig.Configuration
)

func Load(filename string) error {
	var err error
	Config, err = jsonconfig.LoadAbstract(filename, "")
	if err != nil {
		return err
	}
	return nil
}