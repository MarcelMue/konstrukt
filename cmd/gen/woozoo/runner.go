package woozoo

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

	stylefmt := "stroke:%s;stroke-width:%d;fill:none"

	// Polyline 1.
	xp1 := []int{0, 0, 103}
	yp1 := []int{0, 100, 100}
	// Polyline 2.
	xp2 := []int{10, 10, 113}
	yp2 := []int{-10, 90, 90}
	// Polyline 3.
	xp3 := []int{20, 20, 123}
	yp3 := []int{-20, 80, 80}
	// Polyline 4.
	xp4 := []int{30, 30, 133}
	yp4 := []int{-30, 70, 70}
	// Polyline 5.
	xp5 := []int{40, 40, 143}
	yp5 := []int{-40, 60, 60}

	canvas.Def(func() {
		canvas.Gid("unit", func() {
			canvas.Polyline(xp1, yp1)
			canvas.Polyline(xp2, yp2)
			canvas.Polyline(xp3, yp3)
			canvas.Polyline(xp4, yp4)
			canvas.Polyline(xp5, yp5)
		})

		canvas.Gid("runit", func() {
			canvas.ScaleXY(-1, -1, func() {
				canvas.Use(0, 0, "#unit")
			})
		})
	})

	ypositioncounter := 0
	for y := -100; y < height+200; y += 100 {
		xpositioncounter := 0
		for x := -120; x < width+200; x += 100 {
			if xpositioncounter%2 == 0 && ypositioncounter%2 == 1 {
				canvas.Use(x+90, y+10, "#runit", fmt.Sprintf(stylefmt, c3, 6))
			}
			xpositioncounter++
		}
		ypositioncounter++
	}

	ypositioncounter = 0
	for y := -100; y < height+200; y += 100 {
		xpositioncounter := 0
		for x := -120; x < width+200; x += 100 {
			if xpositioncounter%2 == 0 && ypositioncounter%2 == 0 {
				canvas.Use(x, y, "#unit", fmt.Sprintf(stylefmt, c2, 6))
			}
			if xpositioncounter%2 == 1 && ypositioncounter%2 == 0 {
				canvas.Use(x+90, y+10, "#runit", fmt.Sprintf(stylefmt, c3, 6))
			}
			if xpositioncounter%2 == 1 && ypositioncounter%2 == 1 {
				canvas.Use(x, y, "#unit", fmt.Sprintf(stylefmt, c2, 6))
			}
			xpositioncounter++
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
