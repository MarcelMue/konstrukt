package shining

import (
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
	cmd.Flags().StringVar(&f.Filename, flagFilename, "shining.svg", `Name of the output file.`)

	cmd.Flags().IntVar(&f.Height, flagHeight, 500, `Height of the output file in pixels.`)
	cmd.Flags().IntVar(&f.Width, flagWidth, 500, `Width of the output file in pixels.`)

	cmd.Flags().StringVar(&f.Color1, flagColor1, "rgb(153,29,40)", `Color of the accent.`)
	cmd.Flags().StringVar(&f.Color2, flagColor2, "rgb(227,78,25)", `Color of the background.`)
	cmd.Flags().StringVar(&f.Color3, flagColor3, "rgb(65,52,44)", `Color of the foreground.`)
}

func (f *flag) Validate() error {

	return nil
}
