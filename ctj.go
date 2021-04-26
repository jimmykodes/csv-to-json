package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	delimiter   string
	inFilePath  string
	outFilePath string
)

func main() {
	flag.StringVar(&delimiter, "delimiter", ",", "source file delimiter")
	flag.Parse()
	if delimiter == "tab" {
		delimiter = "\t"
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
	inData, err := ioutil.ReadAll(inFile)
	if err != nil {
		checkErr(err)
	}
	rows := strings.Split(string(inData), "\n")
	headerRow := strings.Split(strings.TrimSpace(rows[0]), delimiter)
	var data []interface{}
	for _, row := range rows[1:] {
		obj := make(map[string]interface{})
		row = strings.TrimSpace(row)
		if row == "" {
			// empty row, continue
			continue
		}
		dataRow := strings.Split(row, delimiter)
		if len(dataRow) != len(headerRow) {
			checkErr(fmt.Errorf("column len missmatch:\n%v\n%v", headerRow, dataRow))
		}
		for j, header := range headerRow {
			obj[header] = dataRow[j]
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
