package helper

import (
	"github.com/merge/pkg/interval"
	"testing"
)

func TestPrint(t *testing.T) {
	type args struct {
		title   string
		results []*interval.Interval
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"print", args{
			title:   "",
			results: []*interval.Interval{
				&interval.Interval{
					Low:  1,
					High: 2,
				},
				&interval.Interval{
					Low:  3,
					High: 4,
				},
			},
		}, "[1,2][3,4]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Print(tt.args.title, tt.args.results); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}