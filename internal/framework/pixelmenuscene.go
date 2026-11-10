package framework

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	// "github.com/gopxl/pixel/v2/ext/text"
)


type PixelMenuScene struct {
	data		*PixelMenuData
	bgSprite	*pixel.Sprite
}

func NewPixelMenuScene(data PixelMenuData) (PixelMenuScene) {
	return PixelMenuScene{
		data: &data,
		bgSprite: nil,
	}
}

func (scene PixelMenuScene) Init(game *PixelGameController) {
	bgImagePath := (*scene.data).GetBackgroundImagePath()
	bgImage, _ := game.LoadPicture(bgImagePath)
	if bgImage != nil {
		scene.bgSprite = pixel.NewSprite(bgImage, bgImage.Bounds())
	}
}

func (scene PixelMenuScene) Render(game *PixelGameController, win *opengl.Window, dt float64) {
	win.Clear((*scene.data).GetBackgroundColor())

	if scene.bgSprite != nil {
		scene.bgSprite.Draw(win, pixel.IM)
	}

	// Now draw the menu text
}

func (scene PixelMenuScene) HandleKey(game *PixelGameController, some_key int, dt float64) {
}
