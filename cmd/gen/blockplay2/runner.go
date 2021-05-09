package blockplay2

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

	// Triangle 1.
	xp1 := []int{0, 0, 50}
	yp1 := []int{50, 100, 100}
	// Rectangle 1.
	xp2 := []int{0, 0, 25, 25}
	yp2 := []int{45, 25, 50, 70}
	// Rectangle 2.
	xp3 := []int{0, 0, 25, 25}
	yp3 := []int{20, 0, 25, 45}
	// Triangle 2.
	xp4 := []int{29, 29, 50}
	yp4 := []int{52, 72, 72}
	// Triangle 3.
	xp5 := []int{29, 29, 50}
	yp5 := []int{27, 47, 47}

	canvas.Def(func() {
		canvas.Gid("unit", func() {
			canvas.Polygon(xp1, yp1)
			canvas.Polygon(xp2, yp2)
			canvas.Polygon(xp3, yp3)
			canvas.Polygon(xp4, yp4)
			canvas.Polygon(xp5, yp5)
		})
		canvas.Gid("runit", func() {
			canvas.ScaleXY(-1, -1, func() {
				canvas.Use(0, -45, "#unit")
			})
		})
	})

	ypositioncounter := 0
	for y := -100; y < height+200; y += 155 {
		xpositioncounter := 0
		for x := -120; x < width+200; x += 120 {
			if xpositioncounter%2 == 0 && ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", "fill:"+c3)
				canvas.Use(x, y, "#runit", "fill:"+c2)
			}
			if xpositioncounter%2 == 1 && ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", "fill:"+c2)
				canvas.Use(x, y, "#runit", "fill:"+c3)
			}
			if xpositioncounter%2 == 0 && ypositioncounter%2 == 1 {
				canvas.Use(x, y, "#unit", "fill:"+c2)
				canvas.Use(x, y, "#runit", "fill:"+c3)
			}
			if xpositioncounter%2 == 1 && ypositioncounter%2 == 1 {
				canvas.Use(x, y, "#unit", "fill:"+c3)
				canvas.Use(x, y, "#runit", "fill:"+c2)
			}
			xpositioncounter++
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
