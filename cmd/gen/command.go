package gen

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/marcelmue/konstrukt/cmd/gen/interruptions"
	"github.com/marcelmue/konstrukt/cmd/gen/janein"
	"github.com/marcelmue/konstrukt/cmd/gen/quadrat"
	"github.com/marcelmue/konstrukt/cmd/gen/shining"
)

const (
	name        = "gen"
	description = "Generate files."
)

type Config struct {
	Logger micrologger.Logger
	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	var err error

	var shiningCmd *cobra.Command
	{
		c := shining.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		shiningCmd, err = shining.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var interruptionsCmd *cobra.Command
	{
		c := interruptions.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		interruptionsCmd, err = interruptions.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var quadratCmd *cobra.Command
	{
		c := quadrat.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		quadratCmd, err = quadrat.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var janeinCmd *cobra.Command
	{
		c := janein.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		janeinCmd, err = janein.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
		RunE:  r.Run,
	}

	f.Init(c)

	c.AddCommand(shiningCmd)
	c.AddCommand(interruptionsCmd)
	c.AddCommand(quadratCmd)
	c.AddCommand(janeinCmd)

	return c, nil
}
