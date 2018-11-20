package cache

import (
	"github.com/go-redis/redis"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCache(t *testing.T) {
	Convey("check cache module", t, func() {
		Convey("implement Client interface", func() {
			redisClient := redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			})
			redisCacher := &RedisCacher{redisClient}
			cache := CacherWrapper{Cacher: redisCacher}
			Convey("Set site = github.com ", func() {
				err := cache.Set("sit", "github.com", 0)
				So(err, ShouldBeNil)
			})

			Convey("get site should be github.com", func() {
				value, err := cache.Get("sit")
				So(err, ShouldBeNil)
				So(value, ShouldEqual, "github.com")
			})

			Convey("Set size = 10000 ", func() {
				err := cache.Set("size", 10000, 0)
				So(err, ShouldBeNil)
			})

			Convey("get size should be 10000", func() {
				value, err := cache.Get("size")
				So(err, ShouldBeNil)
				So(value, ShouldEqual, "10000")
			})

			Convey(`Set json = {"type": "boolean","description": "是否为热门"} `, func() {
				err := cache.Set("json", "{\"type\": \"boolean\",\"description\": \"是否为热门\"}", 0)
				So(err, ShouldBeNil)
			})

			Convey(`get json should be {"type": "boolean","description": "是否为热门"} `, func() {
				value, err := cache.Get("json")
				So(err, ShouldBeNil)
				So(value, ShouldEqual, "{\"type\": \"boolean\",\"description\": \"是否为热门\"}")
			})
		})
	})
}
