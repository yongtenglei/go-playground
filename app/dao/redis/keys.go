package redis

const (
	KeyPrefix          = "app:"        // Project name
	KeyPostTime        = "post:time"   // zset 帖子与时间排序
	KeyPostScore       = "post:score"  // zset 帖子与分数
	KeyPostVotedPrefix = "post:voted:" // zset 帖子与投票状态前缀,需要与post_id 拼接
	KeyCommunityPrefix = "community:" // set 社区前缀， 需要与community_id 拼接
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
