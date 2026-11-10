package main

import(
	"fmt"
	"github.com/derekwaters/dungeondelver/internal/dungeon"
	"github.com/derekwaters/dungeondelver/internal/framework"
)

func main() {
	fmt.Println("--- Starting The Game ---")

	var theData dungeon.DungeonGameData
	var theGame framework.PixelGameController
	fmt.Println("--- Launching The Game ---")
	theGame.DebugMode = true
	theGame.Launch(theData)

	fmt.Println("---  Ending The Game  ---")
}