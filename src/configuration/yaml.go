// YAML configuration
package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// uses configuration structures
type YamlConfiguration struct {
	Application Configuration
}

// parse YAML file
func (configuration *Application) ParseConfiguration() {
	file, err := ioutil.ReadFile(configuration.Application.Config.Filename)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(file, &configuration)
	if err != nil {
		log.Fatal(err)
	}
}
