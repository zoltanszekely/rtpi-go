package types

import (
	"fmt"
	"strconv"
)

func unexpectedValueError(v interface{}) error {
	return fmt.Errorf("unexpected value: %s, type: %T", v, v)
}

func toInt(v interface{}) (int, error) {
	switch vv := v.(type) {
	case string:
		i, err := strconv.Atoi(vv)
		if err != nil {
			return 0, fmt.Errorf("can't convert string to int: %s", v)
		}
		return i, nil
	}
	return 0, unexpectedValueError(v)
}

func toString(v interface{}) (string, error) {
	switch vv := v.(type) {
	case string:
		return vv, nil
	}
	return "", unexpectedValueError(v)
}

func toSlice(v interface{}) ([]interface{}, error) {
	switch vv := v.(type) {
	case []interface{}:
		return vv, nil
	}
	return nil, unexpectedValueError(v)
}
