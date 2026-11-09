package dungeon

import (
	"github.com/derekwaters/dungeondelver/internal/framework"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)


type DungeonDebugScene struct {
	framework.PixelGameScene
}

func (scene DungeonDebugScene) render(win *opengl.Window, dt float64) {
	win.Clear(colornames.Blue)
}
