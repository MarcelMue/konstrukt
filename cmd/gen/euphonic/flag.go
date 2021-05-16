package euphonic

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
	flagColor3 = "color3"
	flagColor4 = "color4"
	flagColor5 = "color5"
	flagColor6 = "color6"

	flagRandomize = "randomize"
)

type flag struct {
	Filename string

	Height int
	Width  int

	Color1 string
	Color2 string
	Color3 string
	Color4 string
	Color5 string
	Color6 string

	Randomize bool
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Filename, flagFilename, "euphonic.svg", `Name of the output file.`)

	cmd.Flags().IntVar(&f.Height, flagHeight, 500, `Height of the output file in pixels.`)
	cmd.Flags().IntVar(&f.Width, flagWidth, 500, `Width of the output file in pixels.`)

	cmd.Flags().StringVar(&f.Color1, flagColor1, "#e67e22", `First color.`)
	cmd.Flags().StringVar(&f.Color2, flagColor2, "#f6b93b", `Second color.`)
	cmd.Flags().StringVar(&f.Color3, flagColor3, "#e58e26", `Third color.`)
	cmd.Flags().StringVar(&f.Color4, flagColor4, "#2c3e50", `Fourth color.`)
	cmd.Flags().StringVar(&f.Color5, flagColor5, "#95a5a6", `Fifth color.`)
	cmd.Flags().StringVar(&f.Color4, flagColor6, "#7f8c8d", `Sixth color.`)

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
	if !validate.Color(f.Color3) {
		validate.PrintFailure(flagColor3)
		return microerror.Mask(invalidFlagError)
	}
	if !validate.Color(f.Color4) {
		validate.PrintFailure(flagColor4)
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
