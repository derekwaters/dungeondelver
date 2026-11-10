package framework

import (
	"github.com/gopxl/pixel/v2/ext/text"
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


type PixelMenuData interface {
	GetBackgroundColor() color.RGBA
	GetBackgroundImagePath() string
	GetMenuFont() *text.Atlas
	GetTitle() string
	GetMenuOptions() ([]PixelMenuSelection)
	GetAboutSceneName() string
	// GetOptionsDefs()
}
