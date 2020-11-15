package validate

import (
	"fmt"
	"regexp"
)

// Color validates if a given string is a valid color.
func Color(c string) bool {
	match, err := regexp.MatchString(`(?:#|0x)(?:[a-f0-9]{3}|[a-f0-9]{6})\b|(?:rgb|hsl)a?\([^\)]*\)`, c)
	if err != nil {
		return false
	}
	return match
}

// Size validates if a given int is a valid size.
func Size(s int) bool {
	if s >= 0 {
		return true
	}
	return false
}

// PrintFailure outputs the message of a failed validation to the command line.
func PrintFailure(flagName string) {
	fmt.Println("\033[31m!!!ERROR!!!")
	fmt.Printf("The value for flag %s is not formatted correctly.\n", flagName)
	fmt.Println("!!!ERROR!!!\033[0m")
}
