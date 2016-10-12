package duo

import r "github.com/garyburd/redigo/redis"

type Conn interface {
	Set(k, v interface{}) (res []byte, err error)
	Get(k interface{}) (res []byte, err error)
}

type conn struct {
	connection r.Conn
}

func (c conn) Set(k, v interface{}) ([]byte, error) {
	rep, err := c.connection.Do("SET", k, v)
	if err != nil {
		return nil, err
	}

	return toByteArray(rep)
}

func (c conn) Get(k interface{}) ([]byte, error) {
	rep, err := c.connection.Do("GET", k)
	if err != nil {
		return nil, err
	}

	return toByteArray(rep)
}

func Duo(c r.Conn) Conn {
	return conn{connection: c}
}
