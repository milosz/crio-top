package configuration

import (
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	os.Args = []string{"crio-top", "-configuration", "desktop-configuration.yml"}

	appConfiguration := Application{}

	appConfiguration.ParseFlags()

	if appConfiguration.Application.Config.Filename != "desktop-configuration.yml" {
		t.Errorf("expected \"%s\", but got \"%s\"", "desktop-configuration.yml", appConfiguration.Application.Config.Filename)
	}
}
