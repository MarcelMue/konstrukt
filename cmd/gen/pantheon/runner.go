package pantheon

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

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Square 1.
	xp1 := []int{0, 0, 50, 50}
	yp1 := []int{0, 50, 50, 0}

	// Square 2.
	xp2 := []int{5, 5, 45, 45}
	yp2 := []int{5, 45, 45, 5}

	// Square 3.
	xp3 := []int{10, 10, 40, 40}
	yp3 := []int{10, 40, 40, 10}

	// Square 4.
	xp4 := []int{15, 15, 45, 45}
	yp4 := []int{15, 45, 45, 15}

	canvas.Def()
	canvas.Gid("unit")
	canvas.Polygon(xp1, yp1, "fill:"+c2)
	canvas.Polygon(xp2, yp2, "fill:"+c1)
	canvas.Polygon(xp3, yp3, "fill:"+c2)
	canvas.Polygon(xp4, yp4, "fill:"+c1)
	canvas.Gend()
	canvas.DefEnd()

	for y := -50; y < height+200; y += 60 {
		for x := -50; x < width+200; x += 60 {
			canvas.Use(x, y, "#unit")
		}
	}
	canvas.End()

	return nil
}
