package rgtm

import (
	"reflect"
	"regexp"
	"testing"
)

func TestStringToMap(t *testing.T) {

	m1 := make(map[string]string)

	m1[""] = "Marvin J Wendt"
	m1["first"] = "Marvin"
	m1["middle"] = "J"
	m1["last"] = "Wendt"

	type args struct {
		r *regexp.Regexp
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{name: "#1", args: args{r: regexp.MustCompile("(?P<first>\\w+) (?P<middle>.+) (?P<last>\\w+)"), s: "Marvin J Wendt"}, want: m1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToMap(tt.args.r, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
