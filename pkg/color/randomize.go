package color

import (
	"math/rand"
	"time"
)

// Random returns a random color from the color palette.
func Random() string {
	rand.Seed(time.Now().UnixNano())
	return colorPalette[rand.Intn(len(colorPalette))]
}
