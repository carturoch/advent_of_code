package main

import (
	"reflect"
	"testing"
)

func TestParseClaim(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    Claim
		wantErr bool
	}{
		{
			"Returns error when line has unexpected format",
			args{"234 @ z ,2 3x3"},
			Claim{},
			true,
		},
		{
			"Parses a claim",
			args{"#234 @ 3,2: 3x3"},
			Claim{
				ID:     "#234",
				Left:   3,
				Top:    2,
				Width:  3,
				Height: 3,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseClaim(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseClaim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseClaim() = %v, want %v", got, tt.want)
			}
		})
	}
}
