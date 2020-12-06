package main

import (
	"reflect"
	"testing"
)

func Test_ParseWireIntoSteps(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []Step
	}{
		{
			"Parse two steps step",
			args{"U23 D12"},
			[]Step{
				Step{'U', 23},
				Step{'D', 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseWireIntoSteps(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseWireIntoSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
