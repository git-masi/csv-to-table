# csv-to-table

## About

Convert CSV files to markdown tables.

```csv
a,b,c
1,2,3
4,5,6789
```

becomes:

```md
| a | b | c    |
| - | - | ---- |
| 1 | 2 | 3    |
| 4 | 5 | 6789 |
```

## How to use

Specify the source CSV using the `-src` flag.

```sh
go run ./... -src='some/path/here.csv'
```