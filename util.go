package duo

import "strconv"
import "errors"

func toByteArray(i interface{}) ([]byte, error) {
	if i == nil {
		return nil, nil
	}

	switch i.(type) {
	case string:
		return []byte(i.(string)), nil
	case []uint8:
		return i.([]uint8)[:], nil
	default:
		return nil, errors.New("invalid type")
	}
}

func toByteArrayArray(i interface{}) ([][]byte, error) {
	if i == nil {
		return nil, nil
	}

	var arr []interface{}
	var res [][]byte

	switch i.(type) {
	case []interface{}:
		arr = i.([]interface{})[:]
	default:
		return nil, errors.New("invalid type")
	}

	for _, v := range arr {
		tmp, err := toByteArray(v)
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}

	return res, nil
}

func toString(i interface{}) (string, error) {
	if i == nil {
		return "", nil
	}

	switch i.(type) {
	case int64:
		v, _ := i.(int64)
		return strconv.FormatInt(v, 10), nil
	case string:
		return i.(string), nil
	case []uint8:
		return string(i.([]uint8)[:]), nil
	default:
		return "", errors.New("invalid type")
	}
}

func toInt64(i interface{}) (int64, error) {
	if i == nil {
		return 0, nil
	}

	switch i.(type) {
	case int64:
		return i.(int64), nil
	default:
		return 0, errors.New("invalid type")
	}
}
