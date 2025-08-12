package main

import (
	"bufio"
	"os"
	"testing"
)

func TestEmptyLines(t *testing.T) {
	options := formatterOptions{
		EndOfLines:     "\n",
		IndentSize:     4,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/empty-lines.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/empty-lines.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/empty-lines.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestEmptyLines failed")
	}
}
