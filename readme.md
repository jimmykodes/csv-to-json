#CSV to JSON

Take in a CSV file and spit out a JSON array

The CSV file can be any format, ie tsv, pipe separated, etc. as long as each row is newline separated and is a standalone
entry, with a consistent delimiter character. 

#### note
Currently, this lacks support for quoted delimiters... sorry...

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
```