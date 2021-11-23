package asab

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	// This is a global configuration (also accessible by `App.Config`)
	// It is a global because application can have only one configuration
	// and the config object has to exists before the application is initialized
	Config *ini.File
)

func init() {
	Config = ini.Empty()

	AddConfigDefaults("general", map[string]string{
		"config_file": os.Getenv("ASAB_CONFIG"),
	})

}

func AddConfigDefaults(section string, values map[string]string) {

	// Ensure that the section exists
	s, err := Config.GetSection(section)
	if err != nil {
		s, err = Config.NewSection(section)
		if err != nil {
			panic(err)
		}
	}

	// Set provided values
	for k, v := range values {
		s.NewKey(k, v)
	}
}

func _LoadConfig() *ini.File {

	cfgsec, err := Config.GetSection("general")
	if err != nil {
		panic(err)
	}

	config_file := cfgsec.Key("config_file").String()
	if config_file != "" {
		err = Config.Append(config_file)
		if err != nil {
			log.Errorf("Cannot load configuration file '%s'", config_file)
		}
	}

	return Config
}
