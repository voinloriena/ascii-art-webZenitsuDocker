package asciigo

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

var bannerMap = map[string]string{
	"standard":   "banners/standard.txt",
	"shadow":     "banners/shadow.txt",
	"thinkertoy": "banners/thinkertoy.txt",
}

var fileHashes = map[string]string{
	"banners/standard.txt": "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"banners/shadow.txt":   "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	// "banners/thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
	"banners/thinkertoy.txt": "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52",
}

func HashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Normalize line endings to \n by writing text without \r
		_, err := hash.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			return "", err
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func VerifyFile(filePath string) bool {
	expectedHash, exists := fileHashes[filePath]
	if !exists {
		fmt.Printf("Error: No hash found for %s\n", filePath)
		return false
	}

	computedHash, err := HashFile(filePath)
	if err != nil {
		fmt.Println("Error computing file hash:", err)
		return false
	}

	return computedHash == expectedHash
}

func getLine(filePath string, num int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open banner file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for lineNum := 0; scanner.Scan(); lineNum++ {
		if lineNum == num {
			return scanner.Text(), nil
		}
	}
	return "", nil
}

func IsValidASCII(text string) bool {
	for _, char := range text {
		if char != '\n' && char != '\r' && (char < ' ' || char > '~') {
			return false
		}
	}
	return true
}

func GenerateAsciiArt(input string, banner string) (string, error) {
	bannerFile, exists := bannerMap[banner]
	if !exists {
		return "", fmt.Errorf("bad request: invalid banner '%s' selected", banner)
	}

	if !VerifyFile(bannerFile) {
		return "", fmt.Errorf("internal error: %s file is invalid or tampered", bannerFile)
	}

	if !IsValidASCII(input) {
		return "", fmt.Errorf("not supported symbols: input contains non-ASCII characters (32-126 only)")
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	lines := strings.Split(input, "\n")
	var result strings.Builder
	onlyEmptyLines := true
	emptyLineCount := 0

	for _, line := range lines {
		if line == "" {
			emptyLineCount++
			continue
		}

		for emptyLineCount > 0 {
			result.WriteString("\n")
			emptyLineCount--
		}

		onlyEmptyLines = false

		for i := 0; i < 8; i++ {
			var lineBuilder strings.Builder
			for _, letter := range line {
				lineText, err := getLine(bannerFile, 1+int(letter-' ')*9+i)
				if err != nil {
					return "", err
				}
				lineBuilder.WriteString(lineText)
			}
			result.WriteString(lineBuilder.String())
			result.WriteString("\n")
		}
	}

	if onlyEmptyLines {
		for i := 0; i < len(lines)-1; i++ {
			result.WriteString("\n")
		}
	} else {
		for emptyLineCount > 0 {
			result.WriteString("\n")
			emptyLineCount--
		}
	}

	return result.String(), nil
}
