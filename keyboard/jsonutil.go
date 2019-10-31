package keyboard

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ParseJsonFileToString(path string) string {
	file, openFileErr := os.Open(path)
	if openFileErr != nil {
		log.Panic("Error while opening file: ", openFileErr)
	}

	b, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		log.Panic("Error while reading file: ", readErr)
	}

	return strings.ReplaceAll(string(b), "\n", "")
}
