package gen

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/marcelmue/konstrukt/cmd/gen/blockplay"
	"github.com/marcelmue/konstrukt/cmd/gen/blockplay2"
	"github.com/marcelmue/konstrukt/cmd/gen/euphonic"
	"github.com/marcelmue/konstrukt/cmd/gen/fallingdaggers"
	"github.com/marcelmue/konstrukt/cmd/gen/fiftyfive"
	"github.com/marcelmue/konstrukt/cmd/gen/hex22"
	"github.com/marcelmue/konstrukt/cmd/gen/hourglass"
	"github.com/marcelmue/konstrukt/cmd/gen/interlocking"
	"github.com/marcelmue/konstrukt/cmd/gen/interruptions"
	"github.com/marcelmue/konstrukt/cmd/gen/janein"
	"github.com/marcelmue/konstrukt/cmd/gen/ladysweat"
	"github.com/marcelmue/konstrukt/cmd/gen/modernhive"
	"github.com/marcelmue/konstrukt/cmd/gen/nolock"
	"github.com/marcelmue/konstrukt/cmd/gen/octolines"
	"github.com/marcelmue/konstrukt/cmd/gen/pantheon"
	"github.com/marcelmue/konstrukt/cmd/gen/qbert"
	"github.com/marcelmue/konstrukt/cmd/gen/quadrat"
	"github.com/marcelmue/konstrukt/cmd/gen/riviera"
	"github.com/marcelmue/konstrukt/cmd/gen/shining"
	"github.com/marcelmue/konstrukt/cmd/gen/swiss16"
	"github.com/marcelmue/konstrukt/cmd/gen/triangles"
	"github.com/marcelmue/konstrukt/cmd/gen/waves"
	"github.com/marcelmue/konstrukt/cmd/gen/whitegold"
	"github.com/marcelmue/konstrukt/cmd/gen/whitegold2"
	"github.com/marcelmue/konstrukt/cmd/gen/woozoo"
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

	var fiftyfiveCmd *cobra.Command
	{
		c := fiftyfive.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		fiftyfiveCmd, err = fiftyfive.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ladysweatCmd *cobra.Command
	{
		c := ladysweat.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ladysweatCmd, err = ladysweat.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var modernhiveCmd *cobra.Command
	{
		c := modernhive.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		modernhiveCmd, err = modernhive.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var swiss16Cmd *cobra.Command
	{
		c := swiss16.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		swiss16Cmd, err = swiss16.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var interlockingCmd *cobra.Command
	{
		c := interlocking.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		interlockingCmd, err = interlocking.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var whitegoldCmd *cobra.Command
	{
		c := whitegold.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		whitegoldCmd, err = whitegold.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var fallingdaggersCmd *cobra.Command
	{
		c := fallingdaggers.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		fallingdaggersCmd, err = fallingdaggers.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var whitegold2Cmd *cobra.Command
	{
		c := whitegold2.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		whitegold2Cmd, err = whitegold2.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var blockplayCmd *cobra.Command
	{
		c := blockplay.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		blockplayCmd, err = blockplay.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var blockplay2Cmd *cobra.Command
	{
		c := blockplay2.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		blockplay2Cmd, err = blockplay2.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var octolinesCmd *cobra.Command
	{
		c := octolines.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		octolinesCmd, err = octolines.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var qbertCmd *cobra.Command
	{
		c := qbert.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		qbertCmd, err = qbert.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var hourglassCmd *cobra.Command
	{
		c := hourglass.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		hourglassCmd, err = hourglass.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var wavesCmd *cobra.Command
	{
		c := waves.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		wavesCmd, err = waves.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var rivieraCmd *cobra.Command
	{
		c := riviera.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		rivieraCmd, err = riviera.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var nolockCmd *cobra.Command
	{
		c := nolock.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		nolockCmd, err = nolock.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var pantheonCmd *cobra.Command
	{
		c := pantheon.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		pantheonCmd, err = pantheon.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var hex22Cmd *cobra.Command
	{
		c := hex22.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		hex22Cmd, err = hex22.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var euphonicCmd *cobra.Command
	{
		c := euphonic.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		euphonicCmd, err = euphonic.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var woozooCmd *cobra.Command
	{
		c := woozoo.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		woozooCmd, err = woozoo.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var trianglesCmd *cobra.Command
	{
		c := triangles.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		trianglesCmd, err = triangles.New(c)
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
	c.AddCommand(fiftyfiveCmd)
	c.AddCommand(ladysweatCmd)
	c.AddCommand(modernhiveCmd)
	c.AddCommand(swiss16Cmd)
	c.AddCommand(interlockingCmd)
	c.AddCommand(whitegoldCmd)
	c.AddCommand(fallingdaggersCmd)
	c.AddCommand(whitegold2Cmd)
	c.AddCommand(blockplayCmd)
	c.AddCommand(blockplay2Cmd)
	c.AddCommand(octolinesCmd)
	c.AddCommand(qbertCmd)
	c.AddCommand(hourglassCmd)
	c.AddCommand(wavesCmd)
	c.AddCommand(rivieraCmd)
	c.AddCommand(nolockCmd)
	c.AddCommand(pantheonCmd)
	c.AddCommand(hex22Cmd)
	c.AddCommand(euphonicCmd)
	c.AddCommand(woozooCmd)
	c.AddCommand(trianglesCmd)

	return c, nil
}
