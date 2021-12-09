package logic

import (
	"app/dao/mysql"
	"app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetail(id)

}
