package binary

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "true", args: args{num: 1}, want: true},
		{name: "true", args: args{num: 2}, want: false},
		{name: "true", args: args{num: 4}, want: true},
		{name: "true", args: args{num: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
