package strf

import (
	"fmt"
	"regexp"
	"strconv"
)

var numberTagRegex *regexp.Regexp

func init() {
	numberTagRegex = regexp.MustCompile(`{([0-9]+)}`)
}

// Format formats the given string like C# `String.Format`.
func Format(format string, args ...interface{}) string {
	if len(args) == 0 {
		return format
	}
	return numberTagRegex.ReplaceAllStringFunc(format, func(m string) string {
		// m is like "{123}"
		index, err := strconv.Atoi(m[1 : len(m)-1])
		if err != nil {
			return m
		}
		return fmt.Sprint(args[index])
	})
}

func main() {
	fmt.Print(Format("sf💇‍♂️👬{1}🧖🏿‍♂️{0}-{1}", "zero", 1))
}
