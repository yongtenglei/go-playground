package redis

import (
	"app/models"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteImeExpire = errors.New("Vote time beyond to one week")
	ErrRepeatVote    = errors.New("Have been Voted")
)

func CreatePost(postID int64, communityID int64) (err error) {
	// redis 事务
	pipeline := rdb.TxPipeline()

	// 创建 time zset
	pipeline.ZAdd(getRedisKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 创建 score zset
	pipeline.ZAdd(getRedisKey(KeyPostScore), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 创建 community set
	cKey := getRedisKey(KeyCommunityPrefix + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)

	_, err = pipeline.Exec()

	return
}

func VoteForPost(uid string, postID string, v float64) (err error) {
	// 判断投票限制
	postTime := rdb.ZScore(getRedisKey(KeyPostTime), postID).Val()

	// testblock
	//postTime1 := rdb.ZScore(getRedisKey(KeyPostTime), postID).Val()
	//postScore1 := rdb.ZScore(getRedisKey(KeyPostScore), postID).Val()
	//fmt.Println(getRedisKey(KeyPostTime))
	//fmt.Println("postID", postID)
	//fmt.Println("postTime: ", postTime1)
	//fmt.Println("postScore", postScore1)
	//fmt.Println("now: ", time.Now().Unix())
	//fmt.Println("diffTime", math.Abs(float64(time.Now().Unix()))-postTime1)
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteImeExpire
	}

	// 更新帖子分数
	ov := rdb.ZScore(getRedisKey(KeyPostVotedPrefix+postID), uid).Val()
	var op float64
	// 不允许重复投票
	if ov == v {
		return ErrRepeatVote
	}
	if v > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - v)
	fmt.Println("==============")
	fmt.Println("diff ", diff)
	fmt.Println("==============")

	//redis 事务开始
	pipline := rdb.TxPipeline()

	pipline.ZIncrBy(getRedisKey(KeyPostScore), op*diff*scorePerVote, postID)

	// 保存投票信息
	if v == 0 {
		pipline.ZRem(getRedisKey(KeyPostVotedPrefix+postID), uid)
	} else {
		pipline.ZAdd(getRedisKey(KeyPostVotedPrefix+postID), redis.Z{
			Score:  v,
			Member: uid,
		})
	}

	_, err = pipline.Exec()

	// testblock
	//fmt.Println("==================")
	//postTime2 := rdb.ZScore(getRedisKey(KeyPostTime), postID).Val()
	//postScore2 := rdb.ZScore(getRedisKey(KeyPostScore), postID).Val()
	//fmt.Println(getRedisKey(KeyPostTime))
	//fmt.Println("postID", postID)
	//fmt.Println("postTime: ", postTime2)
	//fmt.Println("postScore", postScore2)
	//fmt.Println("now: ", time.Now().Unix())
	//fmt.Println("diffTime", math.Abs(float64(time.Now().Unix()))-postTime2)
	//fmt.Println("===================")
	//fmt.Println("diffScore", postScore1-postScore2)

	return
}

func GetPostIDsInOrder(p *models.ParamPostListInfo) ([]string, error) {
	//确定redis key
	var key string
	switch p.Order {
	case models.PostTimeOrder:
		key = getRedisKey(KeyPostTime)
	case models.PostScoreOrder:
		key = getRedisKey(KeyPostScore)
	}

	// 确定差寻的起始和终点
	start := (p.Page - 1) * p.NumPerPage
	end := start + p.NumPerPage - 1

	// 分数由大到小查询
	return rdb.ZRevRange(key, int64(start), int64(end)).Result()

}

func GetThumbUPs(idList []string) (thumbups []int64, err error) {
	// Works but not diligent
	//thumbups = make([]int64, 0, len(idList))
	//for _, id := range idList {
	//key := getRedisKey(KeyPostVotedPrefix + id)
	//v := rdb.ZCount(key, "1", "1").Val()
	//thumbups = append(thumbups, v)
	//}

	// 一次链接发送多条请求, 减少RTT
	thumbups = make([]int64, 0, len(idList))
	pipline := rdb.Pipeline()
	for _, id := range idList {
		key := getRedisKey(KeyPostVotedPrefix + id)
		pipline.ZCount(key, "1", "1")
	}

	cmders, err := pipline.Exec()
	if err != nil {
		return nil, err
	}

	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		thumbups = append(thumbups, v)
	}

	return
}

func GetCommunityPostIDsInOrder(p *models.ParamPostListInfo) ([]string, error) {
	// orderKey
	var orderKey string
	switch p.Order {
	case models.PostTimeOrder:
		orderKey = getRedisKey(KeyPostTime)
	case models.PostScoreOrder:
		orderKey = getRedisKey(KeyPostScore)
	}
	//使用zinterstore 把分区帖子set与帖子分数zset(time or score)生成一个新的zset

	cKey := getRedisKey(KeyCommunityPrefix + strconv.Itoa(int(p.CommunityID)))

	// 利用缓存key减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(key).Val() < 1 {
		// 不存在，重新构建zinterstore
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(orderKey, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		if _, err := pipeline.Exec(); err != nil {
			return nil, err
		}
	}

	// 确定差寻的起始和终点
	start := (p.Page - 1) * p.NumPerPage
	end := start + p.NumPerPage - 1

	// 分数由大到小查询
	return rdb.ZRevRange(orderKey, int64(start), int64(end)).Result()
}
