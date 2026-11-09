package framework

import (
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)


type PixelGameScene struct {
}

func (scene PixelGameScene) render(win *opengl.Window, dt float64) {
	win.Clear(colornames.Grey)
}
