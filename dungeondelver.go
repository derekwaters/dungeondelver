package main

import(
	"fmt"
	"github.com/derekwaters/dungeondelver/internal/dungeon"
)

func main() {
	fmt.Println("--- Starting The Game ---")

	fmt.Println(dungeon.GetTileset())

	theGame = DungeonGameData {
		Title: "My Dungeon Game",
		Bounds: pixel.R(0, 0, 1200, 960),
		VSync: true,
		debugMode: true,
	}
	theGame.Launch()

	fmt.Println("---  Ending The Game  ---")
}