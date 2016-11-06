package duo

import r "github.com/garyburd/redigo/redis"

type Conn interface {
	// basic
	Del(k interface{}) error
	Set(k, v interface{}) error
	Get(k interface{}) ([]byte, error)
	GetSet(k, v interface{}) ([]byte, error)
	Expire(k, expiration interface{}) error
	Exists(k interface{}) (bool, error)
	Keys(pattern string) ([]string, error)
	FlushAll() error
	FlushDb() error
	// list
	RPop(k interface{}) ([]byte, error)
	RPush(k interface{}, vals ...interface{}) (int, error)
	RPushX(k, v interface{}) (int, error)

	LPop(k interface{}) ([]byte, error)
	LPush(k interface{}, vals ...interface{}) (int, error)
	LPushX(k, v interface{}) (int, error)

	LIndex(k, i interface{}) ([]byte, error)
	LInsertAfter(k, pivot, v interface{}) (int, error)
	LInsertBefore(k, pivot, v interface{}) (int, error)
	LLen(k interface{}) (int, error)
	LRange(k, start, end interface{}) ([][]byte, error)
	LRem(k, c, v interface{}) (int, error)
	LSet(k, i, v interface{}) error
	LTrim(k, start, end interface{}) error
}

type conn struct {
	connection r.Conn
}

func Duo(c r.Conn) Conn {
	return conn{connection: c}
}
