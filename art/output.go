package art

import (
	"log"
	"os"
)
/*
file, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("Appending another line of text.\n")
if err != nil {
    log.Fatal(err)
}
*/

func Output(lines []string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, line :=  range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Error writing file: %s", err)
		}
	}	
	// output := []byte(line)
	// error := os.WriteFile(fileName, output, 0644)
	// if error != nil {
	// 	log.Fatalf("Error writing file: %s", error)
	// }
}
