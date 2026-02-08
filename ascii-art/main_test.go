package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("hello", banner)
}

func TestHELLO(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("HELLO", banner)
}

func TestHeLloHuMaN(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("HeLlo HuMaN", banner)
}

func Test1Hello2There(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("1Hello 2There", banner)
}

func TestHelloThere(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("Hello\nThere", banner)
}

func TestHelloNewlineNewlineThere(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("Hello\n\nThere", banner)
}

func TestHelloAndThereSpecial(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("{Hello & There #}", banner)
}

func TestHelloThereNumbers(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("hello There 1 to 2!", banner)
}

func TestMaD3IrALiSboN(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("MaD3IrA&LiSboN", banner)
}

func TestSpecialChars(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("1a\"#FdwHywR&/()=", banner)
}

func TestBracesPipeTilde(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("{|}~", banner)
}

func TestBackslashCaretUnderscoreA(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("[\\]^_ 'a", banner)
}

func TestRGB(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("RGB", banner)
}

func TestSymbols(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII(":;<=>?@", banner)
}

func TestAllSymbols(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("\\!\" #$%&'()*+,-./", banner)
}

func TestUppercaseAlphabet(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("ABCDEFGHIJKLMNOPQRSTUVWXYZ", banner)
}

func TestLowercaseAlphabet(t *testing.T) {
	banner, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	printASCII("abcdefghijklmnopqrstuvwxyz", banner)
}
