package environment

import (
	"fmt"
	"log"
	"os"
)

// ReadAndCheckEnvVar reads environment variable varName, killing the program if it's not set.
// Returns environment variable value.
func ReadAndCheckEnvVar(varName string) (varVal string) {
	if varVal = ReadEnvVar(varName); varVal == "" {
		err := fmt.Errorf("%s environment variable not set properly", varName)
		log.Fatal(err)
	}

	return
}

// ReadEnvVar reads environment variable varName, returning it's value if set, empty string otherwise
func ReadEnvVar(varName string) string {
	return os.Getenv(fmt.Sprint(varName))
}
