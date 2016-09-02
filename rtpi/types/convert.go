package types

import (
	"fmt"
	"strconv"
)

func unexpectedValueError(v interface{}) error {
	return fmt.Errorf("unexpected value: %s, type: %T", v, v)
}

func toInt(v interface{}) (int, error) {
	vv, ok := v.(string)
	if !ok {
		return 0, unexpectedValueError(v)
	}
	i, err := strconv.Atoi(vv)
	if err != nil {
		return 0, fmt.Errorf("can't convert string to int: %s", v)
	}
	return i, nil
}

func toString(v interface{}) (string, error) {
	vv, ok := v.(string)
	if !ok {
		return "", unexpectedValueError(v)
	}
	return vv, nil
}

func toSlice(v interface{}) ([]interface{}, error) {
	vv, ok := v.([]interface{})
	if !ok {
		return nil, unexpectedValueError(v)
	}
	return vv, nil
}
