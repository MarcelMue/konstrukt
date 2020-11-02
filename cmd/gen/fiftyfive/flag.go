package fiftyfive

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
	cmd.Flags().StringVar(&f.Filename, flagFilename, "fiftyfive.svg", `Name of the output file.`)

	cmd.Flags().IntVar(&f.Height, flagHeight, 500, `Height of the output file in pixels.`)
	cmd.Flags().IntVar(&f.Width, flagWidth, 500, `Width of the output file in pixels.`)

	cmd.Flags().StringVar(&f.Color1, flagColor1, "#ff3838", `First color.`)
	cmd.Flags().StringVar(&f.Color2, flagColor2, "#000000", `Second color.`)
	cmd.Flags().StringVar(&f.Color3, flagColor3, "#ffffff", `Third color.`)
}

func (f *flag) Validate() error {

	return nil
}
