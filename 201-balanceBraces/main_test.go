package main

import "testing"

func Test_isBracketBalancedAdvanced(t *testing.T) {
	type test struct {
		input  string
		output bool
	}

	tests := []test{
		{"()", true},
		{"(())", true},
		{"))((", false},
		{"()(())", true},
	}

	for _, v := range tests {
		if v.output != IsBracketBalancedAdvanced(v.input) {
			t.Error("Expected", v.output)
		}
	}
}

func Benchmark_IsBracketBalancedAdvanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsBracketBalancedAdvanced("()(())")
	}
}
