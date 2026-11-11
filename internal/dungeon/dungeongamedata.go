package dungeon

import (
	"github.com/gopxl/pixel/v2"
	"github.com/derekwaters/dungeondelver/internal/framework"
)

type DungeonGameData struct {
	score			int
}

func (game DungeonGameData) Init(controller *framework.PixelGameController) {
	controller.Log("--- Initialising Dungeon ---")
	// Load Resources
	controller.Resources.AddFont("menu", "resources/fonts/ancient-modern-tales.ttf", 40)
	controller.Resources.AddFont("menutitle", "resources/fonts/ancient-modern-tales.ttf", 60)
	// Create Scenes
	menuData := NewDungeonMenuData(controller)
	menuScene := framework.NewPixelMenuScene(menuData)
	var debugScene DungeonDebugScene
	var roomScene DungeonRoomScene
	controller.AddScene("menu", &menuScene)
	controller.AddScene("debug", &debugScene)
	controller.AddScene("room", &roomScene)

}

func (game DungeonGameData) GetTitle() string {
	return "Welcome To The Dungeon!"
}

func (game DungeonGameData) GetBounds() pixel.Rect {
	return pixel.R(0, 0, 1200, 960)
}

