package euphonic

import (
	"context"
	"fmt"
	"io"
	"math"
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
		r.flag.Color5 = color.Random()
		r.flag.Color6 = color.Random()
		fmt.Printf("Set Color1:%s Color2:%s Color3:%s Color4:%s Color4:%s Color4:%s\n", r.flag.Color1, r.flag.Color2, r.flag.Color3, r.flag.Color4, r.flag.Color5, r.flag.Color6)
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
	c1, c2, c3, c4, c5, c6 := r.flag.Color1, r.flag.Color2, r.flag.Color3, r.flag.Color4, r.flag.Color5, r.flag.Color6

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())

	hexSide := float64(80)
	stripeWidth := float64(80 / 5) // Hexagon side divided by number of stripes
	x0 := float64(0)
	y0 := float64(0)
	x1 := float64(0)
	y1 := hexSide
	x2, y2 := rotate(x0, y0, stripeWidth, 0, 30)
	x3, y3 := rotate(x1, y1, x1+stripeWidth, y1, 30)
	x5, y5 := rotate(x0, y0, 0, -hexSide, 60)
	x4, y4 := rotate(x5, y5, x5+stripeWidth, y5, 30)
	// Polygon.
	xp1 := []float64{x0, x1, x3, x2, x4, x5}
	yp1 := []float64{y0, y1, y3, y2, y4, y5}

	canvas.Def(func() {
		canvas.Gid("blockc1", func() {
			canvas.Polygonf64(xp1, yp1, "fill:"+c1)
			canvas.Translatef64(x2, y2, func() {
				canvas.Polygonf64(xp1, yp1, "fill:"+c2)
			})
			canvas.Translatef64((x2 * 2), (y2 * 2), func() {
				canvas.Polygonf64(xp1, yp1, "fill:"+c3)
			})
			canvas.Translatef64((x2 * 3), (y2 * 3), func() {
				canvas.Polygonf64(xp1, yp1, "fill:"+c2)
			})
			canvas.Translatef64((x2 * 4), (y2 * 4), func() {
				canvas.Polygonf64(xp1, yp1, "fill:"+c1)
			})
		})

		canvas.Gid("blockc2", func() {
			canvas.Polygonf64(xp1, yp1, "fill:"+c4)
			canvas.Translatef64(x2, y2, func() {
				canvas.Polygonf64(xp1, yp1, "fill:"+c5)
			})
			canvas.Translatef64((x2 * 2), (y2 * 2), func() {
				canvas.Polygonf64(xp1, yp1, "fill:"+c6)
			})
			canvas.Translatef64((x2 * 3), (y2 * 3), func() {
				canvas.Polygonf64(xp1, yp1, "fill:"+c5)
			})
			canvas.Translatef64((x2 * 4), (y2 * 4), func() {
				canvas.Polygonf64(xp1, yp1, "fill:"+c4)
			})
		})

		canvas.Gid("mirrorblockc1", func() {
			canvas.Rotate(180, func() {
				canvas.Translatef64(-(x2 * 10), -(y2 * 10), func() {
					canvas.Polygonf64(xp1, yp1, "fill:"+c3)
					canvas.Translatef64(x2, y2, func() {
						canvas.Polygonf64(xp1, yp1, "fill:"+c2)
					})
					canvas.Translatef64((x2 * 2), (y2 * 2), func() {
						canvas.Polygonf64(xp1, yp1, "fill:"+c1)
					})
					canvas.Translatef64((x2 * 3), (y2 * 3), func() {
						canvas.Polygonf64(xp1, yp1, "fill:"+c2)
					})
					canvas.Translatef64((x2 * 4), (y2 * 4), func() {
						canvas.Polygonf64(xp1, yp1, "fill:"+c3)
					})
				})
			})
		})

		canvas.Gid("mirrorblockc2", func() {
			canvas.Rotate(180, func() {
				canvas.Translatef64(-(x2 * 10), -(y2 * 10), func() {
					canvas.Polygonf64(xp1, yp1, "fill:"+c6)
					canvas.Translatef64(x2, y2, func() {
						canvas.Polygonf64(xp1, yp1, "fill:"+c5)
					})
					canvas.Translatef64((x2 * 2), (y2 * 2), func() {
						canvas.Polygonf64(xp1, yp1, "fill:"+c4)
					})
					canvas.Translatef64((x2 * 3), (y2 * 3), func() {
						canvas.Polygonf64(xp1, yp1, "fill:"+c5)
					})
					canvas.Translatef64((x2 * 4), (y2 * 4), func() {
						canvas.Polygonf64(xp1, yp1, "fill:"+c6)
					})
				})
			})
		})

		canvas.Gid("fullc1", func() {
			canvas.Use(0, 0, "#blockc1")
			canvas.Use(0, 0, "#mirrorblockc1")
		})

		canvas.Gid("fullc2", func() {
			canvas.Use(0, 0, "#blockc2")
			canvas.Use(0, 0, "#mirrorblockc2")
		})
	})

	// width: 138 (side to side)
	// height: 120 (edge to edge)
	// half width: 69

	ypositioncounter := 0
	for y := -200; y < height+200; y += 120 {
		xpositioncounter := 0
		for x := -200; x < width+200; x += 138 {
			if ypositioncounter%2 == 1 {
				canvas.Use(x, y, "#fullc1")
			} else {
				canvas.Use(x+69, y, "#fullc2")
			}
			xpositioncounter++
		}
		ypositioncounter++
	}

	canvas.End()

	return nil
}

func rotate(originx, originy, x, y, angle float64) (float64, float64) {
	radian := angle * (math.Pi / 180)
	s := math.Sin(radian)
	c := math.Cos(radian)

	x = x - originx
	y = y - originy

	xnew := x*c - y*s
	ynew := x*s + y*c

	x = xnew + originx
	y = ynew + originy

	return x, y
}
