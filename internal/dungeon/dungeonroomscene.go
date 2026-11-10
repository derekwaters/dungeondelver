package dungeon

import (
	"github.com/derekwaters/dungeondelver/internal/framework"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)


type DungeonRoomScene struct {
}

func (scene DungeonRoomScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(colornames.Green)
}

func (scene DungeonRoomScene) HandleKey(game *framework.PixelGameController, some_key int, dt float64) {
}