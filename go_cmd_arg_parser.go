package go_cmd_arg_parser

import (
	// "fmt"
	"strings"
	"unicode"
)

// This is a '\'
const escapeRune = rune(92)

func ParseFromCmdString(cmd string) []string {
	lastQuote := rune(0)
	consecutiveEscapeCount := 0
	splitWithQuotes := func(c rune) bool {
		switch {
		case c == lastQuote:
			if consecutiveEscapeCount%2 == 0 {
				lastQuote = rune(0)
			}
			consecutiveEscapeCount = 0
			return false
		case c == escapeRune:
			consecutiveEscapeCount += 1
			return false
		case lastQuote != rune(0):
			consecutiveEscapeCount = 0
			return false
		case unicode.In(c, unicode.Quotation_Mark):
			lastQuote = c
			consecutiveEscapeCount = 0
			return false
		default:
			consecutiveEscapeCount = 0
			return unicode.IsSpace(c)
		}
	}
	result := strings.FieldsFunc(cmd, splitWithQuotes)
	return trimOuterQuotesFromArgs(result)
}

// Remove outer quotation marks from string
func trimOuterQuotesFromArgs(args []string) []string {
	var result []string
	for i := range args {
		arg := args[i]
		chars := []rune(arg)
		if unicode.In(chars[0], unicode.Quotation_Mark) && chars[0] == chars[len(chars)-1] {
			arg = arg[1 : len(arg)-1]
		}
		result = append(result, arg)
	}

	return result
}
