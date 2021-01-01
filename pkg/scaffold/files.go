package scaffold

var PatternCommandTemplate = `package {{ .Name }}

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name        = "{{ .Name }}"
	description = "Draws {{ .Author }}s '{{ .Name }}'."
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

	return c, nil
}
`

var PatternCommandTestTemplate = `package {{ .Name }}

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/giantswarm/micrologger"
)

func Test_command(t *testing.T) {
	testCases := []struct {
		name     string
		c1       string
		c2       string
		filename string
		height   int
		width    int
	}{
		{
			name:     "case 0: default pattern",
			filename: "{{ .Name }}",
		},
		{
			name:     "case 1: in color pattern",
			filename: "{{ .Name }}-in",
			c1:       "#e55039",
			c2:       "#f39c12",
		},
		{
			name:     "case 2: banner resize",
			filename: "{{ .Name }}-wide",
			width:    2000,
			height:   400,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			var err error

			var logger micrologger.Logger
			{
				c := micrologger.Config{}

				logger, err = micrologger.New(c)
				if err != nil {
					t.Fatal("expected", nil, "got", err)
				}
			}

			c := Config{
				Logger: logger,
			}

			testCommand, err := New(c)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			args := []string{}
			if tc.filename != "" {
				args = append(args, []string{"--filename", fmt.Sprintf("../../../samples/%s.svg", tc.filename)}...)
			}
			if tc.c1 != "" {
				args = append(args, []string{"--color1", tc.c1}...)
			}
			if tc.c2 != "" {
				args = append(args, []string{"--color2", tc.c2}...)
			}
			if tc.height != 0 {
				args = append(args, []string{"--height", strconv.Itoa(tc.height)}...)
			}
			if tc.width != 0 {
				args = append(args, []string{"--width", strconv.Itoa(tc.width)}...)
			}

			testCommand.SetArgs(args)

			err = testCommand.Execute()
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
		})
	}
}
`

var PatternErrorTemplate = `package {{ .Name }}

import "github.com/giantswarm/microerror"

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidFlagError = &microerror.Error{
	Kind: "invalidFlagError",
}

// IsInvalidFlag asserts invalidFlagError.
func IsInvalidFlag(err error) bool {
	return microerror.Cause(err) == invalidFlagError
}
`

var PatternFlagTemplate = `package {{ .Name }}

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/marcelmue/konstrukt/pkg/validate"
)

const (
	flagFilename = "filename"

	flagHeight = "height"
	flagWidth  = "width"

	flagColor1 = "color1"
	flagColor2 = "color2"

	flagRandomize = "randomize"
)

type flag struct {
	Filename string

	Height int
	Width  int

	Color1 string
	Color2 string

	Randomize bool
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Filename, flagFilename, "{{ .Name }}.svg", "Name of the output file.")

	cmd.Flags().IntVar(&f.Height, flagHeight, 500, "Height of the output file in pixels.")
	cmd.Flags().IntVar(&f.Width, flagWidth, 500, "Width of the output file in pixels.")

	cmd.Flags().StringVar(&f.Color1, flagColor1, "#bdc581", "First color.")
	cmd.Flags().StringVar(&f.Color2, flagColor2, "#2c3a47", "Second color.")

	cmd.Flags().BoolVar(&f.Randomize, flagRandomize, false, "Randomize all colors in the pattern, ignore other color flags.")
}

func (f *flag) Validate() error {
	if !validate.Color(f.Color1) {
		validate.PrintFailure(flagColor1)
		return microerror.Mask(invalidFlagError)
	}
	if !validate.Color(f.Color2) {
		validate.PrintFailure(flagColor2)
		return microerror.Mask(invalidFlagError)
	}
	if !validate.Size(f.Height) {
		validate.PrintFailure(flagHeight)
		return microerror.Mask(invalidFlagError)
	}
	if !validate.Size(f.Width) {
		validate.PrintFailure(flagWidth)
		return microerror.Mask(invalidFlagError)
	}

	return nil
}
`

var PatternRunnerTemplate = `package {{ .Name }}

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	svg "github.com/marcelmue/svgo"
	"github.com/spf13/cobra"

	"github.com/marcelmue/konstrukt/pkg/color"
	"github.com/marcelmue/konstrukt/pkg/project"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	if r.flag.Randomize {
		r.flag.Color1 = color.Random()
		r.flag.Color2 = color.Random()
		fmt.Printf("Set Color1:%s Color2:%s\n", r.flag.Color1, r.flag.Color2)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	f, err := os.Create(r.flag.Filename)
	if err != nil {
		return microerror.Mask(err)
	}
	width, height := r.flag.Width, r.flag.Height
	c1, c2 := r.flag.Color1, r.flag.Color2

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Add your implementation here!
	canvas.Rect(0, 0, width, height, "fill:"+c1)


	canvas.End()

	return nil
}
`
