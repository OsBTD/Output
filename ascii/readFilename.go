package ascii

import (
	"log"
	"os"
	"strings"
)

var (
	filename   string
	Banner     string
	input      string
	inputsplit []string
)

func FileNameManagement() (string, string, string, []string) {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 3 {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

	// Usage instructions if no arguments or more than 2 arguments are provided
	// the default styling will be standard.txt unless the user chooses otherwise
	Banner := "standard.txt"
	filename := "--output=result.txt"
	input := ""
	if len(args) == 1 {
		args = append(args, filename, Banner)
		input = args[0]
		filename = args[1]
		Banner = args[2]

	} else if len(args) == 2 && strings.HasPrefix(args[0], "--output=") {
		args = append(args, Banner)
		filename = args[0]
		input = args[1]
		Banner = args[2]

	} else if len(args) == 2 && (args[1] == "standard.txt" || args[1] == "shadow.txt" || args[1] == "thinkertoy.txt" || args[1] == "standard" || args[1] == "shadow" || args[1] == "thinkertoy") {
		args = append(args, filename)
		input = args[0]
		Banner = args[1]
		filename = args[2]

	} else if len(args) == 3 {
		filename = args[0]
		input = args[1]
		Banner = args[2]
	}
	if !strings.HasSuffix(Banner, ".txt") {
		Banner = Banner + ".txt"
	}
	if Banner != "standard.txt" && Banner != "shadow.txt" && Banner != "thinkertoy.txt" {
		log.Fatal("Usage: this style is unavailabe\nplease chose one of the available styles\n1 : standard.txt\n2 : shadow.txt\n3 : thinkertoy.txt")
	}
	if !strings.HasPrefix(filename, "--output=") {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

	inputsplit = strings.Split(input, "\\n")
	// we check if the input is valid and only contains printable ascii characters
	for _, line := range inputsplit {
		for _, c := range line {
			if c < 32 || c > 126 {
				log.Fatal("Error : input should only contain printable ascii characters")
			}
		}
	}

	filename = strings.TrimPrefix(filename, "--output=")

	return filename, input, Banner, inputsplit
}
