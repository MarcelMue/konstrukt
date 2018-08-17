package main

import (
	"github.com/giantswarm/micrologger"

	"github.com/marcelmue/konstrukt/command"
)

var (
	description string = "Command line tool for me."
	gitCommit   string = "n/a"
	name        string = "konstrukt"
	source      string = "https://github.com/marcelmue/konstrukt"
)

func main() {
	var err error

	// Create a new logger which is used by all packages.
	var newLogger micrologger.Logger
	{
		newLogger, err = micrologger.New(micrologger.Config{})
		if err != nil {
			panic(err)
		}
	}

	var newCommand *command.Command
	{
		commandConfig := command.DefaultConfig()

		commandConfig.Logger = newLogger

		commandConfig.Description = description
		commandConfig.GitCommit = gitCommit
		commandConfig.Name = name
		commandConfig.Source = source

		newCommand, err = command.New(commandConfig)
		if err != nil {
			panic(err)
		}
	}

	newCommand.CobraCommand().Execute()
}
