package main

import (
	"fmt"
	"os"
	"strings"
)

const charHeight = 8

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . \"File Name .txt\" [banner]")
		return
	}

	file := os.Args[1]

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	input := string(data)

	bannerName := "standard"
	if len(os.Args) > 2 {
		bannerName = os.Args[2]
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
		printASCII(line, banner)
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
