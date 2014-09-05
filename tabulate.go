package main

import "fmt"

// Basic Structure of TableFormat
type TableFormat struct {
	LineTop []string
	LineBelowHeader []string
	LineBetweenRows []string
	LineBottom []string
	HeaderRow []string
	Padding int
	HeaderHide bool
	FitScreen bool
}

var TableFormats map[string]TableFormat

type Tabulate struct {
	Data map[string]interface{}
	Headers map[string]string
	FormatFloat string
	TableFormat TableFormat
	AlignNum string
	AlignStr string
	EmptyVar string
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

func (t *Tabulate) Render() {

}

func (t *Tabulate) SetHeaders(headers map[string]string) {

}

func (t *Tabulate) setFormatting(format string) {

}

func Create(data interface{}) *Tabulate {
	return &Tabulate{Data:data.(map[string]interface{})}
}

func main() {
	data := map[string]string{"wat":"wat"}
	Create(data)
	fmt.Println("waat")
}
