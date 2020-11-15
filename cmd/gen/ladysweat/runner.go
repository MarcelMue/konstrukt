package ladysweat

import (
	"context"
	"fmt"
	"io"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/marcelmue/konstrukt/pkg/color"
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

	// Skewed rectangle.
	xp := []int{50, 70, 80, 60}
	yp := []int{50, 50, 30, 30}
	// Polygon.
	xt := []int{45, 80, 85, 70, 80, 60}
	yt := []int{30, 30, 20, 20, 0, 0}

	canvas.Def()
	canvas.Gid("unit1")
	canvas.Polygon(xp, yp)
	canvas.Gend()
	canvas.Gid("unit2")
	canvas.Polygon(xt, yt)
	canvas.Gend()
	canvas.DefEnd()

	ypositioncounter := 0
	for x := -45; x < width; x += 35 {
		for y := 0; y-60 < height; y += 60 {
			if ypositioncounter%4 == 0 {
				canvas.Use(x, y, "#unit1", "fill:"+c2)
				canvas.Use(x, y, "#unit2", "fill:"+c3)
			}
			if ypositioncounter%4 == 1 {
				canvas.Use(x, y-15, "#unit1", "fill:"+c2)
				canvas.Use(x, y-15, "#unit2", "fill:"+c3)
			}
			if ypositioncounter%4 == 2 {
				canvas.Use(x, y-30, "#unit1", "fill:"+c2)
				canvas.Use(x, y-30, "#unit2", "fill:"+c3)
			}
			if ypositioncounter%4 == 3 {
				canvas.Use(x, y-45, "#unit1", "fill:"+c2)
				canvas.Use(x, y-45, "#unit2", "fill:"+c3)
			}
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
