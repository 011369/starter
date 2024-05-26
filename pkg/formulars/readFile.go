package formulars

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadDataFile() ([]float64, error) {
	args := os.Args[1:]
	if len(args) > 1 {
		return nil, fmt.Errorf("error: usage is \"go run main.go data.txt\"")
	}
	fileName := args[0]
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []float64
	for scanner.Scan() {

		line := scanner.Text()
		pattern := `^-?\d+(\.?\d+)?(\s?\d+(\.?\d+)?)*$`
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("error regexp: %w", err)
		}
		matches := re.FindStringSubmatch(line)
		trimmedGroup := strings.ReplaceAll(matches[2], " ", "")
		newLine := strings.Replace(line, matches[2], trimmedGroup, 1)
		if !validCombo(newLine) {
			return nil, fmt.Errorf("invalid combination of digits")
		} else if validCombo(newLine) {
			num, err := strconv.ParseFloat(newLine, 64)
			if err != nil {
				return nil, err
			}
			data = append(data, num)
		}
	}
	return data, nil
}
func validCombo(s string) bool {
	pattern := `^-?\d+(\.?\d+)?$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	bul := re.MatchString(s)
	return bul

}
