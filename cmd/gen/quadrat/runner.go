package quadrat

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

	// Triangle 1.
	xt1 := []int{0, 25, 0}
	yt1 := []int{50, 50, 25}

	// Triangle 2.
	xt2 := []int{50, 25, 50}
	yt2 := []int{50, 25, 25}

	// Triangle 3.
	xt3 := []int{25, 25, 0}
	yt3 := []int{25, 0, 0}

	canvas.Def()
	canvas.Gid("unit")
	canvas.Polygon(xt1, yt1)
	canvas.Polygon(xt2, yt2)
	canvas.Polygon(xt3, yt3)
	canvas.Gend()
	canvas.DefEnd()

	ypositioncounter := 0
	for x := 0; x < width; x += 50 {
		xpositioncounter := 0
		for y := 0; y < height; y += 50 {
			if (ypositioncounter%2 == 0 && xpositioncounter%2 == 0) ||
				ypositioncounter%2 != 0 && xpositioncounter%2 != 0 {
				canvas.Rect(x, y, 50, 50, "fill: "+c2)
				canvas.Use(x, y, "#unit", "fill:"+c1)
			}
			if ypositioncounter%2 != 0 && xpositioncounter%2 == 0 ||
				ypositioncounter%2 == 0 && xpositioncounter%2 != 0 {
				canvas.Rect(x, y, 50, 50, "fill: "+c1)
				canvas.Use(x, y, "#unit", "fill:"+c2)
			}
			xpositioncounter++
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
