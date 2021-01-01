package whitegold2

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
	stylefmt := "stroke:%s;stroke-width:%d;fill:none"

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Polygon.
	xp := []int{40, 0, 40, 80}
	yp := []int{0, 40, 80, 40}

	// Polyline.
	xl := []int{70, 40, 70}
	yl := []int{10, 40, 70}

	canvas.Def()
	canvas.Gid("unit")
	canvas.Polygon(xp, yp, "fill:"+c2)
	canvas.Polyline(xl, yl, fmt.Sprintf(stylefmt, c1, 10))
	canvas.Gend()
	canvas.Gid("runit")
	canvas.ScaleXY(-1, 1)
	canvas.Use(0, 0, "#unit")
	canvas.Gend()
	canvas.Gend()
	canvas.DefEnd()

	ypositioncounter := 0
	for y := -100; y < height+200; y += 50 {
		for x := -200; x < width+200; x += 100 {
			if ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit")
			}
			if ypositioncounter%2 == 1 {
				canvas.Use(x+30, y, "#runit")
			}
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
