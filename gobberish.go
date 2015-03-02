// Package gobberish provides helper functions for generating random strings
// for testing.
package gobberish

import "errors"
import "math/rand"
import "time"
import "unicode"

// Generate a random utf-8 string of a given character (not byte) length.
// The range of the random characters is the entire printable unicode range.
func GenerateString(length int) string {
	var s []rune
	for i := 0; i < length; i++ {
		s = append(s, CreateRandomRune())
	}

	return string(s)
}

// Generate a random utf-8 string of a given character length that exists on the
// given RangeTables. For a list of valid RangeTables, see
// http://golang.org/pkg/unicode/#pkg-variables
func GenerateStringInRange(length int, tables ...*unicode.RangeTable) string {
	var s []rune
	for i := 0; i < length; i++ {
		s = append(s, CreateRandomRuneInRange(tables))
	}

	return string(s)
}

// Generates a random rune in the printable range.
func CreateRandomRune() rune {
	return CreateRandomRuneInRange(unicode.GraphicRanges)
}

// Generates a random rune in the given RangeTable.
func CreateRandomRuneInRange(tables []*unicode.RangeTable) rune {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(totalInRange(tables))
	x, _ := getItemInRangeTable(i, tables)

	return rune(x)
}

// Returns the nth item contained in the array of ranges.
func getItemInRangeTable(n int, tables []*unicode.RangeTable) (int, error) {
	nPointer := n
	var picked int
	found := false

	for _, table := range tables {
		if found == false {
			for _, r16 := range table.R16 {
				countInRange := int((r16.Hi-r16.Lo)/r16.Stride) + 1
				if nPointer <= countInRange-1 {
					picked = int(r16.Lo) + (int(r16.Stride) * nPointer)
					found = true
					break
				} else {
					nPointer -= countInRange
				}
			}

			if found == false {
				for _, r32 := range table.R32 {
					countInRange := int((r32.Hi-r32.Lo)/r32.Stride) + 1
					if nPointer <= countInRange-1 {
						picked = int(r32.Lo) + (int(r32.Stride) * nPointer)
						found = true
						break
					} else {
						nPointer -= countInRange
					}
				}
			}
		}
	}

	if found == true {
		return picked, nil
	} else {
		return -1, errors.New("Value not found in range")
	}
}

// Counts the total number of items contained in the array of ranges.
func totalInRange(tables []*unicode.RangeTable) int {
	total := 0
	for _, table := range tables {
		for _, r16 := range table.R16 {
			total += int((r16.Hi-r16.Lo)/r16.Stride) + 1
		}
		for _, r32 := range table.R32 {
			total += int((r32.Hi-r32.Lo)/r32.Stride) + 1
		}
	}
	return total
}
