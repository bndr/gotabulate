# Gotabulate - Easily tabulate Data
[![GoDoc](https://godoc.org/github.com/bndr/gotabulate?status.svg)](https://godoc.org/github.com/bndr/gotabulate)
[![Build Status](https://travis-ci.org/bndr/gotabulate.svg?branch=master)](https://travis-ci.org/bndr/gotabulate)

## Summary

Go-Tabulate - Generic Go Library for easy tabulation of your data. 

## Installation

    go get github.com/bndr/gotabulate

## Description
Supported Data Types:
- 2D Array of Int, Int64, Float64, String, interface{}
- Map of String, interface{} (Keys will be used as header)

## Usage
```go

// Create Some Fake Rows
row_1 := []interface{}{"john", 20, "ready"}
row_2 := []interface{}{"bndr", 23, "ready"}

// Create an object from 2D interface array
t := gotabulate.Create([][]interface{}{row_1, row_2})

// Set the Headers (optional)
t.SetHeaders([]string{"age", "status"})

// Set the Empty String (optional)
t.SetEmptyString("None")

// Set Align (Optional)
t.SetAlign("right")

// Print the result: grid, or simple
fmt.Println(t.Render("grid"))

+---------+--------+-----------+
|         |    age |    status |
+=========+========+===========+
|    john |     20 |     ready |
+---------+--------+-----------+
|    bndr |     23 |     ready |
+---------+--------+-----------+

```

## Examples

```
t := gotabulate.Create([][]string{STRING_ARRAY, STRING_ARRAY})
t.SetHeaders(HEADERS) // If not headers are set, the first row will be used.
t.SetEmptyString("None") // Set what will be printed in the empty cell
rendered_string := t.Render("simple") // Render() will return a string

Simple Table
----------------------  ----------------------  ----------------------  -------------  -------------
             Header 1                Header 2                Header 3       Header 4       Header 5 
----------------------  ----------------------  ----------------------  -------------  -------------
          test string           test string 2                    test            row           bndr 

          test string           test string 2                    test            row           bndr 

    4th element empty       4th element empty       4th element empty           None           None 
----------------------  ----------------------  ----------------------  -------------  -------------

Grid Table
+-------------+-------------+-------------+-------------+-------------+
|    Header 1 |    Header 2 |    Header 3 |    Header 4 |    Header 5 |
+=============+=============+=============+=============+=============+
|       10.01 |      12.002 |      -123.5 |    20.00005 |        1.01 |
+-------------+-------------+-------------+-------------+-------------+
|       10.01 |      12.002 |      -123.5 |    20.00005 |        1.01 |
+-------------+-------------+-------------+-------------+-------------+
|       10.01 |      12.002 |      -123.5 |    20.00005 |        None |
+-------------+-------------+-------------+-------------+-------------+

Padded Headers:
+----------------------+----------------------+----------------------+-------------+-------------+
|                      |             Header 1 |             header 2 |    header 3 |    header 4 |
+======================+======================+======================+=============+=============+
|          test string |        test string 2 |                 test |         row |        bndr |
+----------------------+----------------------+----------------------+-------------+-------------+
|          test string |        test string 2 |                 test |         row |        bndr |
+----------------------+----------------------+----------------------+-------------+-------------+
|    4th element empty |    4th element empty |    4th element empty |        None |        None |
+----------------------+----------------------+----------------------+-------------+-------------+

Align Center:
+-------------+-------------+-------------+-------------+-------------+
|   Header 1  |   Header 2  |   Header 3  |   Header 4  |   Header 5  |
+=============+=============+=============+=============+=============+
|    10.01    |    12.002   |    -123.5   |   20.00005  |     1.01    |
+-------------+-------------+-------------+-------------+-------------+
|    10.01    |    12.002   |    -123.5   |   20.00005  |     1.01    |
+-------------+-------------+-------------+-------------+-------------+
|    10.01    |    12.002   |    -123.5   |   20.00005  |     1.01    |
+-------------+-------------+-------------+-------------+-------------+

Align Left:
+-------------+-------------+-------------+-------------+-------------+
| Header 1    | Header 2    | Header 3    | Header 4    | Header 5    |
+=============+=============+=============+=============+=============+
| 10.01       | 12.002      | -123.5      | 20.00005    | 1.01        |
+-------------+-------------+-------------+-------------+-------------+
| 10.01       | 12.002      | -123.5      | 20.00005    | 1.01        |
+-------------+-------------+-------------+-------------+-------------+
| 10.01       | 12.002      | -123.5      | 20.00005    | 1.01        |
+-------------+-------------+-------------+-------------+-------------+
```

### TODO
- Max size of cell
- linewrap in cell