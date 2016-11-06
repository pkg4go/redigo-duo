package duo

import . "github.com/pkg4go/assert"
import "github.com/google/uuid"
import "testing"

func TestList(t *testing.T) {
	a := A{t}

	a.Equal(1, 1)

	k := uuid.New().String()

	c, e := redis.RPushX(k, "not success")
	a.Nil(e)
	a.Equal(c, 0)

	c, e = redis.LPushX(k, "not success")
	a.Nil(e)
	a.Equal(c, 0)

	c, e = redis.RPush(k, 1, 2, 3)
	a.Nil(nil)
	a.Equal(c, 3)

	c, e = redis.RPushX(k, 4)
	a.Nil(nil)
	a.Equal(c, 4)

	c, e = redis.LPush(k, -1, -2, -3)
	a.Nil(nil)
	a.Equal(c, 7)

	c, e = redis.LPushX(k, -4)
	a.Nil(nil)
	a.Equal(c, 8)
	// pop
	v, e := redis.LPop(k)
	a.Nil(e)
	a.Equal(string(v[:]), "-4")

	v, e = redis.RPop(k)
	a.Nil(e)
	a.Equal(string(v[:]), "4")
	// insert
	c, e = redis.LInsertAfter(k, "not exist", "some")
	a.Nil(nil)
	a.Equal(c, -1)

	c, e = redis.LInsertBefore(k, "not exist", "some")
	a.Nil(nil)
	a.Equal(c, -1)

	c, e = redis.LInsertAfter(k, "1", "a1")
	a.Nil(nil)
	a.Equal(c, 7)

	c, e = redis.LInsertBefore(k, "-1", "b-1")
	a.Nil(nil)
	a.Equal(c, 8)
	// index
	v, e = redis.LIndex(k, 4)
	a.Nil(e)
	a.Equal(string(v[:]), "1")

	v, e = redis.LIndex(k, 999)
	a.Nil(e)
	a.Equal(len(v), 0)
	// len
	c, e = redis.LLen(k)
	a.Nil(e)
	a.Equal(c, 8)
	// range
	l, e := redis.LRange(k, 0, 2)
	a.Nil(e)
	a.Equal(string(l[0][:]), "-3")
	a.Equal(string(l[1][:]), "-2")
	a.Equal(string(l[2][:]), "b-1")
	// set
	e = redis.LSet(k, 0, "set")
	a.Nil(e)

	e = redis.LSet(k, 999, "error")
	a.Equal(e.Error(), "ERR index out of range")
	// trim
	e = redis.LTrim(k, 1, -1)
	a.Nil(e)
	c, e = redis.LLen(k)
	a.Nil(e)
	a.Equal(c, 7)
	// remove
	c, e = redis.LRem(k, 0, "-2")
	a.Nil(e)
	a.Equal(c, 1)
}
