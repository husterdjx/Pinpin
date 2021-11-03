package service

import (
	"Pinpin/dao"
	"Pinpin/model"
)

func GetCompetitionDetailsService() (res []model.ComperitionEntails, err error) {
	res, err = dao.GetAllCompetitionsDetails()
	if err != nil {
		return
	}
	if len(res) == 0 {
		return make([]model.ComperitionEntails, 0), nil
	}
	return
}