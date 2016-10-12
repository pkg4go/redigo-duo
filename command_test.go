package duo

import r "github.com/garyburd/redigo/redis"
import c "github.com/pkg4go/convert"
import . "github.com/pkg4go/assert"
import "testing"

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

	_, e := redis.Set("a", "1")
	a.Nil(e)

	v, e := redis.Get("a")
	a.Nil(e)
	a.Equal(c.String(v), "1")

	_, e = redis.Set("a", "2")
	a.Nil(e)

	v, e = redis.Get("a")
	a.Nil(e)
	a.Equal(c.String(v), "2")
}
