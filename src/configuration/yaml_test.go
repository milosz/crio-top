package configuration

import (
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	configuration := Application{
		Application: Configuration{
			Config: struct{ Filename string }{
				Filename: "../../examples/configuration.yaml",
			},
		},
	}

	configuration.ParseConfiguration()

	if configuration.Application.Header.Width.Server != 20 {
		t.Errorf("expected \"%d\", but got \"%d\"", 20, configuration.Application.Header.Width.Server)
	}

	if configuration.Application.Servers[0].Name != "desktop" {
		t.Errorf("expected \"%s\", but got \"%s\"", "desktop", configuration.Application.Servers[0].Name)
	}

	if configuration.Application.Commands[0].Command != "uname -m" {
		t.Errorf("expected \"%s\", but got \"%s\"", "uname -m", configuration.Application.Commands[0].Command)
	}

}
