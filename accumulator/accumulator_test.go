package accumulator

import (
	"github.com/go-redis/redis"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yacen/go-redis-app/idgenerator"
	"testing"
)

func TestCache(t *testing.T) {
	Convey("check accumulator module", t, func() {
		Convey("implement Client interface", func() {

			redisClient := redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			})

			redisIdGenerator := idgenerator.RedisIdGenerator{redisClient}
			idGeneratorWrapper := idgenerator.IdGeneratorWrapper{IdGenerator: redisIdGenerator}

			redisAcc := RedisAccumulator{redisIdGenerator}
			acc := AccumulatorWrapper{idGeneratorWrapper, redisAcc}

			Convey("init number = 0", func() {
				_, err := acc.Reset("number")
				So(err, ShouldBeNil)
			})

			Convey("incr number 1", func() {
				n, err := acc.Incr("number")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 1)
			})

			Convey("incr number 2", func() {
				n, err := acc.Incr("number")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 2)
			})

			Convey("incr number 3", func() {
				n, err := acc.Incr("number")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 3)
			})

			Convey("reset number to 0", func() {
				n, err := acc.Reset("number")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 3)
			})

			Convey("incr number by 10", func() {
				n, err := acc.Incr("number", 10)
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 10)
			})

			Convey("incr number by 99", func() {
				n, err := acc.Incr("number", 99)
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 109)
			})

			Convey("get number ", func() {
				n, err := acc.Get("number")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 109)
			})

			Convey("decr number", func() {
				n, err := acc.Decr("number")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 108)
			})

			Convey("decr number by 88", func() {
				n, err := acc.Decr("number", 88)
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 20)
			})

			Convey("reset number to 0 and get old number 20", func() {
				n, err := acc.Reset("number")
				So(err, ShouldBeNil)
				So(n, ShouldEqual, 20)
			})
		})
	})
}
