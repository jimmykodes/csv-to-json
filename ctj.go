package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	delimiter   string
	inFilePath  string
	outFilePath string
	comma       rune
)

func main() {
	flag.StringVar(&delimiter, "delimiter", ",", "source file delimiter")
	flag.Parse()
	if delimiter == "tab" {
		comma = '\t'
	} else {
		comma = rune(delimiter[0])
	}
	args := flag.Args()
	if len(args) < 1 {
		checkErr("missing source file or destination file")
	}
	inFilePath = args[0]
	if len(args) == 2 {
		outFilePath = args[1]
	}
	inFile, err := os.Open(inFilePath)
	if err != nil {
		checkErr(err)
	}
	defer inFile.Close()
	reader := csv.NewReader(inFile)
	reader.Comma = comma
	reader.LazyQuotes = true
	rows, err := reader.ReadAll()
	if err != nil {
		checkErr(err)
	}
	headerRow := rows[0]
	var data []interface{}
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
	var out io.Writer
	if outFilePath != "" {
		out, err = os.Create(outFilePath)
		checkErr(err)
	} else {
		out = os.Stdout
	}

	err = json.NewEncoder(out).Encode(data)
	checkErr(err)
}

func checkErr(msg interface{}) {
	if msg != nil {
		panic(msg)
	}
}
