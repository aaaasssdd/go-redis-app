package idgenerator

import (
	"github.com/go-redis/redis"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCache(t *testing.T) {
	Convey("check counter module", t, func() {
		Convey("implement Client interface", func() {

			redisClient := redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			})
			redisIdGenerator := &RedisIdGenerator{redisClient}
			idGeneratorWrapper := IdGeneratorWrapper{IdGenerator: redisIdGenerator}

			Convey("init id = 0", func() {
				_, err := idGeneratorWrapper.Reset("id")
				So(err, ShouldBeNil)
			})

			Convey("incr id 1", func() {
				n, err := idGeneratorWrapper.Incr("id")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 1)
			})

			Convey("incr id 2", func() {
				n, err := idGeneratorWrapper.Incr("id")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 2)
			})

			Convey("incr id 3", func() {
				n, err := idGeneratorWrapper.Incr("id")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 3)
			})

			Convey("reset id to 0", func() {
				n, err := idGeneratorWrapper.Reset("id")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 3)
			})

			Convey("incr id by 10", func() {
				n, err := idGeneratorWrapper.Incr("id", 10)
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 10)
			})

			Convey("incr id by 99", func() {
				n, err := idGeneratorWrapper.Incr("id", 99)
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 109)
			})

			Convey("get id ", func() {
				n, err := idGeneratorWrapper.Get("id")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 109)
			})
		})
	})
}
