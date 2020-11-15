package fiftyfive

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

	c1, c2, c3 := r.flag.Color1, r.flag.Color2, r.flag.Color3
	width, height := r.flag.Width, r.flag.Height

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())
	canvas.Rect(0, 0, width, height, "fill:"+c1)

	// Polygon
	xp := []int{0, 80, 100, 80, 80}
	yp := []int{80, 100, 80, 0, 80}
	// Triangle
	xt := []int{0, 80, 80}
	yt := []int{80, 80, 0}

	canvas.Def()
	canvas.Gid("unit1")
	canvas.Polygon(xp, yp)
	canvas.Gend()
	canvas.Gid("unit2")
	canvas.Polygon(xt, yt)
	canvas.Gend()
	canvas.DefEnd()

	for x := 0; x < width; x += 100 {
		for y := 0; y < height; y += 100 {
			canvas.Use(x, y, "#unit1", "fill:"+c2)
			canvas.Use(x, y, "#unit2", "fill:"+c3)
		}
	}
	canvas.End()

	return nil
}
