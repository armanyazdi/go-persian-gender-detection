package persiangenderdetection

import (
	"bufio"
	"os"
	"strings"
	"sync"
	"unicode"
)

var (
	dataset     = make(map[string]string)
	datasetOnce sync.Once
)

// readCSV reads the CSV file and populates the dataset map
// This function is called only once using sync.Once
func readCSV() {
	file, err := os.Open("data/names.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		record := strings.Split(line, ",")
		if len(record) == 2 {
			dataset[record[0]] = record[1]
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// isUnwantedRune checks if a rune is a punctuation, symbol, nonspacing mark, or digit
func isUnwantedRune(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.Is(unicode.Mn, r) || unicode.IsDigit(r)
}

// clearName cleans the input name by removing unwanted runes and replacing specific characters
func clearName(name string) string {
	var builder strings.Builder
	builder.Grow(len(name))

	for _, r := range name {
		if !isUnwantedRune(r) {
			builder.WriteRune(r)
		}
	}

	replacer := strings.NewReplacer(
		"\u200c", " ",
		"آ", "ا",
		"ي", "ی",
		"ك", "ک",
		"ـ", "",
	)

	return strings.TrimSpace(replacer.Replace(builder.String()))
}

// GetGender determines the gender of the given name by consulting the dataset
// It first clears the name of unwanted characters, then checks for prefixes, and finally looks up the gender
func GetGender(name string) string {
	datasetOnce.Do(readCSV)

	clearedName := clearName(name)
	fullName := strings.Fields(clearedName)
	prefixes := map[string]struct{}{
		"سید":   {},
		"سیده":  {},
		"استاد": {},
		"دکتر":  {},
		"مهندس": {},
		"سرکار": {},
	}

	// Remove prefixes from the name
	for len(fullName) > 0 {
		if _, exists := prefixes[fullName[0]]; exists {
			fullName = fullName[1:]
		} else {
			break
		}
	}

	// Check for gender by progressively reducing the full name
	for len(fullName) > 0 {
		firstName := strings.Join(fullName, " ")
		if gender, found := dataset[firstName]; found {
			switch gender {
			case "M":
				return "MALE"
			case "F":
				return "FEMALE"
			}
		}
		fullName = fullName[:len(fullName)-1]
	}

	return "UNKNOWN" // Unknown gender
}
