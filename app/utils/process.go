package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/jszwec/csvutil"
)

type Process struct {
	ImageName   string    `csv:"col1"`
	PID         int       `csv:"col2"`
	SessionName string    `csv:"col3"`
	SessionID   string    `csv:"col4"`
	MemoryUsage float64   `csv:"col5"`
	Status      string    `csv:"col6"`
	Username    string    `csv:"col7"`
	CPUTime     time.Time `csv:"col8"`
	WindowTitle string    `csv:"col9"`
}

func (p *Process) UnmarshalCSV(data []byte) error {
	parts := strings.Split(string(data), ",")
	if len(parts) != 2 {
		return csvutil.ErrFieldCount
	}

	sizeWithSuffix := parts[1]
	sizeWithoutK := strings.TrimRight(sizeWithSuffix, "K")
	size, err := strconv.ParseFloat(sizeWithoutK, 64)
	if err != nil {
		return err
	}

	p.MemoryUsage = size * 1024

	return nil
}

func convertCPUTimeToTime(cpuTime string) (time.Duration, error) {
	durationString := strings.Split(cpuTime, ":")
	if len(durationString) != 3 {
		return 0, errors.New("invalid CPU time format")
	}

	hours, err := strconv.Atoi(durationString[0])
	if err != nil {
		return 0, err
	}

	minutes, err := strconv.Atoi(durationString[1])
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.Atoi(durationString[2])
	if err != nil {
		return 0, err
	}

	return time.Duration(hours*60*60+minutes*60+seconds) * time.Second, nil
}

func convertMemoryUsageToInt(memoryUsage string) (int, error) {
	parts := strings.Split(memoryUsage, " ")
	if len(parts) != 2 {
		return 0, errors.New("invalid memory usage format")
	}

	numericValue, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	kilobytes := numericValue * 1024
	return kilobytes, nil
}
