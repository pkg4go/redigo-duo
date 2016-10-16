
[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

### redigo-duo

```go
package main

import r "github.com/garyburd/redigo/redis"
import "github.com/pkg4go/redigo-duo"
import "log"

var redis duo.Conn

func init() {
	con, err := r.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	redis = duo.Duo(con)
}

func main() {

	_ = redis.Set("a", "1")

	v, _ := redis.Get("a")
	log.Println(string(v[:]))

	_ = redis.Del("a")
	exists, _ := redis.Exists("a")
	log.Println(exists)
}
```

### License
MIT

[travis-img]: https://img.shields.io/travis/pkg4go/redigo-duo.svg?style=flat-square
[travis-url]: https://travis-ci.org/pkg4go/redigo-duo
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: http://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/pkg4go/redigo-duo
