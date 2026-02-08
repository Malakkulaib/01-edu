package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Error: missing arguments")
		return
	}

	input := args[0]
	output := args[1]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)

	words := strings.Fields(text)

	words = applyCommands(words)
	words = fixVowels(words)
	finalText := fixPunctuation(words)

	os.WriteFile(output, []byte(finalText), 0644)
}

func applyCommands(words []string) []string {
	for i := 0; i < len(words); i++ {

		word := words[i]

		if word == "(hex)" {
			words[i-1] = hexToDec(words[i-1])
			words = remove(words, i)
			i--
		} else if word == "(bin)" {
			words[i-1] = binToDec(words[i-1])
			words = remove(words, i)
			i--
		} else if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = remove(words, i)
			i--
		} else if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = remove(words, i)
			i--
		} else if word == "(cap)" {
			words[i-1] = capitalize(words[i-1])
			words = remove(words, i)
			i--
		} else {
			if strings.HasPrefix(word, "(up,") || strings.HasPrefix(word, "(low,") || strings.HasPrefix(word, "(cap,") {

				if i+1 < len(words) {
					numString := strings.TrimSuffix(words[i+1], ")")
					count, _ := strconv.Atoi(numString)

					for j := 1; j <= count; j++ {
						indexToChange := i - j
						if indexToChange >= 0 {
							if strings.HasPrefix(word, "(up,") {
								words[indexToChange] = strings.ToUpper(words[indexToChange])
							} else if strings.HasPrefix(word, "(low,") {
								words[indexToChange] = strings.ToLower(words[indexToChange])
							} else if strings.HasPrefix(word, "(cap,") {
								words[indexToChange] = capitalize(words[indexToChange])
							}
						}
					}

					words = remove(words, i+1)
					words = remove(words, i)
					i--
				}
			}
		}
	}
	return words
}

func fixVowels(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			nextWord := words[i+1]
			if len(nextWord) > 0 {
				firstLetter := rune(nextWord[0])
				if isVowel(firstLetter) {
					words[i] = words[i] + "n"
				}
			}
		}
	}
	return words
}

func fixPunctuation(words []string) string {
	text := strings.Join(words, " ")

	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", ";")

	text = strings.ReplaceAll(text, ",", ", ")
	text = strings.ReplaceAll(text, ".", ". ")
	text = strings.ReplaceAll(text, "!", "! ")
	text = strings.ReplaceAll(text, "?", "? ")
	text = strings.ReplaceAll(text, ":", ": ")
	text = strings.ReplaceAll(text, ";", "; ")

	for strings.Contains(text, ". .") {
		text = strings.ReplaceAll(text, ". .", "..")
	}
	for strings.Contains(text, "! !") {
		text = strings.ReplaceAll(text, "! !", "!!")
	}
	for strings.Contains(text, "? ?") {
		text = strings.ReplaceAll(text, "? ?", "??")
	}

	newWords := strings.Fields(text)
	result := ""
	open := false

	for i := 0; i < len(newWords); i++ {
		w := newWords[i]
		if w == "'" {
			if !open {
				result = result + " '"
				open = true
			} else {
				result = result + "'"
				open = false
			}
		} else {
			if open && strings.HasSuffix(result, "'") {
				result = result + w
			} else if i == 0 {
				result = result + w
			} else {
				result = result + " " + w
			}
		}
	}
	return result
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func hexToDec(hex string) string {
	val, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(val)
}

func binToDec(bin string) string {
	val, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(val)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	lower := strings.ToLower(s)
	return string(unicode.ToUpper(rune(lower[0]))) + lower[1:]
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouhAEIOUH", r)
}
