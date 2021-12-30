package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	REDIS_RANK string = "RANK" //
)

func main() {

	//for i := 0; i < 100; i += 1 {
	//
	//	fmt.Printf("%d \n", get_redis_value(1))
	//}
	r := gin.Default()
	r.POST("/user/:userID/:score", handle_update_score)
	r.GET("/user/:userID", handle_get_ranks)
	r.Run(":8080")

}

func handle_update_score(c *gin.Context) {
	score, _ := strconv.Atoi(c.Param("score"))
	userID, _ := strconv.Atoi(c.Param("userID"))
	fmt.Printf("%d update score %d \n", userID, score)
	redis_value := get_redis_value(score)
	write_to_redis(userID, redis_value)
	c.JSON(200, gin.H{"score": score, "redis_value": redis_value, "rank": get_my_rank(userID)})
}

type user struct {
	ID    interface{} `json:"userID"`
	Score int         `json:"score"`
	Rank  int         `json:"rank"` //从0开始
}

func handle_get_ranks(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))
	fmt.Printf("%d get ranks \n", userID)
	my_rank := get_my_rank(userID)
	limit := 10
	f := my_rank - int64(limit)
	if f < 0 {
		f = 0
	}
	e := my_rank + int64(limit)
	redis_values := GetRedis().ZRevRangeWithScores(REDIS_RANK, f, e).Val()
	ret := make([]*user, 0, len(redis_values))
	for i, v := range redis_values {
		u := &user{
			ID:    v.Member,
			Score: int(int64(v.Score) >> 48),
			Rank:  i,
		}
		ret = append(ret, u)
	}
	c.JSON(200, ret)
}

func write_to_redis(userID int, redis_value int64) {
	z := &redis.Z{
		Score:  float64(redis_value),
		Member: userID,
	}
	GetRedis().ZAdd(REDIS_RANK, z)
}

func get_my_rank(userID int) int64 {
	zRank := GetRedis().ZRevRank(REDIS_RANK, fmt.Sprintf("%d", userID))
	if zRank.Err() == nil {
		return zRank.Val()
	} else {
		return GetRedis().ZCard(REDIS_RANK).Val()
	}
}

var rank_index int64 = 12
var last_ms int64 = 0

func get_redis_value(_score int) int64 {
	score := int64(_score)<<48 + calc_rank_index()
	return score
}
func calc_rank_index() int64 {
	ms := time.Now().UnixNano() / int64(time.Millisecond)
	ms &= 0xFFFFFFFFFFF //44
	if last_ms != ms {
		atomic.AddInt64(&rank_index, -rank_index)
		last_ms = ms
	}
	new_rank_index := atomic.AddInt64(&rank_index, 1)
	ret := 1<<48 - (ms<<4 | new_rank_index%15)
	return ret
}
