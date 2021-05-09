package janein

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

	canvas.Def(func() {
		canvas.Gid("unit", func() {
			canvas.Rect(0, 0, 60, 20)
		})
	})

	for x := -60; x < width+80; x += 80 {
		xpositioncounter := 0
		for y := -40; y < height+20; y += 20 {
			if xpositioncounter%4 == 0 {
				canvas.Rect(x, y, 80, 20, "fill: "+c2)
				canvas.Use(x, y, "#unit", "fill:"+c1)
			}
			if xpositioncounter%4 == 1 {
				canvas.Rect(x+20, y, 80, 20, "fill: "+c1)
				canvas.Use(x+20, y, "#unit", "fill:"+c2)
			}
			if xpositioncounter%4 == 2 {
				canvas.Rect(x+40, y, 80, 20, "fill: "+c2)
				canvas.Use(x+40, y, "#unit", "fill:"+c1)
			}
			if xpositioncounter%4 == 3 {
				canvas.Rect(x+60, y, 80, 20, "fill: "+c1)
				canvas.Use(x+60, y, "#unit", "fill:"+c2)
			}
			xpositioncounter++
		}
	}
	canvas.End()

	return nil
}
