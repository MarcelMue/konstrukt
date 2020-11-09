package interlocking

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
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Polygon
	xp := []int{0, 20, 100, 60, 40, 60}
	yp := []int{20, 0, 80, 120, 100, 80}

	canvas.Def()
	canvas.Gid("unit")
	canvas.Polygon(xp, yp)
	canvas.Gend()
	canvas.Gid("runit")
	canvas.TranslateRotate(40, 160, 180)
	canvas.Use(0, 0, "#unit")
	canvas.Gend()
	canvas.Gend()
	canvas.DefEnd()

	ypositioncounter := 0
	for y := -80; y < width+80; y += 80 {
		for x := -160; x < height+160; x += 160 {
			if ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", "fill:"+c2)
				canvas.Use(x, y, "#runit", "fill:"+c2)
			}
			if ypositioncounter%2 == 1 {
				canvas.Use(x, y, "#unit", "fill:"+c3)
				canvas.Use(x, y, "#runit", "fill:"+c3)
			}
		}
		ypositioncounter++
	}

	canvas.End()

	return nil
}
