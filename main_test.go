package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("it should return the longest string in each column", func(t *testing.T) {
		rows := [][]string{
			{"_", "long", "_"},
			{"_", "_", "longer"},
			{"longest", "_", "_"},
		}

		longest := getLongest(rows)

		if longest[0] != 7 {
			t.Errorf("want 7, got %d", longest[0])
		}

		if longest[1] != 4 {
			t.Errorf("want 4, got %d", longest[1])
		}

		if longest[2] != 6 {
			t.Errorf("want 6, got %d", longest[2])
		}
	})

	t.Run("it should build a markdown table", func(t *testing.T) {
		rows := [][]string{
			{"a", "b", "c"},
			{"1", "2", "3"},
			{"4", "5", "6789"},
		}

		longest := getLongest(rows)

		var builder strings.Builder

		buildMarkdown(&builder, rows, longest)

		got := strings.Split(builder.String(), "\n")

		want := []string{
			"| a | b | c    |",
			"| - | - | ---- |",
			"| 1 | 2 | 3    |",
			"| 4 | 5 | 6789 |",
		}

		for i, row := range want {
			if row != got[i] {
				t.Errorf("want '%s', got '%s'", row, got[i])
			}
		}
	})
}
