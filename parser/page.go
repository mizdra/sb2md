package parser

import "strings"

type Page struct {
	title string
	lines []Line
}

func ParsePage(page_text string) Page {
	line_texts := strings.Split(page_text, "\n")
	title := line_texts[0]
	lines := make([]Line, len(line_texts))
	for i, line_text := range line_texts {
		lines[i] = ParseLine(line_text)
	}
	return Page{title: title, lines: lines}
}
