package strings

import (
	"reflect"
	"testing"
)

func Test_letterCombinationsV1(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test for phone", args: args{digits: "23"}, want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{name: "test for phone", args: args{digits: "456"}, want: []string{"gjm", "gjn", "gjo", "gkm", "gkn", "gko", "glm", "gln", "glo", "hjm", "hjn", "hjo", "hkm", "hkn", "hko", "hlm", "hln", "hlo", "ijm", "ijn", "ijo", "ikm", "ikn", "iko", "ilm", "iln", "ilo"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinationsV1(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinationsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
