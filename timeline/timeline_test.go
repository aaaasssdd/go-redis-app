package timeline

import (
	"fmt"
	"github.com/go-redis/redis"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCache(t *testing.T) {
	Convey("check timeline module", t, func() {
		Convey("implement timeline interface", func() {
			Redis := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
			redisTimeline := RedisTimeline{Redis: Redis}

			timeline := TimeLineWrapper{&redisTimeline}

			Convey("delete jax timeline", func() {
				err := Redis.Del("timeline:username:jax:message").Err()
				So(err, ShouldBeNil)

				Convey("one day, jax post a message", func() {
					err := timeline.Add("timeline:username:jax:message", "hello everyone")
					So(err, ShouldBeNil)

					Convey("other day, jax post a message again", func() {
						err := timeline.Add("timeline:username:jax:message", "today is sunny")
						So(err, ShouldBeNil)

						Convey("get timeline", func() {
							data, err := timeline.FetchRecent("timeline:username:jax:message", 7)
							So(err, ShouldBeNil)
							So(data[0], ShouldEqual, "today is sunny")
							So(data[1], ShouldEqual, "hello everyone")

							Convey("jax post message a month", func() {
								for i := 0; i < 30; i++ {
									err := timeline.Add("timeline:username:jax:message", fmt.Sprintf("post %dth message", i))
									So(err, ShouldBeNil)
								}

								Convey("get timeline a month", func() {
									data, err := timeline.FetchFromIndex("timeline:username:jax:message", 0, 30)
									So(err, ShouldBeNil)
									for i := 0; i < 30; i++ {
										So(fmt.Sprintf("post %dth message", 29-i), ShouldEqual, data[i])
									}
								})
							})
						})
					})
				})
			})

		})
	})
}
