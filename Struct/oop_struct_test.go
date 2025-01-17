package main

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum1",
			args: args{
				numbers: []int{1, 2, 3},
			},
			want: 10,
		},
		{
			name: "sum2",
			args: args{
				numbers: []int{1, 2, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
