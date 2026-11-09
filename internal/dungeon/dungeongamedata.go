package dungeon

import (
	"github.com/derekwaters/dungeondelver/internal/framework"
)

type DungeonGameData struct {
	
	framework.PixelGameData
	
	score			int
}


func (game DungeonGameData) Init() {
	game.Log(":: Initialising Dungeon")
	debugScene := DungeonDebugScene{}
	roomScene := DungeonRoomScene{}
	game.AddScene("debug", debugScene)
	game.AddScene("room", roomScene)
}


