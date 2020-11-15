package interlocking

import (
	"github.com/giantswarm/microerror"
	"github.com/marcelmue/konstrukt/pkg/validate"
	"github.com/spf13/cobra"
)

const (
	flagFilename = "filename"

	flagHeight = "height"
	flagWidth  = "width"

	flagColor1 = "color1"
	flagColor2 = "color2"
	flagColor3 = "color3"
)

type flag struct {
	Filename string

	Height int
	Width  int

	Color1 string
	Color2 string
	Color3 string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Filename, flagFilename, "interlocking.svg", `Name of the output file.`)

	cmd.Flags().IntVar(&f.Height, flagHeight, 500, `Height of the output file in pixels.`)
	cmd.Flags().IntVar(&f.Width, flagWidth, 500, `Width of the output file in pixels.`)

	cmd.Flags().StringVar(&f.Color1, flagColor1, "#f7f1e3", `First color.`)
	cmd.Flags().StringVar(&f.Color2, flagColor2, "#fed330", `Second color.`)
	cmd.Flags().StringVar(&f.Color3, flagColor3, "#0fb9b1", `Third color.`)
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
	if !validate.Color(f.Color3) {
		validate.PrintFailure(flagColor3)
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
