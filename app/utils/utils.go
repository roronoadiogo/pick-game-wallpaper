package utils

import (
	"os"
	"os/exec"

	"github.com/roronoadiogo/pick-game-wallpaper/app/config"
)

var logger = &config.AppLogger{Logger: config.InitializeConfigs()}

func FindProcessGame() {
	cmd := exec.Command("TASKLIST", "/V", "/FO", "CSV")
	result, err := cmd.Output()

	if err != nil {
		logger.Error("Error not found any process", err)
	}
	exportCsvProcess(result)
}

func exportCsvProcess(taskCsv []byte) (*os.File, error) {
	filename := "data_test.csv"
	file, err := os.Create(filename)

	if err != nil {
		logger.Error("Error creating file:", err)
	}

	defer file.Close()

	_, err = file.Write(taskCsv)
	if err != nil {
		logger.Error("Error writing to the file:", err)
	}

	return file, nil
}
