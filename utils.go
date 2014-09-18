package main

import "strconv"
import "fmt"

func createFromString(data [][]string) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))

	for index, el := range data {
		rows[index] = &TabulateRow{Elements: el}
	}
	return rows
}

func createFromMixed(data [][]interface{}) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for index_1, element := range data {
		normalized := make([]string, len(element))
		for index, el := range element {
			switch el.(type) {
			case int32:
				quoted := strconv.QuoteRuneToASCII(el.(int32))
				normalized[index] = quoted[1 : len(quoted)-1]
			case int:
				normalized[index] = strconv.Itoa(el.(int))
			case int64:
				normalized[index] = strconv.FormatInt(el.(int64), 10)
			case bool:
				normalized[index] = strconv.FormatBool(el.(bool))
			case float64:
				normalized[index] = strconv.FormatFloat(el.(float64), 'f', -1, 64)
			case uint64:
				normalized[index] = strconv.FormatUint(el.(uint64), 10)
			default:
				normalized[index] = fmt.Sprintf("%#v", el)
			}
		}
		rows[index_1] = &TabulateRow{Elements: normalized}
	}
	return rows
}

func createFromInt(data [][]int) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for index_1, arr := range data {
		row := make([]string, len(arr))
		for index, el := range arr {
			row[index] = strconv.Itoa(el)
		}
		rows[index_1] = &TabulateRow{Elements: row}
	}
	return rows
}

func createFromInt32(data [][]int32) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for index_1, arr := range data {
		row := make([]string, len(arr))
		for index, el := range arr {
			quoted := strconv.QuoteRuneToASCII(el)
			row[index] = quoted[1 : len(quoted)-1]
		}
		rows[index_1] = &TabulateRow{Elements: row}
	}
	return rows
}

func createFromInt64(data [][]int64) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for index_1, arr := range data {
		row := make([]string, len(arr))
		for index, el := range arr {
			row[index] = strconv.FormatInt(el, 10)
		}
		rows[index_1] = &TabulateRow{Elements: row}
	}
	return rows
}

func createFromBool(data [][]bool) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	for index_1, arr := range data {
		row := make([]string, len(arr))
		for index, el := range arr {
			row[index] = strconv.FormatBool(el)
		}
		rows[index_1] = &TabulateRow{Elements: row}
	}
	return rows
}

func createFromMapMixed(data map[string][]interface{}) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	return rows
}

func createFromMapString(data map[string][]string) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	return rows
}

func getRow(data []interface{}) *TabulateRow {
	normalized := make([]string, len(data))
	for index, element := range data {
		switch element.(type) {
		case int32:
			quoted := strconv.QuoteRuneToASCII(element.(int32))
			normalized[index] = quoted[1 : len(quoted)-1]
		case int:
			normalized[index] = strconv.Itoa(element.(int))
		case int64:
			normalized[index] = strconv.FormatInt(element.(int64), 10)
		case bool:
			normalized[index] = strconv.FormatBool(element.(bool))
		case float64:
			normalized[index] = strconv.FormatFloat(element.(float64), 'f', -1, 64)
		case uint64:
			normalized[index] = strconv.FormatUint(element.(uint64), 10)
		default:
			normalized[index] = fmt.Sprintf("%#v", element)
		}
	}
	return &TabulateRow{Elements: normalized}
}

func inSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
