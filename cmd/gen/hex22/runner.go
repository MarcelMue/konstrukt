package hex22

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
	xp1 := []int{0, 0, 35, 35}
	yp1 := []int{20, 100, 80, 0}
	// Mirrored Triangle.
	xp2 := []int{0, 0, -35, -35}
	yp2 := []int{20, 100, 80, 0}

	canvas.Def(func() {
		canvas.Gid("block", func() {
			canvas.Polygon(xp1, yp1)
		})
		canvas.Gid("mirrorblock", func() {
			canvas.Polygon(xp2, yp2)
		})

		canvas.Gid("full", func() {
			transx := -18
			transy := -70
			rotate := float64(60)

			// Build up actual structure.
			canvas.Use(0, 0, "#block", "fill:"+c3)
			canvas.Use(0, 0, "#mirrorblock", "fill:"+c2)
			canvas.RotateTranslate(transx, transy, rotate, func() {
				canvas.Use(0, 0, "#block", "fill:"+c4)
				canvas.Use(0, 0, "#mirrorblock", "fill:"+c3)
				canvas.RotateTranslate(transx, transy, rotate, func() {
					canvas.Use(0, 0, "#block", "fill:"+c2)
					canvas.Use(0, 0, "#mirrorblock", "fill:"+c4)
					canvas.RotateTranslate(transx, transy, rotate, func() {
						canvas.Use(0, 0, "#block", "fill:"+c3)
						canvas.Use(0, 0, "#mirrorblock", "fill:"+c2)
						canvas.RotateTranslate(transx, transy, rotate, func() {
							canvas.Use(0, 0, "#block", "fill:"+c4)
							canvas.Use(0, 0, "#mirrorblock", "fill:"+c3)
							canvas.RotateTranslate(transx, transy, rotate, func() {
								canvas.Use(0, 0, "#block", "fill:"+c2)
								canvas.Use(0, 0, "#mirrorblock", "fill:"+c4)
							})
						})
					})
				})
			})
		})

	})

	for y := -200; y < height+80; y += 243 {
		xpositioncounter := 0
		for x := -200; x < width+160; x += 210 {
			if xpositioncounter%2 == 0 {
				canvas.Use(x, y+121, "#full")
			} else {
				canvas.Use(x, y, "#full")
			}

			xpositioncounter++

		}
	}
	canvas.End()

	return nil
}
