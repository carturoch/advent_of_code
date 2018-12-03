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

func TestFreqTable_GetCharsMoreFreqThan(t *testing.T) {
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
			[]string{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTrios := tt.table.GetCharsMoreFreqThan(tt.args.val); !reflect.DeepEqual(gotTrios, tt.wantTrios) {
				t.Errorf("FreqTable.GetCharsMoreFreqThan() = %v, want %v", gotTrios, tt.wantTrios)
			}
		})
	}
}
