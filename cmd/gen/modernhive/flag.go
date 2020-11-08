package modernhive

import (
	"github.com/spf13/cobra"
)

const (
	flagFilename = "filename"

	flagHeight = "height"
	flagWidth  = "width"

	flagColor1 = "color1"
	flagColor2 = "color2"
)

type flag struct {
	Filename string

	Height int
	Width  int

	Color1 string
	Color2 string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Filename, flagFilename, "modernhive.svg", `Name of the output file.`)

	cmd.Flags().IntVar(&f.Height, flagHeight, 500, `Height of the output file in pixels.`)
	cmd.Flags().IntVar(&f.Width, flagWidth, 500, `Width of the output file in pixels.`)

	cmd.Flags().StringVar(&f.Color1, flagColor1, "#00a8ff", `First color.`)
	cmd.Flags().StringVar(&f.Color2, flagColor2, "#ffffff", `Second color.`)
}

func (f *flag) Validate() error {

	return nil
}