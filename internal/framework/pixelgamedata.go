package framework

import (
	"fmt"
	"time"
	
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

type PixelGameData struct {
	
	// Setup Data
	Title						string
	Bounds						pixel.Rect
	VSync						bool

	// Game Management Data
	Scenes						map[string]PixelGameScene
	CurrentScene				string

	// Pixel Data
	win							*opengl.Window
	frames						int
	second						<-chan time.Time
	debugMode					bool

	// textAtlas					*text.Atlas
	// debugAtlas					*text.Atlas
}


func (game PixelGameData) Launch() {
	game.Log(":: Starting Pixel")
	opengl.Run(game.run)
	game.Log(":: Finishing Pixel")
}

func (game PixelGameData) AddScene(name string, scene PixelGameScene) {
	game.Scenes[name] = scene
	if len(game.Scenes) == 1 {
		game.SetScene(name)
	}
}

func (game PixelGameData) SetScene(name string) {
	game.Log(":: Changing Scene to " + name)
	_, exist := game.Scenes[name]
	if exist {
		game.CurrentScene = name
	}
}


func (game PixelGameData) Log(msg string) {
	if (game.debugMode) {
		fmt.Println(msg)
	}
}


func (game PixelGameData) Init() {
	game.Log(":: Initialising Game")
}





func (game PixelGameData) run() {
	// Initiate the Scene list
	game.Scenes = make(map[string]PixelGameScene)

	// Initiate the window
	cfg := opengl.WindowConfig{
		Title: game.Title,
		Bounds: game.Bounds,
		VSync: game.VSync,
	}

	win, err := opengl.NewWindow(cfg)
	game.win = win
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	win.SetSmooth(true)

	game.Init()

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		
		game.Scenes[game.CurrentScene].render(win, dt)
		
		win.Update()
	}
}