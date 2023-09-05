package main

import "strings"

// A function to turn snake case into camel case
func snakeToCamel(s string) string {
	var result string
	capitalizeNext := false
	for _, c := range s {
		if c == '_' {
			capitalizeNext = true
		} else if capitalizeNext {
			result += strings.ToUpper(string(c))
			capitalizeNext = false
		} else {
			result += string(c)
		}
	}
	return result
}

// A function to turn first character upper case
func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
