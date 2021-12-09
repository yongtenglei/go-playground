package logic

import (
	"app/dao/mysql"
	"app/dao/redis"
	"app/models"
	"app/pkg/snowflake"
	"strconv"

	"go.uber.org/zap"
)

func CreatePost(post *models.Post) error {
	// 生成 Post Id
	var id int64 = snowflake.GenID()
	post.ID = id

	// 保存到数据库中
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("logic: Mysql Create Post failed", zap.Error(err))
		return err
	}

	if err := redis.CreatePost(post.ID, post.CommunityID); err != nil {
		zap.L().Error("logic: Redis Create Post failed", zap.Error(err))
		return err
	}

	return nil
}

func PostDetile(pid int64) (postDetail *models.APIPostDetail, err error) {
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic: GetPostDetile failed", zap.Int64("pid", pid), zap.Error(err))
		return nil, err
	}

	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("logic: GetPostDetile failed", zap.Int64("uid", post.AuthorID), zap.Error(err))
		return nil, err
	}

	community, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("logic: GetPostDetile failed", zap.Int64("cid", post.CommunityID), zap.Error(err))
		return nil, err
	}

	postDetail = &models.APIPostDetail{
		AuthorName:      user.UserName,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func PostList(page int64, numPerPage int64) (postList []*models.APIPostDetail, err error) {
	// 获取数据
	posts, err := mysql.GetPostList(page, numPerPage)
	if err != nil {
		zap.L().Error("logic: GetPostList failed", zap.Error(err))
		return nil, err
	}

	postList = make([]*models.APIPostDetail, 0, len(posts))
	// 拼接APIPostDetail
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("logic: GetPostDetile failed", zap.Int64("uid", post.AuthorID), zap.Error(err))
			continue
		}

		community, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("logic: GetPostDetile failed", zap.Int64("cid", post.CommunityID), zap.Error(err))
			continue
		}

		postDetile := &models.APIPostDetail{
			AuthorName:      user.UserName,
			Post:            post,
			CommunityDetail: community,
		}

		postList = append(postList, postDetile)
	}

	return
}

// VoteForPost 使用简单的投票分数算法
// 投一票加432分 (一天86400s需要投票200票维持一天, 即 86400 / 200)
// 逻辑:
// 1. when direction = 1 or -1:
// * when original direction == 0 -> original direction = 1 / -1
// * when original direction == -1 / 1 -> original direction = 1 / -1
// 2. when direction = 0:
// * when original direction == 1, direction = 0
// * when original direction == -1, direction = 0
// 可简化为:
// when direction > original direction 加分
// when direction < original direction 减分
// 为减轻存储压力做出以下限制:
// 帖子发表一星期内可以进行表态,一星期后不能进行表态
// 即到期后
// 1. 将redis中的赞成反对态度存储到mysql中
// 2. 删除redis中的KeyPostVotedPrefix键
func VoteForPost(uid int64, voteData *models.ParamVoteData) error {
	if err := redis.VoteForPost(strconv.FormatInt(uid, 10), strconv.FormatInt(voteData.PostID, 10), float64(voteData.Direction)); err != nil {
		zap.L().Error("logic: VoteForPost failed", zap.Error(err))
		return err
	}
	return nil
}

func PostList2(p *models.ParamPostListInfo) (postDetailList []*models.APIPostDetail, err error) {
	// redis 中获取id列表
	idList, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("logic: GetPostIDsInOrder failed", zap.Error(err))
		return
	}

	if len(idList) == 0 {
		zap.L().Warn("logic: GetPostIDsInOrder successfully, but maybe empty")
		return
	}

	// 获取 thumbup 数量
	thumbups := make([]int64, 0, len(idList))
	thumbups, err = redis.GetThumbUPs(idList)

	// Mysql中获取详细信息
	postList, err := mysql.GetPostList2(idList)
	if err != nil {
		zap.L().Error("logic: GetPostList2 by order failed", zap.Error(err))
	}

	postDetailList = make([]*models.APIPostDetail, 0, len(postList))
	// 拼接APIPostDetail
	for idx, post := range postList {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("logic: GetPostDetile failed", zap.Int64("uid", post.AuthorID), zap.Error(err))
			continue
		}

		community, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("logic: GetPostDetile failed", zap.Int64("cid", post.CommunityID), zap.Error(err))
			continue
		}

		postDetile := &models.APIPostDetail{
			AuthorName:      user.UserName,
			ThumbupNum:      thumbups[idx],
			Post:            post,
			CommunityDetail: community,
		}

		postDetailList = append(postDetailList, postDetile)
	}
	return
}

func PostCommunityList(p *models.ParamPostListInfo) (postDetailList []*models.APIPostDetail, err error) {
	// redis 中获取id列表
	idList, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("logic: GetPostIDsInOrder failed", zap.Error(err))
		return
	}

	if len(idList) == 0 {
		zap.L().Warn("logic: GetPostIDsInOrder successfully, but maybe empty")
		return
	}

	// 获取 thumbup 数量
	thumbups := make([]int64, 0, len(idList))
	thumbups, err = redis.GetThumbUPs(idList)

	// Mysql中获取详细信息
	postList, err := mysql.GetPostList2(idList)
	if err != nil {
		zap.L().Error("logic: GetPostList2 by order failed", zap.Error(err))
	}

	postDetailList = make([]*models.APIPostDetail, 0, len(postList))
	// 拼接APIPostDetail
	for idx, post := range postList {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("logic: GetPostDetile failed", zap.Int64("uid", post.AuthorID), zap.Error(err))
			continue
		}

		community, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("logic: GetPostDetile failed", zap.Int64("cid", post.CommunityID), zap.Error(err))
			continue
		}

		postDetile := &models.APIPostDetail{
			AuthorName:      user.UserName,
			ThumbupNum:      thumbups[idx],
			Post:            post,
			CommunityDetail: community,
		}

		postDetailList = append(postDetailList, postDetile)
	}
	return
}

func SwithQueryPostListMode(p *models.ParamPostListInfo) (postDetailList []*models.APIPostDetail, err error) {
	if p.CommunityID == 0 {
		postDetailList, err = PostList2(p)
	} else {
		postDetailList, err = PostCommunityList(p)
	}
	if err != nil {
		return nil, err
	}
	return
}
