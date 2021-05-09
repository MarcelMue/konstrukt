package nolock

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
	xp1 := []int{0, 10, 50, 30, 20, 30}
	yp1 := []int{10, 0, 40, 60, 50, 40}

	canvas.Def(func() {
		canvas.Gid("unit", func() {
			canvas.Polygon(xp1, yp1)
		})
		canvas.Gid("runit", func() {
			canvas.TranslateRotate(30, 90, 180, func() {
				canvas.Use(0, 0, "#unit")
			})
		})
	})

	ypositioncounter := 0
	for y := -100; y < height+200; y += 50 {
		for x := -120; x < width+200; x += 100 {
			if ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", "fill:"+c2)
				canvas.Use(x, y, "#runit", "fill:"+c2)
			}
			if ypositioncounter%2 == 1 {
				canvas.Use(x+50, y, "#unit", "fill:"+c3)
				canvas.Use(x+50, y, "#runit", "fill:"+c3)
			}
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
