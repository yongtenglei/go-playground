package mysql

import (
	"app/models"
	"database/sql"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`

	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("No community")
			err = nil
			return
		}
		zap.L().Error("No community", zap.Error(err))
		return
	}
	return
}

func GetCommunityDetail(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`
	err = db.Get(communityDetail, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("No such a community")
			err = nil
			return
		}
		zap.L().Error("GetCommunityDetail failed", zap.Error(err))
		return
	}
	return
}
