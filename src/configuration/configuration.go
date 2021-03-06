// general configuration
package configuration

import (
	"flag"

	"golang.org/x/crypto/ssh"
)

// configuration structures
type Configuration struct {
	Config struct {
		Filename string
	}
	Refresh struct {
		Window int
		Data   int
	}
	Header struct {
		Width struct {
			Server int
		}
	}
	Servers  []ServerConfiguration
	Commands []CommandConfiguration
}

type ServerConfiguration struct {
	Name    string
	Server  string
	User    string
	Port    int
	Session *ssh.Session
	Client  *ssh.Client
}

type CommandConfiguration struct {
	Name    string
	Command string
	Width   int
}

type Application struct {
	Application Configuration
}

// initialize configuration structures
func Initialize() (configuration *Application) {
	configuration = &Application{}

	configuration.ParseFlags()
	configuration.ParseConfiguration()

	return configuration
}

func (configuration *Application) ParseFlags() {
	flag.StringVar(&configuration.Application.Config.Filename, "configuration", "configuration.yaml", "path to a configuration file")
	flag.Parse()
}
