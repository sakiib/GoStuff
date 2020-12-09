package main

import "testing"

func TestTotal(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "first test",
			args: args{
				[]int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "second test",
			args: args{
				args: []int{2, 3, 4},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Total(tt.args.args...); got != tt.want {
				t.Errorf("Total() = %v, want %v", got, tt.want)
			}
		})
	}
}
