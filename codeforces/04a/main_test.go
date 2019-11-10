package main

import "testing"

func Test_hasEvenFactor(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "no",
			num:  0,
			want: "NO",
		},
		{
			name: "no",
			num:  1,
			want: "NO",
		},
		{
			name: "no",
			num:  2,
			want: "NO",
		},
		{
			name: "no",
			num:  3,
			want: "NO",
		},
		{
			name: "yes",
			num:  4,
			want: "YES",
		},
		{
			name: "yes",
			num:  6,
			want: "YES",
		},
		{
			name: "yes",
			num:  7,
			want: "NO",
		},
		{
			name: "yes",
			num:  8,
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasEvenFactor(tt.num); got != tt.want {
				t.Errorf("hasEvenFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
