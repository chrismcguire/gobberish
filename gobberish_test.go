package gobberish

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode"
)

// getItemInRangeTable

func TestGetItemInSingle16RangeTable(t *testing.T) {
	ranges := [][]int{{5, 15, 2}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned, _ := getItemInRangeTable(4, tables)

	assert.Equal(
		t,
		13,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemInSingle32RangeTable(t *testing.T) {
	ranges := [][]int{{100000, 100010, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned, _ := getItemInRangeTable(7, tables)

	assert.Equal(
		t,
		100007,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemWithFirstItemInRange(t *testing.T) {
	ranges := [][]int{{7, 10, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned, _ := getItemInRangeTable(0, tables)

	assert.Equal(
		t,
		7,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemWithLastItemInRange(t *testing.T) {
	ranges := [][]int{{0, 5, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned, _ := getItemInRangeTable(5, tables)

	assert.Equal(
		t,
		5,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemInMultiRangeRangeTable(t *testing.T) {
	ranges := [][]int{{0, 5, 1}, {100000, 100010, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned, _ := getItemInRangeTable(9, tables)

	assert.Equal(
		t,
		100003,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemWithFirstItemInSecondRange(t *testing.T) {
	ranges := [][]int{{10, 20, 1}, {100100, 100200, 20}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned, _ := getItemInRangeTable(11, tables)

	assert.Equal(
		t,
		100100,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemWithLastItemInSecondRange(t *testing.T) {
	ranges := [][]int{{10, 20, 1}, {100100, 100200, 20}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned, _ := getItemInRangeTable(16, tables)

	assert.Equal(
		t,
		100200,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemWithMultipleRangeTables(t *testing.T) {
	ranges := [][]int{{0, 10, 2}, {100010, 100020, 2}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}
	secondRanges := [][]int{{10, 13, 1}, {100000, 100021, 7}}
	secondTable := createRangeTable(secondRanges)
	tables = append(tables, &secondTable)

	returned, _ := getItemInRangeTable(14, tables)

	assert.Equal(
		t,
		12,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemWithLastItemInMultipleRangeTables(t *testing.T) {
	ranges := [][]int{{0, 10, 2}, {100010, 100020, 2}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}
	secondRanges := [][]int{{10, 13, 1}, {100000, 100021, 7}}
	secondTable := createRangeTable(secondRanges)
	tables = append(tables, &secondTable)

	returned, _ := getItemInRangeTable(19, tables)

	assert.Equal(
		t,
		100021,
		returned,
		"Retreived value from range did not match expected",
	)
}

func TestGetItemWhenOutsideRange(t *testing.T) {
	ranges := [][]int{{0, 10, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned, err := getItemInRangeTable(42, tables)

	assert.Equal(
		t,
		-1,
		returned,
		"Retreive value from range when error did not match expected",
	)
	assert.Equal(
		t,
		errors.New("Value not found in range"),
		err,
		"Error not raised when getting outside of range",
	)
}

// totalInRange

func TestTotalInRangeWithEmptySlice(t *testing.T) {
	var tables []*unicode.RangeTable

	returned := totalInRange(tables)
	assert.Equal(t, 0, returned, "Returned non-zero value")
}

func TestTotalInRangeWithSingleRange(t *testing.T) {
	ranges := [][]int{{5, 7, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned := totalInRange(tables)

	assert.Equal(t, 3, returned, "Incorrectly counted range")
}

func TestTotalInRangeWithSingleRangeTable(t *testing.T) {
	ranges := [][]int{{0, 10, 2}, {100000, 100010, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned := totalInRange(tables)

	assert.Equal(t, 17, returned, "Incorrectly counted range")
}

func TestTotalInRangeWithMultipleRangeTables(t *testing.T) {
	ranges := [][]int{{7, 14, 7}, {1000000, 1000001, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}
	secondRanges := [][]int{{8, 10, 1}, {100000, 100021, 7}}
	secondTable := createRangeTable(secondRanges)
	tables = append(tables, &secondTable)

	returned := totalInRange(tables)

	assert.Equal(t, 11, returned, "Incorrectly counted range.")
}

func TestTotalInRangeWithSingleValueRange(t *testing.T) {
	ranges := [][]int{{5, 5, 1}}
	table := createRangeTable(ranges)
	tables := []*unicode.RangeTable{&table}

	returned := totalInRange(tables)

	assert.Equal(t, 1, returned, "Incorrectly counted range")
}

func createRangeTable(ranges [][]int) unicode.RangeTable {
	var sixteens []unicode.Range16
	var thirtyTwos []unicode.Range32

	for _, numRange := range ranges {
		if numRange[1] <= 65535 {
			sixteen := unicode.Range16{
				Lo:     uint16(numRange[0]),
				Hi:     uint16(numRange[1]),
				Stride: uint16(numRange[2]),
			}
			sixteens = append(sixteens, sixteen)
		} else {
			thirtyTwo := unicode.Range32{
				Lo:     uint32(numRange[0]),
				Hi:     uint32(numRange[1]),
				Stride: uint32(numRange[2]),
			}
			thirtyTwos = append(thirtyTwos, thirtyTwo)
		}
	}

	table := unicode.RangeTable{R16: sixteens, R32: thirtyTwos}

	return table
}
