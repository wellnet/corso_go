package main

import (
	"regexp"
	"strings"
	"unicode"
)

func main() {

}

func FindWords(text string) []string {
	re := regexp.MustCompile(`\w+`) // w+ is equal to [0-9A-Za-z_]+
	wordsList := re.FindAllString(text, -1)

	return wordsList
}

func FindWords2(text string) []string {
	re := regexp.MustCompile(`[0-9A-Za-z_]+`) // w+ is equal to [0-9A-Za-z_]+
	wordsList := re.FindAllString(text, -1)

	return wordsList
}

// SplitOnNonLetters splits a string on non-letter runes.
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

// SplitOnNonLettersOrNumbers splits a string on non-letter or numbers runes.
func SplitOnNonLettersOrNumbers(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) && !unicode.IsNumber(char) }
	return strings.FieldsFunc(s, notALetter)
}

// SplitOnWords splits a string on words ()[0-9A-Za-z_]+) runes.
func SplitOnWords(s string) []string {
	notALetter := func(char rune) bool {
		return !unicode.IsLetter(char) && !unicode.IsNumber(char) && !unicode.Is(&unicode.RangeTable{
			R16:         []unicode.Range16{unicode.Range16{Lo: 95, Hi: 95, Stride: 1}},
			R32:         []unicode.Range32{},
			LatinOffset: 1,
		}, char)
	}
	return strings.FieldsFunc(s, notALetter)
}
