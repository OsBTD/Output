package ascii

import (
	"io"
	"log"
	"os"
)

func PrintArt(string, []string, map[rune]([]string), string) {
	filename, input, Banner, _ = FileNameManagement()
	var result string
	// we start by checking if the user input only contains literal newlines
	// if so we print newlines accordingly
	for idx, line := range inputsplit {
		// also if there's empty strings resulting from the spliting we print a newline
		if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
			result += "\n"
			continue
		} else if len(line) == 0 && !Checknewline(inputsplit) {
			result += "\n"
		} else if len(line) != 0 && !Checknewline(inputsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(line); j++ {
					inputrune := rune(line[j])
					result += Replace[inputrune][i]
				}
				// after each line is printed we print a newline
				result += "\n"
			}
		}
	}
	resfile, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file : ", err)
	}
	defer resfile.Close()

	_, err = io.WriteString(resfile, result)
	if err != nil {
		log.Fatal("Error writing to file : ", err)
	}
}
