package configuration

import (
	"flag"
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	os.Args = []string{"crio-top", "--configuration", "desktop-configuration.yml"}

	appConfiguration := Application{}

	appConfiguration.ParseFlags()

	if appConfiguration.Application.Config.Filename != "desktop-configuration.yml" {
		t.Errorf("expected \"%s\", but got \"%s\"", "desktop-configuration.yml", appConfiguration.Application.Config.Filename)
	}
}

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

func TestInitialize(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"crio-top", "--configuration", "../../examples/configuration.yaml"}

	var appConfiguration = Initialize()

	if appConfiguration.Application.Config.Filename != "../../examples/configuration.yaml" {
		t.Errorf("expected \"%s\", but got \"%s\"", "../../examples/configuration.yaml", appConfiguration.Application.Config.Filename)
	}
}
