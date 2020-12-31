package interruptions

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

	// Rectangle size offsets in each column.
	ypos := []int{0, -80, -75, -5, -145, -100, -50, -10, -130, -120, -20}
	// Repeating color sequence.
	ycol := []string{c1, c1, c2, c3, c3, c2, c2, c1, c3, c3, c2, c1, c2}

	ypositioncounter := 0
	for x := 5; x < width; x += 15 {
		xpositioncounter := 0
		for y := ypos[ypositioncounter%len(ypos)]; y < height; y += 160 {
			canvas.Rect(x, y, 5, 150, "fill: "+ycol[(xpositioncounter+ypositioncounter)%len(ycol)])
			xpositioncounter++
		}
		ypositioncounter++
	}
	canvas.End()

	return nil
}
