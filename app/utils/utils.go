package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/jszwec/csvutil"
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

func exportCsvProcess(taskCsv []byte) {
	//var process []Process

	csvReader := csv.NewReader(strings.
		NewReader(`"dwm.exe","1632","Console","1","83,944 K","Running","N/A","0:00:49","DWM Notification Window"`))

	userHeader, err := csvutil.Header(Process{}, "csv")
	if err != nil {
		log.Fatal(err)
	}

	dec, err := csvutil.NewDecoder(csvReader, userHeader...)
	if err != nil {
		log.Fatal(err)
	}

	var proc []Process
	for {

		var p Process

		if err := dec.Decode(&p); err != nil {
			break
		}

		proc = append(proc, p)

	}

	for _, p := range proc {
		fmt.Printf("%+v\n", p)
	}

}
