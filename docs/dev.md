# Notes for quick development of patterns

- `svgplay` allows easy setup of a quick iteration framework: https://github.com/marcelmue/svgo/blob/master/svgplay/svgplay.go
- Using `shift+enter` to redraw the svgs in `svgplay` prevents lag
- Transfer content after testing into `konstrukt`

Sample which can be easily transformed to fit into `konstrukt`:
```go
package main

import (
  "os"
  "github.com/marcelmue/svgo"
)

func main() {
  canvas := svg.New(os.Stdout)
  width := 500
  height := 500
  c1, c2 := "#ffffff", "#000000"

  canvas.Start(width, height)
  canvas.Rect(0, 0, width, height, "fill:"+c1)

  // Triangle.
  xp1 := []int{0, 0, 50}
  yp1 := []int{50, 100, 100}

  canvas.Def(func() {
    canvas.Gid("unit", func() {
            canvas.Polygon(xp1, yp1,"fill:"+c2)
    })
  })

  for y := -100; y < height+200; y += 155 {
    for x := -120; x < width+200; x += 120 {
      canvas.Use(x, y, "#unit")
    }
  }
  canvas.End()
}
```
