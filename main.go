package main

import "github.com/roronoadiogo/pick-game-wallpaper/app/config"

var logger = config.InitializeConfigs()

func main() {

	logger.Info("Start Application")

}
