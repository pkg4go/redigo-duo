package duo

import r "github.com/garyburd/redigo/redis"

type Conn interface {
	Set(k, v interface{}) error
	Get(k interface{}) ([]byte, error)
	Expire(k, expiration interface{}) error
}

type conn struct {
	connection r.Conn
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

func Duo(c r.Conn) Conn {
	return conn{connection: c}
}
