package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	source string
}

func main() {
	config := Config{}

	flag.StringVar(&config.source, "src", "", "The path to the source CSV")
	flag.Parse()

	if config.source == "" {
		panic("Source cannot be empty")
	}

	file, err := os.Open(config.source)
	if err != nil {
		panic(err)
	}

	var rows [][]string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), ","))
	}

	longest := getLongest(rows)

	var builder strings.Builder

	buildMarkdown(&builder, rows, longest)

	fmt.Print(builder.String())
}

func getLongest(rows [][]string) map[int]int {
	// map[column]length
	longest := map[int]int{}

	for col := 0; col < len(rows[0]); col++ {
		for row := 0; row < len(rows); row++ {
			if len(rows[row][col]) > longest[col] {
				longest[col] = len(rows[row][col])
			}
		}
	}

	return longest
}

func buildMarkdown(builder *strings.Builder, rows [][]string, longest map[int]int) {
	for row := 0; row < len(rows); row++ {
		for col := 0; col < len(rows[row]); col++ {
			cell := rows[row][col]

			if col == 0 {
				builder.WriteString("| ")
			} else {
				builder.WriteString(" ")
			}

			builder.WriteString(cell)

			for range longest[col] - len(cell) {
				builder.WriteString(" ")
			}

			builder.WriteString(" |")
		}

		builder.WriteString("\n")

		if row == 0 {
			for col := 0; col < len(rows[row]); col++ {
				if col == 0 {
					builder.WriteString("| ")
				} else {
					builder.WriteString(" ")
				}

				for range longest[col] {
					builder.WriteString("-")
				}

				builder.WriteString(" |")
			}

			builder.WriteString("\n")
		}
	}
}
