package interruptions

import (
	"context"
	"io"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
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

	c1 := r.flag.Color1
	c2 := r.flag.Color2
	c3 := r.flag.Color3

	width, height := r.flag.Width, r.flag.Height

	canvas := svg.New(f)
	canvas.Start(width, height)
	ypos := []int{0, -80, -75, -5, -145, -100, -50, -10, -130, -120, -20}
	ycol := []string{c1, c1, c2, c3, c3, c2, c2, c1, c3, c3, c2, c1, c2}
	ypositioncounter := 0
	for x := 5; x < width; x += 15 {
		xpositioncounter := 0
		for y := ypos[ypositioncounter%len(ypos)]; y < height; y += 160 {
			canvas.Rect(x, y, 5, 150, "fill: "+ycol[(xpositioncounter+ypositioncounter)%len(ycol)])
			xpositioncounter++
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
