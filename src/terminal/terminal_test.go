package terminal

import (
	"bytes"
	"golang/src/background"
	"golang/src/configuration"
	"io"
	"testing"
	"time"
)

func TestClear(t *testing.T) {
	var output io.Writer = new(bytes.Buffer)

	Clear(output)

	var got string = output.(*bytes.Buffer).String()
	var expected string = "\033[2J\033[H"

	if got != expected {
		t.Errorf("expected \"%s\", but got \"%s\"", expected, got)
	}
}

func TestCreateDateLine(t *testing.T) {
	var output io.Writer = new(bytes.Buffer)
	var datetime time.Time = time.Date(2000, time.January, 19, 0, 1, 0, 0, time.UTC)

	CreateDateLine(output, datetime)

	var got string = output.(*bytes.Buffer).String()
	var expected string = "Date: " + (datetime).Format(time.UnixDate)

	if got != expected {
		t.Errorf("expected \"%s\", but got \"%s\"", expected, got)
	}
}

func TestCreateEmptyLine(t *testing.T) {
	var output io.Writer = new(bytes.Buffer)

	CreateEmptyLine(output)

	var got string = output.(*bytes.Buffer).String()
	var expected string = "\n\n"

	if got != expected {
		t.Errorf("expected \"%s\", but got \"%s\"", expected, got)
	}
}

func TestPrintHeader(t *testing.T) {
	var output io.Writer = new(bytes.Buffer)

	appConfiguration := &configuration.Application{
		Application: configuration.Configuration{
			Header: struct {
				Width struct{ Server int }
			}{
				Width: struct{ Server int }{Server: 10},
			},
			Commands: []configuration.CommandConfiguration{
				{
					Name:  "0123456789",
					Width: 10,
				},
				{
					Name:  "Command001",
					Width: 10,
				},
				{
					Name:  "Command 02",
					Width: 10,
				},
				{
					Name:  "C3",
					Width: 5,
				},
			},
		},
	}

	PrintHeader(output, appConfiguration)

	var got string = output.(*bytes.Buffer).String()
	var expected string = "    Server 0123456789 Command001 Command 02    C3 \n"

	if got != expected {
		t.Errorf("expected \"%s\", but got \"%s\"", expected, got)
	}
}

func TestPrintRows(t *testing.T) {
	var output io.Writer = new(bytes.Buffer)

	appConfiguration := &configuration.Application{
		Application: configuration.Configuration{
			Header: struct {
				Width struct{ Server int }
			}{
				Width: struct{ Server int }{Server: 10},
			},
			Commands: []configuration.CommandConfiguration{
				{
					Width: 5,
				},
				{
					Width: 5,
				},
				{
					Width: 5,
				},
				{
					Width: 5,
				},
				{
					Width: 5,
				},
				{
					Width: 5,
				},
			},
		},
	}

	serverOutput := background.Output{
		Servers: []background.ServerOutput{
			{
				Server: configuration.ServerConfiguration{
					Name: "server name",
				},
				Values: []string{
					"100%", "30%", "10M", "1G", "", "-",
				},
			},
		},
	}

	PrintRows(output, appConfiguration, serverOutput)

	var got string = output.(*bytes.Buffer).String()
	var expected string = "server name  100%   30%   10M    1G           - \n"

	if got != expected {
		t.Errorf("expected \"%s\", but got \"%s\"", expected, got)
	}
}
