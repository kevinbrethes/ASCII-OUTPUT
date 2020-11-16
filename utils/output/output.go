package output

import (
	"log"
	"os"
)

func FileName(arg string) string {
	name := arg[9:]
	return name
}

func WriteFile(arg string) string {
	file, err := os.Create(FileName(arg))
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return FileName(arg)
}
