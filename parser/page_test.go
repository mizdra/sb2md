package parser

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestParsePage(t *testing.T) {
	type args struct {
		page_text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "basic",
			args: args{page_text: "title\nbody1\nbody2"},
		},
		{
			name: "trailing newline",
			args: args{page_text: "title\nbody1\n"},
		},
		{
			name: "title only",
			args: args{page_text: "title"},
		},
		{
			name: "empty title",
			args: args{page_text: "\nbody1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snaps.MatchSnapshot(t, ParsePage(tt.args.page_text))
		})
	}
}
