package duo

import r "github.com/garyburd/redigo/redis"
import c "github.com/pkg4go/convert"
import . "github.com/pkg4go/assert"
import "github.com/pkg4go/slicex"
import "github.com/google/uuid"
import "testing"
import "time"

var redis Conn

func init() {
	con, err := r.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	redis = Duo(con)
}

func TestSetGetDel(t *testing.T) {
	a := A{t}

	k := uuid.New().String()

	e := redis.Set(k, "1")
	a.Nil(e)

	v, e := redis.Get(k)
	a.Nil(e)
	a.Equal(c.String(v), "1")

	e = redis.Set(k, "2")
	a.Nil(e)

	v, e = redis.Get(k)
	a.Nil(e)
	a.Equal(c.String(v), "2")
	// del
	e = redis.Del(k)
	a.Nil(e)

	v, e = redis.Get(k)
	a.Nil(e)
	a.Equal(len(v), 0)
}

func TestGetSet(t *testing.T) {
	a := A{t}

	k := uuid.New().String()

	v, e := redis.GetSet(k, "ooo")
	a.Nil(e)
	a.Equal(string(v[:]), "")
	v, e = redis.GetSet(k, "kkk")
	a.Nil(e)
	a.Equal(string(v[:]), "ooo")
	v, e = redis.GetSet(k, "_")
	a.Nil(e)
	a.Equal(string(v[:]), "kkk")
}

func TestExpireExistsKeys(t *testing.T) {
	a := A{t}

	k := uuid.New().String()

	e := redis.Set(k, "2")
	a.Nil(e)
	e = redis.Expire(k, "1")
	a.Nil(e)
	v, e := redis.Get(k)
	a.Nil(e)
	a.Equal(c.String(v), "2")
	// keys
	keys, e := redis.Keys("*")
	a.Nil(e)
	a.Equal(slicex.Contains(keys, k), true)
	a.Equal(slicex.Contains(keys, k+"not exists"), false)
	// exists
	exists, e := redis.Exists(k)
	a.Nil(e)
	a.Equal(exists, true)

	time.Sleep(1200 * time.Millisecond)

	exists, e = redis.Exists(k)
	a.Nil(e)
	a.Equal(exists, false)

	v, e = redis.Get(k)
	a.Nil(e)
	a.Equal(len(v), 0)
}
