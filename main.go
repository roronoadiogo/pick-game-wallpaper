package main

import (
	"github.com/roronoadiogo/pick-game-wallpaper/app/config"
	"github.com/roronoadiogo/pick-game-wallpaper/app/utils"
)

var logger = &config.AppLogger{Logger: config.InitializeConfigs()}

func main() {

	logger.Info("Start Application")

	utils.FindProcessGame()

}
