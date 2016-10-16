package duo

func (c conn) Del(k interface{}) error {
	_, err := c.connection.Do("DEL", k)
	return err
}

func (c conn) Set(k, v interface{}) error {
	_, err := c.connection.Do("SET", k, v)
	return err
}

func (c conn) Get(k interface{}) ([]byte, error) {
	rep, err := c.connection.Do("GET", k)
	if err != nil {
		return nil, err
	}

	return toByteArray(rep)
}

func (c conn) Expire(k, expiration interface{}) error {
	_, err := c.connection.Do("EXPIRE", k, expiration)
	return err
}

func (c conn) Exists(k interface{}) (bool, error) {
	rep, err := c.connection.Do("EXISTS", k)
	if err != nil {
		return false, err
	}

	i, err := toInt64(rep)
	if err != nil {
		return false, err
	}

	return i == 1, nil
}
