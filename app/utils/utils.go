package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func FindProcessGame() {

	cmd := exec.Command("TASKLIST", "/V", "/FO", "CSV")
	result, err := cmd.Output()

	if err != nil {
		fmt.Println("Error not found any process", err)
	}

	exportCsvProcess(result)

	fmt.Println(result)

}

func exportCsvProcess(taskCsv []byte) (*os.File, error) {

	filename := "data_test.csv"
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	defer file.Close()

	_, err = file.Write(taskCsv)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
	}

	return file, nil

}
