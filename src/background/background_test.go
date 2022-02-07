package background

import (
	"golang/src/configuration"
	"strings"
	"sync"
	"testing"
)

func TestInitialize(t *testing.T) {
	appConfiguration := &configuration.Application{
		Application: configuration.Configuration{
			Servers: []configuration.ServerConfiguration{
				{
					Name: "server x",
				},
			},
			Commands: []configuration.CommandConfiguration{
				{
					Name: "command 1",
				},
				{
					Name: "command 2",
				},
			},
		},
	}

	var backgroundProcess Output = Initialize(appConfiguration)

	if backgroundProcess.Servers[0].Server.Name != "server x" {
		t.Errorf("expected \"%s\", but got \"%s\"", "server x", backgroundProcess.Servers[0].Server.Name)
	}

	if len(backgroundProcess.Servers[0].Values) != 2 {
		t.Errorf("expected \"%d\", but got \"%d\"", 2, len(backgroundProcess.Servers[0].Values))
	}
}

type ServerExecutionMock struct {
	wg           sync.WaitGroup
	serverCalls  int
	commandCalls int
}

func (d *ServerExecutionMock) ExecuteOnServer(configuration *configuration.Application, soutput *ServerOutput) {
	defer d.wg.Done()
	d.serverCalls += 1
	for k := range configuration.Application.Commands {
		d.commandCalls += 1
		soutput.Values[k] = strings.Join([]string{soutput.Server.Name, configuration.Application.Commands[k].Name}, " ")
	}
}

func TestExecute(t *testing.T) {
	appConfiguration := &configuration.Application{
		Application: configuration.Configuration{
			Servers: []configuration.ServerConfiguration{
				{
					Name: "server 1",
				},
				{
					Name: "server 2",
				},
			},
			Commands: []configuration.CommandConfiguration{
				{
					Name: "command 1",
				},
				{
					Name: "command 2",
				},
				{
					Name: "command 3",
				},
			},
		},
	}

	var output Output = Initialize(appConfiguration)

	dd := &ServerExecutionMock{}

	dd.wg.Add(len(output.Servers))
	Execute(appConfiguration, &output, dd)
	dd.wg.Wait()

	if strings.Join(output.Servers[0].Values, ";") != "server 1 command 1;server 1 command 2;server 1 command 3" {
		t.Errorf("expected \"%s\", but got \"%s\"", "server 1 command 1;server 1 command 2;server 1 command 3", strings.Join(output.Servers[0].Values, ";"))
	}

	if dd.serverCalls != len(output.Servers) {
		t.Errorf("expected \"%d\", but got \"%d\"", len(output.Servers), dd.serverCalls)
	}

	if dd.commandCalls != len(output.Servers)*len(appConfiguration.Application.Commands) {
		t.Errorf("expected \"%d\", but got \"%d\"", len(output.Servers)*len(appConfiguration.Application.Commands), dd.commandCalls)
	}
}
