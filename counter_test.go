package counter

import (
	"testing"
)

func compareIntSlices(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestNewCounterFromString(t *testing.T) {
	type expected struct {
		key   string
		value int
	}
	tests := []struct {
		value    string
		expected []expected
	}{
		{
			"abbcccddddeeeee",
			[]expected{
				{"a", 1},
				{"b", 2},
				{"c", 3},
				{"d", 4},
				{"e", 5},
			},
		},
	}

	for i, test := range tests {
		counter, err := NewCounter(test.value)
		if err != nil {
			t.Fatalf("tests[%d]. failed with error %q", i, err)
		}
		for _, e := range test.expected {
			if counter[e.key] != e.value {
				t.Fatalf("tests[%d]. expected %d. got %d", i, e.value, counter[e.key])
			}
		}
	}
}

func TestNewCounterFromIntSlice(t *testing.T) {
	type expected struct {
		key   int
		value int
	}
	tests := []struct {
		value    []int
		expected []expected
	}{
		{
			[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5},
			[]expected{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{5, 5},
			},
		},
	}

	for i, test := range tests {
		counter, err := NewCounter(test.value)
		if err != nil {
			t.Fatalf("tests[%d]. failed with error %q", i, err)
		}
		for _, e := range test.expected {
			if counter[e.key] != e.value {
				t.Fatalf("tests[%d]. expected %d. got %d", i, e.value, counter[e.key])
			}
		}
	}
}

func TestNewCounterFromFloatSlice(t *testing.T) {
	type expected struct {
		key   float64
		value int
	}
	tests := []struct {
		value    []float64
		expected []expected
	}{
		{
			[]float64{
				1.1, 2.2, 2.2, 3.3, 3.3,
				3.3, 4.4, 4.4, 4.4, 4.4,
				5.5, 5.5, 5.5, 5.5, 5.5,
			},
			[]expected{
				{1.1, 1},
				{2.2, 2},
				{3.3, 3},
				{4.4, 4},
				{5.5, 5},
			},
		},
	}

	for i, test := range tests {
		counter, err := NewCounter(test.value)
		if err != nil {
			t.Fatalf("tests[%d]. failed with error %q", i, err)
		}
		for _, e := range test.expected {
			if counter[e.key] != e.value {
				t.Fatalf("tests[%d]. expected %d. got %d", i, e.value, counter[e.key])
			}
		}
	}
}

func TestNewCounterFromMap(t *testing.T) {
	type expected struct {
		key   interface{}
		value int
	}
	tests := []struct {
		value    map[interface{}]int
		expected []expected
	}{
		{
			map[interface{}]int{
				"one":   1,
				"two":   2,
				"three": 3,
				"four":  4,
				"five":  5,
			},
			[]expected{
				{"one", 1},
				{"two", 2},
				{"three", 3},
				{"four", 4},
				{"five", 5},
			},
		},
	}

	for i, test := range tests {
		counter, err := NewCounter(test.value)
		if err != nil {
			t.Fatalf("tests[%d]. failed with error %q", i, err)
		}
		for _, e := range test.expected {
			if counter[e.key] != e.value {
				t.Fatalf("tests[%d]. expected %d. got %d", i, e.value, counter[e.key])
			}
		}
	}
}
