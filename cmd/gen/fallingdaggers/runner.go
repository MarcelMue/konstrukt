package fallingdaggers

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
		fmt.Printf("Set Color1:%s Color2:%s\n", r.flag.Color1, r.flag.Color2)
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
	c1, c2 := r.flag.Color1, r.flag.Color2

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Polygon.
	xp := []int{50, 25, 75}
	yp := []int{0, 25, 75}

	canvas.Def(func() {
		canvas.Gid("unit", func() {
			canvas.Polygon(xp, yp)
		})
		canvas.Gid("runit", func() {
			canvas.ScaleXY(-1, 1, func() {
				canvas.Use(0, 0, "#unit")
			})
		})
	})

	ypositioncounter := 0
	for y := -56; y < height+200; y += 56 {
		for x := -200; x < width+200; x += 200 {
			if ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", "fill:"+c2)
				canvas.Use(x, y, "#runit", "fill:"+c2)
			}
			if ypositioncounter%2 == 1 {
				canvas.Use(x+100, y, "#unit", "fill:"+c2)
				canvas.Use(x+100, y, "#runit", "fill:"+c2)
			}
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
