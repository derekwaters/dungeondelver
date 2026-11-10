package dungeon

import (
	"github.com/derekwaters/dungeondelver/internal/framework"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)


type DungeonDebugScene struct {
}

func (scene DungeonDebugScene) Init(game *framework.PixelGameController) {
	
}

func (scene DungeonDebugScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(colornames.Blue)
}

func (scene DungeonDebugScene) HandleKey(game *framework.PixelGameController, some_key int, dt float64) {
}