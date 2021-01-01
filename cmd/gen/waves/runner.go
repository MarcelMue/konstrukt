package waves

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	svg "github.com/marcelmue/svgo"
	"github.com/spf13/cobra"

	"github.com/marcelmue/konstrukt/pkg/color"
	"github.com/marcelmue/konstrukt/pkg/project"
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
	canvas.Desc(project.PatternDesc())
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Polygon.
	xp1 := []int{0, 0, 20, 20, 30, 30, 30}
	yp1 := []int{0, 20, 50, 80, 80, 50, 40}

	canvas.Def()
	canvas.Gid("shape")
	canvas.Polygon(xp1, yp1)
	canvas.Gend()
	canvas.Gid("unit")
	canvas.Use(0, 0, "#shape")
	canvas.Gend()
	canvas.Gid("runit")
	canvas.ScaleXY(-1, -1)
	canvas.Use(0, 0, "#shape")
	canvas.Gend()
	canvas.Gend()
	canvas.DefEnd()

	for y := -100; y < height+200; y += 160 {
		xpositioncounter := 0
		for x := -120; x < width+200; x += 50 {
			if xpositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", "fill:"+c2)
				canvas.Use(x, y+80, "#runit", "fill:"+c3)
			}
			if xpositioncounter%2 == 1 {
				canvas.Use(x, y+80, "#unit", "fill:"+c2)
				canvas.Use(x, y, "#runit", "fill:"+c3)
			}

			xpositioncounter++
		}
	}
	canvas.End()

	return nil
}
