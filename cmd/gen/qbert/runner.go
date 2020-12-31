package qbert

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

	patternWidth := 72
	patternHeight := 54

	patternWidthHalf := patternWidth / 2

	patternHeightThird := patternHeight / 3
	patternHeightFrac := int(float64(patternHeight) / 1.5)

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Desc(project.PatternDesc())

	// Left Rhombus
	xpLeft := []int{0, 0, patternWidthHalf, patternWidthHalf}
	ypLeft := []int{patternHeightThird, patternHeight, patternHeightFrac, 0}
	canvas.Def()
	canvas.Gid("left")
	canvas.Polygon(xpLeft, ypLeft)
	canvas.Gend()
	canvas.DefEnd()

	// Right Rhombus
	xpRight := []int{0, 0, patternWidthHalf, patternWidthHalf}
	ypRight := []int{0, patternHeightFrac, patternHeight, patternHeightThird}
	canvas.Def()
	canvas.Gid("right")
	canvas.Polygon(xpRight, ypRight)
	canvas.Gend()
	canvas.DefEnd()

	canvas.Rect(0, 0, width, height, "fill:"+c1)

	for y := 0; y < height; y += patternHeight * 2 {
		for x := -patternWidthHalf; x < width; x += patternWidth {
			canvas.Use(x, y, "#left", "fill:"+c2)
			canvas.Use(x+patternWidthHalf, y, "#right", "fill:"+c3)
			canvas.Use(x+patternWidthHalf, y+patternHeight, "#left", "fill:"+c2)
			canvas.Use(x+patternWidth, y+patternHeight, "#right", "fill:"+c3)
		}
	}
	canvas.End()

	return nil
}
