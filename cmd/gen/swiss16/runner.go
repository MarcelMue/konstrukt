package swiss16

import (
	"context"
	"fmt"
	"io"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/marcelmue/konstrukt/pkg/color"
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
		r.flag.Color3 = color.Random()
		fmt.Printf("Set Color1:%s Color2:%s Color3:%s\n", r.flag.Color1, r.flag.Color2, r.flag.Color3)
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
	c1, c2, c3 := r.flag.Color1, r.flag.Color2, r.flag.Color3

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Polygon.
	xp := []int{0, 0, 40, 60}
	yp := []int{0, 50, 90, 60}

	canvas.Def()
	canvas.Gid("unit")
	canvas.Polygon(xp, yp)
	canvas.Gend()
	canvas.Gid("runit")
	canvas.ScaleXY(-1, 1)
	canvas.Use(0, 0, "#unit")
	canvas.Gend()
	canvas.Gend()
	canvas.DefEnd()

	ypositioncounter := 0
	for y := -60; y < height; y += 60 {
		for x := -120; x < width; x += 120 {
			if ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", "fill:"+c2)
				canvas.Use(x, y, "#runit", "fill:"+c3)
			}
			if ypositioncounter%2 == 1 {
				canvas.Use(x+60, y, "#unit", "fill:"+c2)
				canvas.Use(x+60, y, "#runit", "fill:"+c3)
			}
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
