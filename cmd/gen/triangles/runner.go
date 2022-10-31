package triangles

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
		r.flag.Color4 = color.Random()
		fmt.Printf("Set Color1:%s Color2:%s Color3:%s Color4:%s\n", r.flag.Color1, r.flag.Color2, r.flag.Color3, r.flag.Color4)
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
	c1, c2, c3, c4 := r.flag.Color1, r.flag.Color2, r.flag.Color3, r.flag.Color4

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Triangle.
	xp1 := []int{0, 0, 50}
	yp1 := []int{50, 100, 100}
	// Triangle.
	xp2 := []int{0, 0, -50}
	yp2 := []int{-50, -100, -100}
	// Triangle.
	xp3 := []int{0, 0, 100, 100}
	yp3 := []int{0, 100, 100, 0}

	canvas.Def(func() {
		canvas.Gid("unit", func() {
			canvas.Polygon(xp1, yp1)
			canvas.Polygon(xp2, yp2)
		})
		canvas.Gid("background", func() {
			canvas.Polygon(xp3, yp3)
		})
	})

	ypositioncounter := 0
	for y := -100; y < height+200; y += 100 {
		xpositioncounter := 0
		for x := -100; x < width+200; x += 100 {
			if ypositioncounter%2 == 0 && xpositioncounter%2 == 1 {
				canvas.Use(x, y, "#background", "fill:"+c4)
			}
			if ypositioncounter%2 == 1 && xpositioncounter%2 == 0 {
				canvas.Use(x, y, "#background", "fill:"+c4)
			}
			xpositioncounter++
		}
		ypositioncounter++
	}
	ypositioncounter = 0
	for y := -100; y < height+200; y += 100 {
		for x := -100; x < width+200; x += 100 {
			if ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", "fill:"+c2)
			}
			if ypositioncounter%2 == 1 {
				canvas.Use(x, y, "#unit", "fill:"+c3)
			}
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
