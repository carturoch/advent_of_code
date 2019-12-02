package main

import "testing"

func TestFuelPerModule(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Zero mass",
			args{0},
			0,
		},
		{
			"Smaller mass",
			args{2},
			0,
		},
		{
			"Regular mass",
			args{12},
			2,
		},
		{
			"Large mass",
			args{100756},
			33583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuelPerModule(tt.args.mass); got != tt.want {
				t.Errorf("FuelPerModule() = %v, want %v", got, tt.want)
			}
		})
	}
}
