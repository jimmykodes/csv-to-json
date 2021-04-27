package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		delimiter string
		src       io.ReadCloser
		dest      io.WriteCloser
		comma     rune
	)

	flag.StringVar(&delimiter, "delimiter", ",", "source file delimiter")
	flag.Parse()
	if delimiter == "tab" {
		comma = '\t'
	} else {
		comma = rune(delimiter[0])
	}
	args := flag.Args()
	var err error
	switch len(args) {
	case 0:
		src = os.Stdin
		dest = os.Stdout
	case 1:
		src, err = os.Open(args[0])
		checkErr(err)
		dest = os.Stdout
	case 2:
		src, err = os.Open(args[0])
		checkErr(err)
		dest, err = os.Create(args[1])
		checkErr(err)
	default:
		panic("invalid number of args")
	}
	defer src.Close()
	defer dest.Close()

	reader := csv.NewReader(src)
	reader.Comma = comma
	reader.LazyQuotes = true
	rows, err := reader.ReadAll()
	checkErr(err)

	headerRow := rows[0]
	var data []map[string]interface{}
	for _, row := range rows[1:] {
		obj := make(map[string]interface{})
		if len(row) != len(headerRow) {
			checkErr(fmt.Errorf("column len missmatch:\n%v\n%v", headerRow, row))
		}
		for j, header := range headerRow {
			obj[header] = row[j]
		}
		data = append(data, obj)
	}

	err = json.NewEncoder(dest).Encode(data)
	checkErr(err)
}

func checkErr(msg interface{}) {
	if msg != nil {
		panic(msg)
	}
}
