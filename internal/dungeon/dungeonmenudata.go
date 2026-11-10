package dungeon

import (
	"github.com/derekwaters/dungeondelver/internal/framework"
	"github.com/gopxl/pixel/v2/ext/text"
	"golang.org/x/image/colornames"
	"image/color"
)

type PixelMenuSelection int

const (
	PixelMenuSelection_NewGame = iota
	PixelMenuSelection_Options
	PixelMenuSelection_About
	PixelMenuSelection_Quit
)

var PixelMenuSelectionName = map[PixelMenuSelection]string{
	PixelMenuSelection_NewGame:		"New Game",
	PixelMenuSelection_Options:		"Options",
	PixelMenuSelection_About:		"About",
	PixelMenuSelection_Quit:		"Quit",
}

type DungeonMenuData struct {
	menuFont	*text.Atlas
}

func NewDungeonMenuData(controller *framework.PixelGameController) DungeonMenuData {
	var data = DungeonMenuData {}
	data.menuFont, _ = controller.LoadFont("resources/fonts/Chibi.ttf", 40)
	return data
}

func (data DungeonMenuData) GetBackgroundColor() color.RGBA {
	return colornames.Yellow
}

func (data DungeonMenuData) GetBackgroundImagePath() string {
	return "resources/menubg.png"
}

func (data DungeonMenuData) GetMenuFont() *text.Atlas {
	return data.menuFont
}

func (data DungeonMenuData) GetTitle() string {
	return "Welcome To The Dungeon"
}

func (data DungeonMenuData) GetMenuOptions() ([]framework.PixelMenuSelection) {
	return []framework.PixelMenuSelection{
		PixelMenuSelection_NewGame,
		PixelMenuSelection_About,
		PixelMenuSelection_Quit,
	}
}

func  (data DungeonMenuData) GetAboutSceneName() string {
	return "debug"
}
