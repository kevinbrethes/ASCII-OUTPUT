package ascii

import (
	"log"
	"os"
)

func appendFile(fileName, text string) { //write text in a file
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	file.WriteString(text)
}

func returnCharacter(characterNumber int, ascii []string) []string { //return an array of 8 lines for each characters
	characterNumber = characterNumber + 8*characterNumber
	var line []string
	for i := characterNumber; i < characterNumber+8; i++ {
		line = append(line, ascii[i])
	}
	return line
}

func returnCharacterLine(text string, ascii []string) []string { //return an array that contain each lines of the function above
	var lineToReturn []string
	for _, val := range text { //val = character of the word
		value := int(val) - 32
		for _, line := range returnCharacter(value, ascii) {
			lineToReturn = append(lineToReturn, line)
		}
	}
	return lineToReturn
}

func WriteTextInAscii(text string, ascii []string, fileName string) { //put the text in a txt file in ASCII art
	characters := returnCharacterLine(text, ascii)

	lines := ""
	for i := 0; i < 8; i++ {
		lines += characters[i]
		j := 8
		for i+j < len(characters) {
			lines += characters[i+j]
			j += 8
		}
		lines += "\n"
	}

	appendFile(fileName, lines)
}
