package modernhive

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
	stylefmt := "stroke:%s;stroke-width:%d"

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Polyline.
	xpl := []int{20, 0, -20, -20, 0, 20}
	ypl := []int{15, 0, 15, 60, 75, 60}

	canvas.Def(func() {
		canvas.Gid("unit", func() {
			canvas.Polyline(xpl, ypl, "fill:none")
		})
		canvas.Gid("runit", func() {
			canvas.TranslateRotate(0, 100, 180, func() {
				canvas.Use(0, 0, "#unit")
			})
		})
	})

	offset := (width + height) / 2
	ypositioncounter := 0
	for x := 0; x < width+offset; x += 60 {
		xpositioncounter := 0
		for y := -offset; y < height; y += 81 {
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
