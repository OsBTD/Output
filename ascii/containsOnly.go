package ascii

import "strings"

func ContainsOnly(string) bool {
	_, input, _, _ = FileNameManagement()
	for i := 0; i < len(input); i++ {
		if !strings.ContainsAny(string(input[i]), "\\n") {
			return false
		}
	}
	return true
}
