package utils

import (
	"fmt"
	"os/exec"
)

func FindProcessGame() {

	cmd := exec.Command("TASKLIST", "/V /NH /FI")
	result, err := cmd.Output()

	if err != nil {
		fmt.Println("Error not found any process")
	}

	fmt.Println(result)

}
