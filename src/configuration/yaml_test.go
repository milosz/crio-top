package configuration

import (
	"testing"
)

func TestParseConfiguration(t *testing.T) {
	appConfiguration := Application{
		Application: Configuration{
			Config: struct{ Filename string }{
				Filename: "../../examples/configuration.yaml",
			},
		},
	}

	appConfiguration.ParseConfiguration()

	if appConfiguration.Application.Header.Width.Server != 20 {
		t.Errorf("expected \"%d\", but got \"%d\"", 20, appConfiguration.Application.Header.Width.Server)
	}

	if appConfiguration.Application.Servers[0].Name != "desktop" {
		t.Errorf("expected \"%s\", but got \"%s\"", "desktop", appConfiguration.Application.Servers[0].Name)
	}

	if appConfiguration.Application.Commands[0].Command != "uname -m" {
		t.Errorf("expected \"%s\", but got \"%s\"", "uname -m", appConfiguration.Application.Commands[0].Command)
	}

}
