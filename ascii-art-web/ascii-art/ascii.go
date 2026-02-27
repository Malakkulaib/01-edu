package ascii_art

import (
	"os"
	"strings"
)

const charHeight = 8

func LoadBanner(path string) (map[rune][]string, error) {
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

func RenderASCII(s string, banner map[rune][]string) string {
	var sb strings.Builder
	for line := 0; line < charHeight; line++ {
		for _, ch := range s {
			if art, ok := banner[ch]; ok {
				sb.WriteString(art[line])
			} else {
				sb.WriteString("        ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
