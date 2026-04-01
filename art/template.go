package art

import (
	"fmt"
	"os"
	"strings"
)

func GenerateTemplate(banner string) []string {
	content, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Printf("Error: could not read banner file: %v\n", err)
		return nil
	}
	return strings.Split(string(content), "\n")
}
