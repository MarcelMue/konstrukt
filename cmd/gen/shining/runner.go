package shining

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

	xp := []int{50, 70, 70, 50, 30, 30}
	yp := []int{40, 50, 75, 85, 75, 50}
	xl := []int{0, 0, 50, 100, 100}
	yl := []int{100, 40, 10, 40, 100}
	stylefmt := "stroke:%s;stroke-width:%d;fill:%s"
	width, height := r.flag.Width, r.flag.Height

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Def()
	canvas.Gid("unit")
	canvas.Polyline(xl, yl, "fill:none")
	canvas.Polygon(xp, yp)
	canvas.Gend()
	canvas.Gid("runit")
	canvas.TranslateRotate(150, 180, 180)
	canvas.Use(0, 0, "#unit")
	canvas.Gend()
	canvas.Gend()
	canvas.DefEnd()

	canvas.Rect(0, 0, width, height, "fill:"+r.flag.Background)
	canvas.Gstyle(fmt.Sprintf(stylefmt, r.flag.Foreground, 12, r.flag.Accent))
	for y := -33; y < height; y += 130 {
		for x := -50; x < width; x += 100 {
			canvas.Use(x, y, "#unit")
			canvas.Use(x, y, "#runit")
		}
	}
	canvas.Gend()
	canvas.End()

	return nil
}
