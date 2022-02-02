// crio-top

package main

import (
	"golang/src/background"
	"golang/src/configuration"
	"golang/src/terminal"
	"io"
	"os"
	"time"
)

var output io.Writer = os.Stdout

func main() {
	// load configuration
	appConfiguration := configuration.New()

	// create shared structures to display data
	backgroundProcess := background.New(appConfiguration)

	// run background processes
	background.Execute(appConfiguration, &backgroundProcess)

	// display results
	for {
		// clear screen and return to the right-left corner
		terminal.Clear(output)

		// display date
		terminal.CreateDateLine(output, time.Now())

		// create empty line
		terminal.CreateEmptyLine(output)

		// display header
		terminal.PrintHeader(output, appConfiguration)

		// display data
		terminal.PrintRows(output, appConfiguration, backgroundProcess)

		// sleep for a defined duration
		time.Sleep(time.Duration(appConfiguration.Application.Refresh.Window) * time.Second)
	}
}
