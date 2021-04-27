# CSV to JSON

Take in a CSV file and spit out a JSON array

The CSV file should meet a few requirements: 
- have a header row with the keys for each column
- each row should have an entry for each column, even if empty
- each row should be its own line (newline delimited)
- each column should be delimited by a standard character

The default delimiter character is "," but others can be passed using the `-delimiter` flag.

The program can read from stdin, or an optional infile, and will output to stdout or an option outfile. This makes
it useful for chaining, especially by piping the output to something like `jq` for pretty printing the json.

## Install
```shell
# Clone the repo:
git clone git@github.com:jimmykodes/csv-to-json.git

# cd to repo
cd csv-to-json

# install
go install ctj.go
```

## Running
```shell
# basic usage
ctj [-delimiter CHAR] [in_file] [out_file]
# take a standard csv and put the json on stdout
ctj data.csv
# specify a different delimiter
ctj -delimiter '|' data.csv
# special case for tabs
ctj -delimiter tab data.csv
# save output json
ctj data.csv output.json
# suggestion: pipe to jq for nice json in formatting
ctj data.csv | jq 
# useful in a chain
curl "https://someUrl.com/get_a_csv" | ctj | jq
```
