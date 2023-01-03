package parser

type Line struct {
	text string
}

func ParseLine(line_text string) Line {
	return Line{text: line_text}
}
