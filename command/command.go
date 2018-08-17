// Package command implements the root command for the command line tool.
package command

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/marcelmue/konstrukt/command/create"
)

// Config represents the configuration used to create a new root command.
type Config struct {
	// Dependencies.
	Logger micrologger.Logger

	// Settings.
	Description  string
	GitCommit    string
	Name         string
	Source       string
	DisableCache bool
}

// DefaultConfig provides a default configuration to create a new root command
// by best effort.
func DefaultConfig() Config {
	return Config{
		// Dependencies.
		Logger: nil,

		// Settings.
		Description:  "",
		GitCommit:    "",
		Name:         "",
		Source:       "",
		DisableCache: false,
	}
}

// New creates a new root command.
func New(config Config) (*Command, error) {
	var err error

	var createCommand *create.Command
	{
		commandConfig := create.DefaultConfig()

		commandConfig.Logger = config.Logger

		createCommand, err = create.New(commandConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newCommand := &Command{
		// Internals.
		cobraCommand: nil,
	}

	newCommand.cobraCommand = &cobra.Command{
		Use:   config.Name,
		Short: config.Description,
		Long:  config.Description,
		Run:   newCommand.Execute,
	}

	newCommand.cobraCommand.AddCommand(createCommand.CobraCommand())

	return newCommand, nil
}

type Command struct {
	// Internals.
	cobraCommand *cobra.Command
}

// CobraCommand returns the spf13/cobra command
func (c *Command) CobraCommand() *cobra.Command {
	return c.cobraCommand
}

// Execute is called to actuall run the main command
func (c *Command) Execute(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, nil)
}
