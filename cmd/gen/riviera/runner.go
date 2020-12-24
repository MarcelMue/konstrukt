package riviera

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

	stylefmt := "stroke:%s;stroke-width:%d;fill:none"

	// Polyline 1.
	xp1 := []int{40, 40, 60, 80, 80, 60, 40}
	yp1 := []int{20, 40, 70, 40, 20, -10, 20}

	// Polyline 2.
	xp2 := []int{60, 20, 60}
	yp2 := []int{-50, 30, 110}

	canvas.Def()
	canvas.Gid("unit")
	canvas.Polyline(xp1, yp1, "fill:"+c3)
	canvas.Polyline(xp2, yp2, fmt.Sprintf(stylefmt, c2, 15))
	canvas.Gend()
	canvas.DefEnd()
	ypositioncounter := 0
	for y := -100; y < height+200; y += 80 {
		for x := -120; x < width+200; x += 80 {
			if ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit")
			}
			if ypositioncounter%2 == 1 {
				canvas.Use(x+40, y, "#unit")
			}

		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
