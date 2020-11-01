package shining

import (
	"github.com/spf13/cobra"
)

const (
	flagFilename = "filename"

	flagHeight = "height"
	flagWidth  = "width"

	flagAccent     = "accent"
	flagBackground = "background"
	flagForeground = "foreground"
)

type flag struct {
	Filename string

	Height int
	Width  int

	Accent     string
	Background string
	Foreground string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Filename, flagFilename, "shining.svg", `Name of the output file.`)

	cmd.Flags().IntVar(&f.Height, flagHeight, 500, `Height of the output file in pixels.`)
	cmd.Flags().IntVar(&f.Width, flagWidth, 500, `Width of the output file in pixels.`)

	cmd.Flags().StringVar(&f.Accent, flagAccent, "rgb(153,29,40)", `Color of the accent.`)
	cmd.Flags().StringVar(&f.Background, flagBackground, "rgb(227,78,25)", `Color of the background.`)
	cmd.Flags().StringVar(&f.Foreground, flagForeground, "rgb(65,52,44)", `Color of the foreground.`)
}

func (f *flag) Validate() error {

	return nil
}
