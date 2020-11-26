package blockplay

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

	canvas.Def()
	canvas.Gid("unit")
	canvas.Polygon(xp1, yp1, "fill:"+c2)
	canvas.Polygon(xp2, yp2, "fill:"+c2)
	canvas.Polygon(xp3, yp3, "fill:"+c2)
	canvas.Polygon(xp4, yp4, "fill:"+c2)
	canvas.Polygon(xp5, yp5, "fill:"+c2)
	canvas.Gend()
	canvas.Gid("runit")
	canvas.ScaleXY(1, -1)
	canvas.Use(0, 10, "#unit")
	canvas.Gend()
	canvas.Gend()
	canvas.Gid("r2unit")
	canvas.ScaleXY(-1, 1)
	canvas.Use(10, 0, "#unit")
	canvas.Gend()
	canvas.Gend()
	canvas.Gid("r3unit")
	canvas.ScaleXY(-1, 1)
	canvas.Use(10, 0, "#runit")
	canvas.Gend()
	canvas.Gend()
	canvas.DefEnd()

	ypositioncounter := 0
	for y := -220; y < height+200; y += 220 {
		for x := -140; x < width+200; x += 140 {

			canvas.Use(x, y, "#unit")
			canvas.Use(x, y, "#runit")
			canvas.Use(x, y, "#r2unit")
			canvas.Use(x, y, "#r3unit")
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
