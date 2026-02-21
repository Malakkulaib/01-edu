package main

import (
	"fmt"
	"os"
	"strings"
)

const charHeight = 8

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	var input, bannerName, outputFile string

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.HasPrefix(arg, "--output=") {
			outputFile = strings.TrimPrefix(arg, "--output=")
		} else if input == "" {
			input = arg
		} else if bannerName == "" {
			bannerName = arg
		}
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
	var outputLines []string
	for _, line := range lines {
		if line == "" {
			continue
		}
		for lineNum := 0; lineNum < charHeight; lineNum++ {
			var sb strings.Builder
			for _, ch := range line {
				if art, ok := banner[ch]; ok {
					sb.WriteString(art[lineNum])
				} else {
					sb.WriteString("        ")
				}
			}
			outputLines = append(outputLines, sb.String())
		}
	}

	if outputFile == "" {
		for _, l := range outputLines {
			fmt.Println(l)
		}
	} else {
		err := os.WriteFile(outputFile, []byte(strings.Join(outputLines, "\n")+"\n"), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
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

func printASCII(s string, banner map[rune][]string) {
	for line := 0; line < charHeight; line++ {
		for _, ch := range s {
			if art, ok := banner[ch]; ok {
				fmt.Print(art[line])
			} else {
				fmt.Print("        ")
			}
		}
		fmt.Println()
	}
}
