package modernhive

import (
	"context"
	"fmt"
	"io"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
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

	canvas := svg.New(f)
	canvas.Start(width, height)
	c1 := r.flag.Color1
	c2 := r.flag.Color2
	stylefmt := "stroke:%s;stroke-width:%d"

	// Background fill
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Polyline
	xpl := []int{20, 0, -20, -20, 0, 20}
	ypl := []int{15, 0, 15, 60, 75, 60}

	canvas.Def()
	canvas.Gid("unit")
	canvas.Polyline(xpl, ypl, "fill:none")
	canvas.Gend()
	canvas.Gid("runit")
	canvas.TranslateRotate(0, 100, 180)
	canvas.Use(0, 0, "#unit")
	canvas.Gend()
	canvas.Gend()
	canvas.DefEnd()

	ypositioncounter := 0
	for x := 0; x < width+500; x += 60 {
		xpositioncounter := 0
		for y := -500; y < height; y += 81 {
			xfinal := x - (xpositioncounter * 25)
			yfinal := y + (ypositioncounter * 20)

			canvas.Use(xfinal, yfinal, "#unit", fmt.Sprintf(stylefmt, c2, 10))
			canvas.Use(xfinal, yfinal, "#runit", fmt.Sprintf(stylefmt, c2, 10))

			xpositioncounter++
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
