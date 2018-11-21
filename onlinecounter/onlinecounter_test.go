package onlinecounter

import (
	"github.com/go-redis/redis"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestCache(t *testing.T) {
	Convey("check accumulator module", t, func() {
		Convey("implement Client interface", func() {
			redisClient := redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			})
			redisOnlineCounter := RedisOnlineCounter{redisClient}
			onlineCounter := OnlineCounterWrapper{redisOnlineCounter}

			Convey("Today is 2018-10-01", func() {
				timeKey := "2018-10-01"
				redisClient.Del(timeKey)
				Convey("generate 10000 online", func() {
					for i := 0; i < 10000; i++ {
						id := i*10000 + rand.Intn(10000)
						err := onlineCounter.Online(timeKey, int64(id))
						So(err, ShouldBeNil)
					}

					Convey("should have 10000 online", func() {
						num, err := onlineCounter.Count(timeKey)
						So(err, ShouldBeNil)
						So(num, ShouldEqual, 10000)
					})
				})

			})
		})
	})
}
