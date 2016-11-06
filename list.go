package duo

import "github.com/pkg4go/convert"

func (c conn) RPop(k interface{}) ([]byte, error) {
	rep, err := c.connection.Do("RPOP", k)
	if err != nil {
		return nil, err
	}

	return toByteArray(rep)
}

func (c conn) RPush(k interface{}, vals ...interface{}) (int, error) {
	var args []interface{}
	args = append(args, k)
	args = append(args, vals...)

	rep, err := c.connection.Do("RPUSH", args...)

	if err != nil {
		return 0, err
	}

	return convert.Int(rep)
}

func (c conn) RPushX(k, v interface{}) (int, error) {
	rep, err := c.connection.Do("RPUSHX", k, v)
	if err != nil {
		return 0, err
	}

	return convert.Int(rep)
}

func (c conn) LPop(k interface{}) ([]byte, error) {
	rep, err := c.connection.Do("LPOP", k)
	if err != nil {
		return nil, err
	}

	return toByteArray(rep)
}

func (c conn) LPush(k interface{}, vals ...interface{}) (int, error) {
	var args []interface{}
	args = append(args, k)
	args = append(args, vals...)

	rep, err := c.connection.Do("LPUSH", args...)

	if err != nil {
		return 0, err
	}

	return convert.Int(rep)
}

func (c conn) LPushX(k, v interface{}) (int, error) {
	rep, err := c.connection.Do("LPUSHX", k, v)
	if err != nil {
		return 0, err
	}

	return convert.Int(rep)
}

func (c conn) LIndex(k, i interface{}) ([]byte, error) {
	rep, err := c.connection.Do("LINDEX", k, i)
	if err != nil {
		return nil, err
	}

	return toByteArray(rep)
}

func (c conn) LInsertAfter(k, pivot, v interface{}) (int, error) {
	rep, err := c.connection.Do("LINSERT", k, "AFTER", pivot, v)
	if err != nil {
		return 0, err
	}

	return convert.Int(rep)
}

func (c conn) LInsertBefore(k, pivot, v interface{}) (int, error) {
	rep, err := c.connection.Do("LINSERT", k, "BEFORE", pivot, v)
	if err != nil {
		return 0, err
	}

	return convert.Int(rep)
}

func (c conn) LLen(k interface{}) (int, error) {
	rep, err := c.connection.Do("LLEN", k)
	if err != nil {
		return 0, err
	}

	return convert.Int(rep)
}

func (c conn) LRange(k, start, end interface{}) ([][]byte, error) {
	rep, err := c.connection.Do("LRANGE", k, start, end)

	if err != nil {
		return nil, err
	}

	return toByteArrayArray(rep)
}

func (c conn) LRem(k, count, v interface{}) (int, error) {
	rep, err := c.connection.Do("LREM", k, count, v)
	if err != nil {
		return 0, err
	}

	return convert.Int(rep)
}

func (c conn) LSet(k, i, v interface{}) error {
	_, err := c.connection.Do("LSET", k, i, v)
	return err
}

func (c conn) LTrim(k, start, end interface{}) error {
	_, err := c.connection.Do("LTRIM", k, start, end)
	return err
}
