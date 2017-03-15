package gotabulate

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var HEADERS = []string{"Header 1", "Header 2", "Header 3", "Header 4", "Header 5"}
var INT_ARRAY = []int{1, 2, 3, 1000, 200}
var INT64_ARRAY = []int64{100, 500, 600, 1000, 10000}
var FLOAT_ARRAY = []float64{10.01, 12.002, -123.5, 20.00005, 1.01}
var STRING_ARRAY = []string{"test string", "test string 2", "test", "row", "bndr"}
var MIXED_ARRAY = []interface{}{"string", 1, 1.005, "another string", -2}
var EMPTY_ARRAY = []string{"4th element empty", "4th element empty", "4th element empty"}
var MIXED_MAP = map[string][]interface{}{"header1": MIXED_ARRAY, "header2": MIXED_ARRAY}

// Test Setters
func TestSetFormat(t *testing.T) {
	tabulate := Create([][]float64{FLOAT_ARRAY, FLOAT_ARRAY, FLOAT_ARRAY})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetFloatFormat('f')
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/table_float"))
}

// Test Align Left, Align Right, Align Center
func TestPads(t *testing.T) {
	tabulate := Create([][]float64{FLOAT_ARRAY, FLOAT_ARRAY, FLOAT_ARRAY})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetAlign("left")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/table_float_right_pad"))
	tabulate.SetAlign("right")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/table_float"))
	tabulate.SetAlign("center")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/table_float_center_pad"))

}

func TestSetHeaders(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY})
	tabulate.SetHeaders(HEADERS)
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/test_headers"))
}

func TestHeadersFirstRow(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, STRING_ARRAY})
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/test_first_row"))
}

func TestEmptyString(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, EMPTY_ARRAY})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/test_empty_element"))
}

func TestMaxColWidth(t *testing.T) {
	// TODO
}

func TestSingleColumn(t *testing.T) {
	tab := Create([][]string{
		{"test"},
	})
	tab.SetMaxCellSize(20)
	tab.SetWrapStrings(true)
	tab.Render("grid")
}

// Test Simple
func TestSimpleFloats(t *testing.T) {
	tabulate := Create([][]float64{FLOAT_ARRAY, FLOAT_ARRAY, FLOAT_ARRAY[:len(FLOAT_ARRAY)-1]})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/simple_floats"))
}

func TestSimpleInts(t *testing.T) {
	tabulate := Create([][]int{INT_ARRAY, INT_ARRAY, INT_ARRAY[:len(INT_ARRAY)-1]})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/simple_ints"))
}

func TestSimpleInts64(t *testing.T) {
	tabulate := Create([][]int64{INT64_ARRAY, INT64_ARRAY, INT64_ARRAY[:len(INT64_ARRAY)-1]})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/simple_ints64"))
}

func TestSimpleString(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, STRING_ARRAY, EMPTY_ARRAY})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/simple_strings"))
}

func TestSimpleMixed(t *testing.T) {
	tabulate := Create([][]interface{}{MIXED_ARRAY, MIXED_ARRAY})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/simple_mixed"))
}

func TestSimpleMapMixed(t *testing.T) {
	tabulate := Create(MIXED_MAP)
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/simple_map_mixed"))
}

// Test Grid

func TestGridFloats(t *testing.T) {
	tabulate := Create([][]float64{FLOAT_ARRAY, FLOAT_ARRAY, FLOAT_ARRAY[:len(FLOAT_ARRAY)-1]})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/grid_floats"))
}

func TestGridInts(t *testing.T) {
	tabulate := Create([][]int{INT_ARRAY, INT_ARRAY, INT_ARRAY[:len(INT_ARRAY)-1]})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/grid_ints"))
}

func TestGridInts64(t *testing.T) {
	tabulate := Create([][]int64{INT64_ARRAY, INT64_ARRAY, INT64_ARRAY[:len(INT64_ARRAY)-1]})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/grid_ints64"))
}

func TestGridString(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, STRING_ARRAY, EMPTY_ARRAY})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/grid_strings"))
}

func TestGridMixed(t *testing.T) {
	tabulate := Create([][]interface{}{MIXED_ARRAY, MIXED_ARRAY})
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/grid_mixed"))
}

func TestGridMapMixed(t *testing.T) {
	tabulate := Create(MIXED_MAP)
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/grid_map_mixed"))
}

func TestPaddedHeader(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, STRING_ARRAY, EMPTY_ARRAY})
	tabulate.SetHeaders([]string{"Header 1", "header 2", "header 3", "header 4"})
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/test_padded_headers"))
}

func TestHideLineBelowHeader(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, STRING_ARRAY, EMPTY_ARRAY})
	tabulate.SetHeaders([]string{"Header 1", "header 2", "header 3", "header 4"})
	tabulate.SetHideLines([]string{"belowheader"})
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/test_hide_lines"))
}
func TestWrapCells(t *testing.T) {
	tabulate := Create([][]string{[]string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, STRING_ARRAY, []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, STRING_ARRAY})
	tabulate.SetHeaders([]string{"Header 1", "header 2", "header 3", "header 4"})
	tabulate.SetMaxCellSize(16)
	tabulate.SetWrapStrings(true)
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/test_string_wrap_grid"))
}
func TestWrapCellsSimple(t *testing.T) {
	tabulate := Create([][]string{[]string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, STRING_ARRAY, []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, STRING_ARRAY})
	tabulate.SetHeaders([]string{"Header 1", "header 2", "header 3", "header 4"})
	tabulate.SetMaxCellSize(16)
	tabulate.SetWrapStrings(true)
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/test_string_wrap_simple"))
}
func TestMultiByteString(t *testing.T) {
	tabulate := Create([][]string{
		{"朝", "おはようございます"},
		{"昼", "こんにちわ"},
		{"夜", "こんばんわ"},
	})
	tabulate.SetHeaders([]string{"時間帯", "挨拶"})
	tabulate.SetMaxCellSize(10)
	tabulate.SetWrapStrings(true)
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/test_multibyte_string"))
}
func readTable(path string) string {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func TestSplitCell(t *testing.T) {
	tab := Create([][]string{
		{"header", "value"},
		{"test1", "This is a really long string, yaaaay it works, Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis"},
		{"test2", "AAAAAAAAAAAAAAAAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBCCCCCCCCCCCCCCCCCCCCCCCCCCEEEEEEEEEEEEEEEEEEEEEDDDDDDDDDDDDDDd"},
	})
	tab.SetMaxCellSize(20)
	tab.SetWrapStrings(true)
	tab.SetWrapDelimiter(' ')
	tab.SetSplitConcat("-")
	assert.Equal(t, tab.Render("grid"), readTable("_tests/smart_wrap"))
}

func TestTitlesGrid(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, STRING_ARRAY, EMPTY_ARRAY})
	tabulate.SetTitle("Title One", "center")
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("grid"), readTable("_tests/grid_strings_titled"))
}

func TestTitlesPlain(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, STRING_ARRAY, EMPTY_ARRAY})
	tabulate.SetTitle("Make Titles Great Again", "left")
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("plain"), readTable("_tests/plain_strings_titled"))
}

func TestTitlesSimple(t *testing.T) {
	tabulate := Create([][]string{STRING_ARRAY, STRING_ARRAY, EMPTY_ARRAY})
	tabulate.SetTitle("Simple Title", "right")
	tabulate.SetHeaders(HEADERS)
	tabulate.SetEmptyString("None")
	assert.Equal(t, tabulate.Render("simple"), readTable("_tests/simple_strings_titled"))
}
