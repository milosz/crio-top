// terminal related helpers and functions
package terminal

import (
	"fmt"
	"golang/src/background"
	"golang/src/configuration"
	"io"
	"strconv"
	"time"
)

// clear the screen and go to the top left corner
func Clear(w io.Writer) {
	fmt.Fprintf(w, "\033[2J")
	fmt.Fprintf(w, "\033[H")
}

// create a date line
func CreateDateLine(w io.Writer, datetime time.Time) {
	fmt.Fprintf(w, "Date: "+datetime.Format(time.UnixDate))
}

// create an empty line
func CreateEmptyLine(w io.Writer) {
	fmt.Fprintf(w, "\n\n")
}

// display header
func PrintHeader(w io.Writer, configuration *configuration.Application) {
	fmt.Fprintf(w, "%"+strconv.Itoa(configuration.Application.Header.Width.Server)+"s ", "Server")

	for i := range configuration.Application.Commands {
		fmt.Fprintf(w, "%"+strconv.Itoa(configuration.Application.Commands[i].Width)+"s ", configuration.Application.Commands[i].Name)
	}
	fmt.Fprintf(w, "\n")

}

// display rows
func PrintRows(w io.Writer, applicationConfiguration *configuration.Application, output background.Output) {
	for i := range output.Servers {
		fmt.Fprintf(w, "%"+strconv.Itoa(applicationConfiguration.Application.Header.Width.Server)+"s ", output.Servers[i].Server.Name)
		for j := range applicationConfiguration.Application.Commands {
			fmt.Fprintf(w, "%"+strconv.Itoa(applicationConfiguration.Application.Commands[j].Width)+"s ", output.Servers[i].Values[j])
		}
		fmt.Fprintf(w, "\n")
	}
}
