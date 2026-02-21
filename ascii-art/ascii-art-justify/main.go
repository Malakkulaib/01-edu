package main

import (
	"fmt"
	"os"
	"strings"
)

const charHeight = 8
const defaultWidth = 80

func main() {
	input := ""
	align := "left"
	bannerName := ""

	args := os.Args[1:]
	for _, arg := range args {
		if strings.HasPrefix(arg, "--align=") {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) == 2 {
				align = parts[1]
			} else {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
				return
			}
		} else if input == "" {
			input = arg
		} else if bannerName == "" {
			bannerName = arg
		}
	}

	if input == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	if bannerName == "" {
		fmt.Println("Please enter banner name!")
		return
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
		printASCII(line, banner, align, defaultWidth)
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

func printASCII(s string, banner map[rune][]string, align string, width int) {
	words := strings.Fields(s)
	wordArts := make([][]string, len(words))
	wordWidths := make([]int, len(words))
	for i, w := range words {
		wordArts[i] = make([]string, charHeight)
		wWidth := 0
		for _, ch := range w {
			if art, ok := banner[ch]; ok {
				for l := 0; l < charHeight; l++ {
					wordArts[i][l] += art[l]
				}
				wWidth += len(art[0])
			} else {
				for l := 0; l < charHeight; l++ {
					wordArts[i][l] += "        "
				}
				wWidth += 8
			}
		}
		wordWidths[i] = wWidth
	}

	for line := 0; line < charHeight; line++ {
		lineStr := ""
		switch align {
		case "justify":
			if len(words) > 1 {
				totalWordsWidth := 0
				for _, w := range wordWidths {
					totalWordsWidth += w
				}
				spaceCount := len(words) - 1
				spaceWidth := 0
				if spaceCount > 0 {
					spaceWidth = (width - totalWordsWidth) / spaceCount
					if spaceWidth < 0 {
						spaceWidth = 0
					}
				}

				for i, art := range wordArts {
					lineStr += art[line]
					if i < len(wordArts)-1 {
						lineStr += strings.Repeat(" ", spaceWidth)
					}
				}
			} else {
				for _, art := range wordArts {
					lineStr += art[line]
				}
			}
		case "right":
			for _, art := range wordArts {
				lineStr += art[line]
			}
			padding := width - len(lineStr)
			if padding < 0 {
				padding = 0
			}
			lineStr = strings.Repeat(" ", padding) + lineStr
		case "center":
			for _, art := range wordArts {
				lineStr += art[line]
			}
			padding := (width - len(lineStr)) / 2
			if padding < 0 {
				padding = 0
			}
			lineStr = strings.Repeat(" ", padding) + lineStr
		default:
			for _, art := range wordArts {
				lineStr += art[line]
			}
		}
		fmt.Println(lineStr)
	}
}
