package strf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var numberTagRegex *regexp.Regexp

func init() {
	numberTagRegex = regexp.MustCompile(`{([0-9]+)}`)
}

// Format formats the given string like C# `String.Format`.
func Format(format string, args ...interface{}) string {
	return numberTagRegex.ReplaceAllStringFunc(format, func(m string) string {
		// m is like "{123}"
		index, err := strconv.Atoi(m[1 : len(m)-1])
		if err != nil {
			return m
		}
		return fmt.Sprint(args[index])
	})
}

// FormatCore formats the given string like C# `String.Format`.
func FormatCore(format string, args ...interface{}) string {
	return fmt.Sprintf(strings.Replace(format, "{}", "%v", -1), args...)
}
