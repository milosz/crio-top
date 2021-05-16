// crio-top

package main

import (
	"golang/src/background"
	"golang/src/configuration"
	"golang/src/terminal"
	"time"
)

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
		terminal.Clear()

		// display date
		terminal.CreateDateLine()

		// create empty line
		terminal.CreateEmptyLine()

		// display header
		terminal.PrintHeader(appConfiguration)

		// display data
		terminal.PrintRows(appConfiguration, backgroundProcess)

		// sleep for a defined duration
		time.Sleep(time.Duration(appConfiguration.Application.Refresh.Window) * time.Second)
	}
}
