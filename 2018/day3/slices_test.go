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

func TestInitLayout(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want Layout
	}{
		{
			"Inits a 2x2",
			args{2},
			Layout{
				{Empty, Empty},
				{Empty, Empty},
			},
		},
		{
			"Inits a 3x3",
			args{3},
			Layout{
				{Empty, Empty, Empty},
				{Empty, Empty, Empty},
				{Empty, Empty, Empty},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitLayout(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitLayout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyClaim(t *testing.T) {
	type args struct {
		c Claim
		l *Layout
	}
	tests := []struct {
		name string
		args args
		want *Layout
	}{
		{
			"Applies a claim to an empty layout",
			args{
				Claim{
					ID:     "123",
					Left:   1,
					Top:    1,
					Width:  2,
					Height: 1,
				},
				&Layout{
					{Empty, Empty, Empty},
					{Empty, Empty, Empty},
					{Empty, Empty, Empty},
				},
			},
			&Layout{
				{Empty, Empty, Empty},
				{Empty, Marked, Marked},
				{Empty, Empty, Empty},
			},
		},
		{
			"Marks collisions",
			args{
				Claim{
					ID:     "123",
					Left:   1,
					Top:    1,
					Width:  2,
					Height: 1,
				},
				&Layout{
					{Empty, Empty, Empty},
					{Empty, Marked, Empty},
					{Empty, Empty, Empty},
				},
			},
			&Layout{
				{Empty, Empty, Empty},
				{Empty, Conflicted, Marked},
				{Empty, Empty, Empty},
			},
		},
		{
			"Ignores previous collisions",
			args{
				Claim{
					ID:     "123",
					Left:   1,
					Top:    1,
					Width:  2,
					Height: 1,
				},
				&Layout{
					{Empty, Empty, Empty},
					{Empty, Conflicted, Empty},
					{Empty, Empty, Empty},
				},
			},
			&Layout{
				{Empty, Empty, Empty},
				{Empty, Conflicted, Marked},
				{Empty, Empty, Empty},
			},
		},
		{
			"Ignores layout overflows",
			args{
				Claim{
					ID:     "123",
					Left:   1,
					Top:    1,
					Width:  2,
					Height: 1,
				},
				&Layout{
					{Empty, Empty},
					{Empty, Marked},
				},
			},
			&Layout{
				{Empty, Empty},
				{Empty, Conflicted},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyClaim(tt.args.c, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyClaim() = %v, want %v", got, tt.want)
			}
		})
	}
}
