package counter

import (
	"fmt"
	"reflect"
	"strings"
)

// T represents an empty interface type.
type T interface{}

// Counter represents a counter.
type Counter map[T]int

// NewCounter creates a new counter.
func NewCounter(arg T) (counter Counter, err error) {
	value := reflect.ValueOf(arg)
	switch value.Kind() {
	case reflect.String:
		substrings := strings.Split(value.String(), "")
		if len(substrings) > 0 {
			tSlice, e := toTSlice(substrings)
			if e != nil {
				err = e
			}
			counter = make(Counter)
			createMappingFromSlice(counter, tSlice)
		}
	case reflect.Slice:
		tSlice, e := toTSlice(value.Interface())
		if err != nil {
			err = e
		}
		counter = make(Counter)
		createMappingFromSlice(counter, tSlice)
	case reflect.Map:
		// Only maps of type map[T]int are valid
		elType := reflect.TypeOf(arg).Elem()
		if elType.Kind() != reflect.Int {
			err = fmt.Errorf("the map element must be of type 'int'")
		}
		tMap, e := toTMap(value.Interface())
		if e != nil {
			err = e
		}
		counter = Counter(tMap)
	default:
		err = fmt.Errorf("unsupported argument")
	}
	return
}

// Update updates the counter using another counter.
func (counter Counter) Update(from Counter) {
	for k, v := range from {
		if _, ok := counter[k]; ok {
			counter[k] += v
		} else {
			counter[k] = v
		}
	}
}

// Subtract subtracts counts in the counter using
// counts from another counter.
func (counter Counter) Subtract(from Counter) {
	for k, v := range from {
		if _, ok := counter[k]; ok {
			counter[k] -= v
		} else {
			counter[k] = -v
		}
	}
}

// Copy creates a copy of the counter.
func (counter Counter) Copy() Counter {
	copy := make(Counter)
	for k, v := range counter {
		copy[k] = v
	}
	return copy
}

func createMappingFromSlice(counter Counter, slice []T) {
	for _, s := range slice {
		if _, ok := counter[s]; ok {
			counter[s]++
		} else {
			counter[s] = 1
		}
	}
}

func toTSlice(arg T) (out []T, err error) {
	slice := reflect.ValueOf(arg)
	if slice.Kind() != reflect.Slice {
		err = fmt.Errorf("value is not a slice")
		return
	}
	c := slice.Len()
	out = make([]T, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, nil
}

func toTMap(arg T) (out map[T]int, err error) {
	m := reflect.ValueOf(arg)
	if m.Kind() != reflect.Map {
		err = fmt.Errorf("value is not a map")
		return
	}
	keys := m.MapKeys()
	out = make(map[T]int)
	for _, key := range keys {
		out[key.Interface()] = int(m.MapIndex(key).Int())
	}
	return out, nil
}
