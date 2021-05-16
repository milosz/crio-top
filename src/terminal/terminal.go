// terminal related helpers and functions
package terminal

import (
	"fmt"
	"golang/src/background"
	"golang/src/configuration"
	"strconv"
	"time"
)

// clear the screen and go to the top left corner
func Clear() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

// create a date line
func CreateDateLine() {
	date := time.Now()
	fmt.Println("Date: " + date.Format(time.UnixDate))
}

// create an empty line
func CreateEmptyLine() {
	fmt.Println()
}

// display header
func PrintHeader(configuration *configuration.Application) {
	fmt.Printf("%"+strconv.Itoa(configuration.Application.Header.Width.Server)+"s ", "Server")

	for i := range configuration.Application.Commands {
		fmt.Printf("%"+strconv.Itoa(configuration.Application.Commands[i].Width)+"s ", configuration.Application.Commands[i].Name)
	}
	fmt.Println()

}

// display rows
func PrintRows(applicationConfiguration *configuration.Application, output background.Output) {
	for i := range output.Servers {
		fmt.Printf("%"+strconv.Itoa(applicationConfiguration.Application.Header.Width.Server)+"s ", output.Servers[i].Server.Name)
		for j := range applicationConfiguration.Application.Commands {
			fmt.Printf("%"+strconv.Itoa(applicationConfiguration.Application.Commands[j].Width)+"s ", output.Servers[i].Values[j])
		}
		fmt.Println()
	}
}
