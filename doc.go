package duo

// package main

// import r "github.com/garyburd/redigo/redis"
// import "github.com/pkg4go/redigo-duo"
// import "log"

// var redis duo.Conn

// func init() {
// 	con, err := r.Dial("tcp", ":6379")
// 	if err != nil {
// 		panic(err)
// 	}

// 	redis = duo.Duo(con)
// }

// func main() {

// 	_ = redis.Set("a", "1")

// 	v, _ := redis.Get("a")
// 	log.Println(string(v[:]))

// 	_ = redis.Del("a")
// 	exists, _ := redis.Exists("a")
// 	log.Println(exists)
// }
