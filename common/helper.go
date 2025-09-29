package common

import (
	"bufio"
	"os"
)

func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func ReadFileTask4(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	line := ""
	for scanner.Scan() {
		if scanner.Text() == "" {
			lines = append(lines, line)
			line = ""
		} else {
			line += scanner.Text() + " "
		}
	}

	if (line) != "" {
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
