package godash

import "testing"

func TestAbs(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		expected float64
	}{
		{0, 0},
		{-1, 1},
		{10, 10},
		{3.14, 3.14},
		{-96, 96},
		{-10e-12, 10e-12},
	}
	for _, test := range tests {
		actual := Abs(test.param)
		if actual != test.expected {
			t.Errorf("Expected Abs(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestSign(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		expected float64
	}{
		{0, 0},
		{-1, -1},
		{10, 1},
		{3.14, 1},
		{-96, -1},
		{-10e-12, -1},
	}
	for _, test := range tests {
		actual := Sign(test.param)
		if actual != test.expected {
			t.Errorf("Expected Sign(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}
