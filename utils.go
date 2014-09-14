package main

import "fmt"
import "strconv"

func normalize(data []interface{}) []*TabulateRow {
	rows := make([]*TabulateRow, len(data))
	normalized := make([]string, len(data))

	for index, element := range data {
		switch v := element.(type) {
		case int:
			normalized = append(normalized, strconv.Itoa(element.(int)))
		case int64:
			normalized = append(normalized, strconv.FormatInt(element.(int64), 10))
		case bool:
			normalized = append(normalized, strconv.FormatBool(element.(bool)))
		case float64:
			normalized = append(normalized, strconv.FormatFloat(element.(float64), 'f', -1, 64))
		case uint64:
			normalized = append(normalized, strconv.FormatUint(element.(uint64), 10))
		case []interface{}:
			rows[index] = normalize(data)
		default:
			fmt.Println(v)
		}
	}

	return []TabulateRow{}
}
