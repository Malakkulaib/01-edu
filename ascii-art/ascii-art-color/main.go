package main

import (
	"fmt"
	"os"
	"strings"
)

const charHeight = 8

var colors = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"reset":   "\033[0m",
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	var input, bannerName, colorName, colorSubstring string
	colorCode := ""

	if strings.HasPrefix(os.Args[1], "--color=") {
		colorName = strings.TrimPrefix(os.Args[1], "--color=")
		code, ok := colors[colorName]
		if !ok {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			return
		}
		colorCode = code

		if len(os.Args) == 4 {
			colorSubstring = ""
			input = os.Args[2]
			bannerName = os.Args[3]
		} else if len(os.Args) >= 5 {
			colorSubstring = os.Args[2]
			input = os.Args[3]
			bannerName = os.Args[4]
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			return
		}

	} else {
		input = os.Args[1]
		if len(os.Args) >= 3 {
			bannerName = os.Args[2]
		} else {
			fmt.Println("Please enter banner name!")
			return
		}
	}

	bannerFile := bannerName + ".txt"
	banner, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		printASCII(line, banner, colorCode, colorSubstring)
	}
}

func loadBanner(path string) (map[rune][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	banner := make(map[rune][]string)

	ch := 32
	for i := 0; i < len(lines); {
		if lines[i] == "" {
			i++
			continue
		}
		if i+charHeight > len(lines) {
			break
		}
		banner[rune(ch)] = lines[i : i+charHeight]
		i += charHeight
		ch++
	}

	return banner, nil
}

func printASCII(s string, banner map[rune][]string, colorCode string, colorSubstring string) {
	positions := make([]bool, len(s))

	if colorCode != "" && colorSubstring != "" {
		start := 0
		for {
			idx := strings.Index(s[start:], colorSubstring)
			if idx == -1 {
				break
			}
			for i := start + idx; i < start+idx+len(colorSubstring); i++ {
				positions[i] = true
			}
			start += idx + len(colorSubstring)
		}
	} else if colorCode != "" && colorSubstring == "" {
		for i := range positions {
			positions[i] = true
		}
	}

	for line := 0; line < charHeight; line++ {
		for i, ch := range s {
			if art, ok := banner[ch]; ok {
				if positions[i] {
					fmt.Print(colorCode + art[line] + colors["reset"])
				} else {
					fmt.Print(art[line])
				}
			} else {
				fmt.Print("        ")
			}
		}
		fmt.Println()
	}
}
