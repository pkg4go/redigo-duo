package duo

import "errors"

func toByteArray(i interface{}) ([]byte, error) {
	switch i.(type) {
	case string:
		return []byte(i.(string)), nil
	case []uint8:
		return i.([]uint8)[:], nil
	default:
		return nil, errors.New("invalid type")
	}
}
