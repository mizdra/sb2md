package parser

import (
	"reflect"
	"testing"
)

func TestParsePage(t *testing.T) {
	type args struct {
		page_text string
	}
	tests := []struct {
		name string
		args args
		want Page
	}{
		{
			name: "basic",
			args: args{page_text: "title\nbody1\nbody2"},
			want: Page{title: "title", lines: []Line{{text: "title"}, {text: "body1"}, {text: "body2"}}},
		},
		{
			name: "trailing newline",
			args: args{page_text: "title\nbody1\n"},
			want: Page{title: "title", lines: []Line{{text: "title"}, {text: "body1"}, {text: ""}}},
		},
		{
			name: "title only",
			args: args{page_text: "title"},
			want: Page{title: "title", lines: []Line{{text: "title"}}},
		},
		{
			name: "empty title",
			args: args{page_text: "\nbody1"},
			want: Page{title: "", lines: []Line{{text: ""}, {text: "body1"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParsePage(tt.args.page_text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePage() = %v, want %v", got, tt.want)
			}
		})
	}
}
