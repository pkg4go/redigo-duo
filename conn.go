package duo

import r "github.com/garyburd/redigo/redis"

type Conn interface {
	// basic
	Del(l interface{}) error
	Set(k, v interface{}) error
	Get(k interface{}) ([]byte, error)
	Expire(k, expiration interface{}) error
	Exists(k interface{}) (bool, error)
}

type conn struct {
	connection r.Conn
}

func Duo(c r.Conn) Conn {
	return conn{connection: c}
}
