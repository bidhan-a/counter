package counter

import (
	"testing"
)

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

func TestCounterUpdate(t *testing.T) {
	type expected struct {
		key   interface{}
		value int
	}
	tests := []struct {
		to       interface{}
		from     interface{}
		expected []expected
	}{
		{
			"abbcccddddeeeee",
			[]string{"a", "b", "c"},
			[]expected{
				{"a", 2},
				{"b", 3},
				{"c", 4},
				{"d", 4},
				{"e", 5},
			},
		},
		{
			[]int{1, 3, 3, 5, 5, 5},
			map[int]int{1: 2, 3: 1},
			[]expected{
				{1, 3},
				{3, 3},
				{5, 3},
			},
		},
	}
	for i, test := range tests {
		toCounter, err := NewCounter(test.to)
		if err != nil {
			t.Fatalf("tests[%d]. failed with error %q", i, err)
		}
		fromCounter, err := NewCounter(test.from)
		if err != nil {
			t.Fatalf("tests[%d]. failed with error %q", i, err)
		}

		toCounter.Update(fromCounter)
		for _, e := range test.expected {
			if toCounter[e.key] != e.value {
				t.Fatalf("tests[%d]. expected %d. got %d", i, e.value, toCounter[e.key])
			}
		}
	}

}

func TestCounterSubtract(t *testing.T) {
	type expected struct {
		key   interface{}
		value int
	}
	tests := []struct {
		to       interface{}
		from     interface{}
		expected []expected
	}{
		{
			"abbcccddddeeeee",
			[]string{"a", "b", "c", "f", "f"},
			[]expected{
				{"a", 0},
				{"b", 1},
				{"c", 2},
				{"d", 4},
				{"e", 5},
				{"f", -2},
			},
		},
		{
			[]int{1, 3, 3, 5, 5, 5},
			map[int]int{1: 2, 3: 1},
			[]expected{
				{1, -1},
				{3, 1},
				{5, 3},
			},
		},
	}
	for i, test := range tests {
		toCounter, err := NewCounter(test.to)
		if err != nil {
			t.Fatalf("tests[%d]. failed with error %q", i, err)
		}
		fromCounter, err := NewCounter(test.from)
		if err != nil {
			t.Fatalf("tests[%d]. failed with error %q", i, err)
		}

		toCounter.Subtract(fromCounter)
		for _, e := range test.expected {
			if toCounter[e.key] != e.value {
				t.Fatalf("tests[%d]. expected %d. got %d", i, e.value, toCounter[e.key])
			}
		}
	}

}

func TestCounterCopy(t *testing.T) {
	type expected struct {
		key   interface{}
		value int
	}

	counter1, _ := NewCounter("abbcccddddeeeee")
	counter2, _ := NewCounter([]int{1, 2, 2, 3, 3, 3})

	tests := []struct {
		counter  Counter
		expected []expected
	}{
		{
			counter1,
			[]expected{
				{"a", 1},
				{"b", 2},
				{"c", 3},
				{"d", 4},
				{"e", 5},
			},
		},
		{
			counter2,
			[]expected{
				{1, 1},
				{2, 2},
				{3, 3},
			},
		},
	}
	for i, test := range tests {
		copy := test.counter.Copy()
		for _, e := range test.expected {
			if copy[e.key] != e.value {
				t.Fatalf("tests[%d]. expected %d. got %d", i, e.value, copy[e.key])
			}
		}
	}

}
