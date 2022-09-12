package twopoints

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{s: "abcabca"}, want: 3},
		{name: "test", args: args{s: "a"}, want: 1},
		{name: "test", args: args{s: "aaaaa"}, want: 1},
		{name: "test", args: args{s: "pwwkew"}, want: 3},
		{name: "test", args: args{s: "wwpkew"}, want: 4},
		{name: "test", args: args{s: ""}, want: 0},
		{name: "test", args: args{s: " "}, want: 1},
		{name: "test", args: args{s: "abba"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring0301(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, want: 49},
		{name: "test", args: args{height: []int{1, 1}}, want: 1},
		{name: "test", args: args{height: []int{1, 2, 2, 1}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
