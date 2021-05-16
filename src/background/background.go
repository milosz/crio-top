// background operations

package background

import (
	"golang/src/configuration"
	"strings"
	"time"
)

// shared structure to store data
type Output struct {
	Servers []ServerOutput
}

type ServerOutput struct {
	Server configuration.ServerConfiguration
	Values []string
}

// initialize shared structures
func New(configuration *configuration.Application) (output Output) {
	for i := range configuration.Application.Servers {
		newServer := configuration.Application.Servers[i]

		var values []string
		for range configuration.Application.Commands {
			values = append(values, "")
		}
		output.Servers = append(output.Servers, ServerOutput{Server: newServer, Values: values})
	}
	return output
}

// check data in the background
func Execute(configuration *configuration.Application, output *Output) {
	for i := range output.Servers {
		go ExecuteOnServer(configuration, &(*output).Servers[i])
	}
}

// check data
func ExecuteOnServer(configuration *configuration.Application, output *ServerOutput) {
	for {
		SSHConnect(&output.Server)
		for i := range configuration.Application.Commands {
			SSHCreateSession(&output.Server)
			out := SSHExec(&output.Server, configuration.Application.Commands[i].Command)
			output.Values[i] = strings.Trim(out, "\n")
			SSHCloseSession(&output.Server)
		}
		SSHClose(&output.Server)

		time.Sleep(time.Duration(configuration.Application.Refresh.Data) * time.Second)
	}
}
