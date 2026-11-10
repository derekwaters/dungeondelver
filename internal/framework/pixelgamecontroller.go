package framework

import (
	"fmt"
	"time"
	"os"
	"image"
	"io/ioutil"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type PixelGameController struct {
	
	// Game Data
	gameData					PixelGameData

	// Game Management Data
	Scenes						map[string]PixelGameScene
	CurrentScene				string

	// Pixel Data
	win							*opengl.Window
	frames						int
	second						<-chan time.Time
	DebugMode					bool

	// textAtlas					*text.Atlas
	// debugAtlas					*text.Atlas


	// Tile Manager?
	// Menu Manager?
	// Options Manager?
	// SaveFile Manager?
}

func (game *PixelGameController) Launch(gameData PixelGameData) {
	game.Log(":: Starting Pixel")
	game.gameData = gameData
	runGame = game
	opengl.Run(run)
	game.Log(":: Finishing Pixel")
}

func (game *PixelGameController) AddScene(name string, scene PixelGameScene) {
	game.Scenes[name] = scene
	if len(game.Scenes) == 1 {
		game.SetScene(name)
	}
}

func (game *PixelGameController) SetScene(name string) {
	_, exist := game.Scenes[name]
	if exist {
		game.Log(":: Scene Changed to ", name)
		game.CurrentScene = name
	}
}

func (game *PixelGameController) LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (game *PixelGameController) LoadFont(path string, size float64) (*text.Atlas, error) {
	face, err := game.LoadTTF(path, size)
	if err != nil {
		return nil, err
	}
	return text.NewAtlas(face, text.ASCII), nil
}

func (game *PixelGameController) LoadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}



func (game *PixelGameController) Log(msg ...interface{}) {
	if (game.DebugMode) {
		fmt.Println(msg)
	}
}


func (game *PixelGameController) Init() {
	game.gameData.Init(game)
	// Init all the scenes
	for _, scene := range game.Scenes {
		scene.Init(game)
	}
}

func (game *PixelGameController) DoRun() {
	game.Scenes = make(map[string]PixelGameScene)

	// Initiate the window
	cfg := opengl.WindowConfig{
		Title: game.gameData.GetTitle(),
		Bounds: game.gameData.GetBounds(),
		VSync: true,
	}

	win, err := opengl.NewWindow(cfg)
	game.win = win
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	win.SetSmooth(true)

	game.Log(":: About to Init Game")
	game.Init()
	game.Log(":: Init Game Finished")

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		
		game.Scenes[game.CurrentScene].Render(game, win, dt)
		
		win.Update()
	}
}



var runGame	*PixelGameController

func run() {
	// Initiate the Scene list
	runGame.DoRun()
}