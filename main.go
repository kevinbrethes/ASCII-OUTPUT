// In this project we want to convert the passed arguments into an ascii art
// We'll use the file "standard.txt" which contain the ASCII table
// Each characters is composed of 8 lines

package main

import (
	"bufio"
	"log"
	"os"

	ascii "../ASCII-OUTPUT/utils/ascii_convert"
	backline "../ASCII-OUTPUT/utils/backlineSupport"
	errors "../ASCII-OUTPUT/utils/errors"
	font "../ASCII-OUTPUT/utils/font"
	output "../ASCII-OUTPUT/utils/output"
)

func scanLines(path string) ([]string, error) { //put each lines of the txt file in an array

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func main() {
	arguments := os.Args[1:]

	if errors.HandlingError(arguments) { //stop the program if an error occured
		return
	}

	text := arguments[0]
	font := font.Choice(arguments)
	filename := arguments[2][9:]

	output.WriteFile(arguments[2])

	lines, err := scanLines(font) //put each lines of the txt file in the array "lines"
	if err != nil {
		log.Fatal(err)
		return
	}

	textArray := backline.BacklineSupport(text)

	for _, words := range textArray {
		ascii.WriteTextInAscii(words, lines, filename)
	}
}
