package art

import (
	"log"
	"os"
)

func Output(fileName, input, banner string) {
	output := []byte(Draw(input, banner))
	error := os.WriteFile(fileName, output, 0644)
	if error != nil {
		log.Fatalf("Error writing file: %s", error)
	}
}
