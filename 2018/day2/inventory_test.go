package main

import (
	"reflect"
	"testing"
)

func TestCountChars(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			"Returns empty map for an empty string",
			args{""},
			map[string]int{},
		},
		{
			"Counts every char",
			args{"aba"},
			map[string]int{
				"a": 2,
				"b": 1,
			},
		},
		{
			"Counts spaces also",
			args{"aba  bbb"},
			map[string]int{
				"a": 2,
				"b": 4,
				" ": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountChars(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreqTable_GetCharsEqFreq(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name      string
		table     FreqTable
		args      args
		wantTrios []string
	}{
		{
			"Detects equal and greater values",
			FreqTable{
				"a": 2,
				"b": 1,
				"c": 3,
			},
			args{2},
			[]string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTrios := tt.table.GetCharsEqFreq(tt.args.val); !reflect.DeepEqual(gotTrios, tt.wantTrios) {
				t.Errorf("FreqTable.GetCharsEqFreq() = %v, want %v", gotTrios, tt.wantTrios)
			}
		})
	}
}

func TestCommon(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			"Returns empty when there are no matches",
			args{
				"abc",
				"zyx",
			},
			"",
		},
		{
			"Matches chars in same position",
			args{
				"abcpq",
				"zbcxr",
			},
			"bc",
		},
		{
			"Support different size strings",
			args{
				"abcde",
				"ab",
			},
			"ab",
		},
		{
			"Returns empty when one string is empty",
			args{
				"0",
				"ab",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Common(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("Common() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
