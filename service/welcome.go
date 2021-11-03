package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/model"
	"context"
	"errors"
	"gorm.io/gorm"
)
func CompetitionCreateService(ctx context.Context,param http_param.Competition) (err error) {
	if param.Contact_qq == "" && param.Contact_wechat == "" && param.Contact_tel == "" { //联系方式三选一触发器
		param.Contact_notnull = 0
	} else {
		param.Contact_notnull = 1
	}
	competition := model.Competition_recruit{
		Model:						gorm.Model{},
		Title:						param.Title,
		Contact_qq:					param.Contact_qq,
		Contact_wechat:				param.Contact_wechat,
		Contact_tel: 				param.Contact_tel,
		Deadline:					param.Deadline,
		Competition_introduction:	param.Competition_introduction,
		Demanding_num:				param.Demanding_num,
		Now_num:					param.Now_num,
		Demanding_introduction:		param.Demanding_introduction,
		Master_name:				param.Master_name,
		Master_sex:					param.Master_sex,
		Master_gradeandmajor:		param.Master_gradeandmajor,
		Master_introduction:		param.Master_introduction,
		Teammate_introduction:		param.Teammate_introduction,
		Contact_notnull:			param.Contact_notnull,
	}
	err = dao.CompetitionCreate(competition)
	if err != nil {
		return errors.New("服务器发生错误")
	}
	return nil
}
