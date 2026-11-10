package framework

import (
	// "github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

type PixelGameScene interface {
	Render(game *PixelGameController, win *opengl.Window, dt float64)
	HandleKey(game *PixelGameController, some_key int, dt float64)
}
