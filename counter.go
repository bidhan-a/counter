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
func NewCounter(arg T) (Counter, error) {
	var counter Counter
	value := reflect.ValueOf(arg)
	switch value.Kind() {
	case reflect.String:
		substrings := strings.Split(value.String(), "")
		if len(substrings) > 0 {
			tSlice, err := toTSlice(substrings)
			if err != nil {
				return nil, err
			}
			counter = make(Counter)
			createMappingFromSlice(counter, tSlice)
		}
	case reflect.Slice:
		tSlice, err := toTSlice(value.Interface())
		if err != nil {
			return nil, err
		}
		counter = make(Counter)
		createMappingFromSlice(counter, tSlice)
	case reflect.Map:
		// Only maps of type map[T]int are valid
		elType := reflect.TypeOf(arg).Elem()
		if elType.Kind() != reflect.Int {
			return nil, fmt.Errorf("the map element must be of type 'int'")
		}
		tMap, err := toTMap(value.Interface())
		if err != nil {
			return nil, err
		}
		counter = Counter(tMap)
	}

	return counter, nil
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
