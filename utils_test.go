package osgridconverter

import (
	"testing"
)

func TestToDegrees(t *testing.T) {
	testcases := []struct {
		input    float64
		expected float64
	}{
		{
			input:    1.0,
			expected: 57.29577951308232,
		},
		{
			input:    0.5,
			expected: 28.64788975654116,
		},
		{
			input:    -1.0,
			expected: -57.29577951308232,
		},
	}

	for _, testcase := range testcases {
		d := toDegrees(testcase.input)

		if d != testcase.expected {
			t.Errorf("Error coverting radians to degrees: expected %v, got %v", testcase.expected, d)
		}
	}
}

func TestToRadians(t *testing.T) {
	testcases := []struct {
		input    float64
		expected float64
	}{
		{
			input:    57.29577951308232,
			expected: 1.0,
		},
		{
			input:    28.64788975654116,
			expected: 0.5,
		},
		{
			input:    -57.29577951308232,
			expected: -1.0,
		},
	}

	for _, testcase := range testcases {
		d := toRadians(testcase.input)

		if d != testcase.expected {
			t.Errorf("Error coverting degrees to radians: expected %v, got %v", testcase.expected, d)
		}
	}
}
