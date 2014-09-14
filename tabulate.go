package main

import "fmt"
import "reflect"

// Basic Structure of TableFormat
type TableFormat struct {
	LineTop         []string
	LineBelowHeader []string
	LineBetweenRows []string
	LineBottom      []string
	HeaderRow       []string
	Padding         int
	HeaderHide      bool
	FitScreen       bool
}

var TableFormats map[string]TableFormat

type Tabulate struct {
	Data        []*TabulateRow
	Headers     []string
	FormatFloat string
	TableFormat TableFormat
	AlignNum    string
	AlignStr    string
	EmptyVar    string
}

type TabulateRow struct {
	Elements []string
}

func (t *Tabulate) normalize() {

}

func (t *Tabulate) drawLine() {

}

func (t *Tabulate) padLine() {

}

func (t *Tabulate) format() {

}

func (t *Tabulate) align() {

}

func (t *Tabulate) Render() string {
	return ""
}

func (t *Tabulate) SetHeaders(headers map[string]string) *Tabulate {
	return t
}

func (t *Tabulate) SetFormatting(format string) *Tabulate {
	return t
}

func Create(data interface{}) *Tabulate {
	t := &Tabulate{}

	switch v := data.(type) {
	case [][]interface{}, [][]string:
		fmt.Println("2D array interface{}")
	case [][]int32, [][]int64, [][]int:
		fmt.Println("Int 64 2d Array")
	case []string:
		fmt.Println("String array")
	case []int32, []int64, []int:
		fmt.Println("Int 32 or int 64")
	case []interface{}:
		t.Data = normalize(data.([]interface{}))
	default:
		fmt.Println(v)
	}

	return t
}

func main() {
	data := []interface{}{"test", "test", "test2"}
	Create([]int{1, 2, 3, 4, 5})
	Create([]string{"1", "2"})
	Create([][]interface{}{{1, 2, 3}, {1, 2, 3}})
	Create(data)
	fmt.Println("waat")
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
