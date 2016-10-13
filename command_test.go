package duo

import r "github.com/garyburd/redigo/redis"
import c "github.com/pkg4go/convert"
import . "github.com/pkg4go/assert"
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

func TestSetGet(t *testing.T) {
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
}

func TestExpire(t *testing.T) {
	a := A{t}

	k := uuid.New().String()

	e := redis.Set(k, "2")
	a.Nil(e)
	e = redis.Expire(k, "1")
	a.Nil(e)
	v, e := redis.Get(k)
	a.Nil(e)
	a.Equal(c.String(v), "2")

	time.Sleep(1 * time.Second)

	v, e = redis.Get(k)
	a.Nil(e)
	a.Equal(len(v), 0)
}
