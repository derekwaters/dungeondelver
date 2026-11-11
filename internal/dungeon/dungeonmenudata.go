package dungeon

import (
	"github.com/derekwaters/dungeondelver/internal/framework"
	"golang.org/x/image/colornames"
)

func NewDungeonMenuData(controller *framework.PixelGameController) framework.PixelMenuData {
	options := []framework.PixelMenuSelection{
		framework.PixelMenuSelection_NewGame,
		framework.PixelMenuSelection_About,
		framework.PixelMenuSelection_Quit,
	}
	data := framework.PixelMenuData{
		BackgroundColor: colornames.Midnightblue,
		MenuTextColor: colornames.Darkred,
		BackgroundImageTag: "menubg",
		MenuFontTag: "menu",
		TitleFontTag: "menutitle",
		Title: "Delve The Dungeon",
		MenuOptions: options,
		NewGameSceneName: "room",
		AboutSceneName: "debug",
	}
	return data
}
