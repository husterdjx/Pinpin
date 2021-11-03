package http_param

import (
	"Pinpin/util"
	"time"
)

type Competition struct {
	Title						string		`form:"title" binding:"required"`
	Contact_qq					string		`form:"contact_qq" binding:"omitempty"`
	Contact_wechat				string		`form:"contact_wechat" binding:"omitempty"`
	Contact_tel 				string		`form:"contact_tel" binding:"omitempty"`
	Deadline					time.Time	`form:"deadline" binding:"required"`
	Competition_introduction	string		`form:"competition_introduction" binding:"omitempty"`
	Demanding_num				int64		`form:"demanding_num" binding:"required"`
	Now_num						int64		`form:"now_num" binding:"required"`
	Demanding_introduction		string		`form:"demanding_introduction" binding:"required,max=200"`
	Master_name					string		`form:"master_name" binding:"omitempty"`
	Master_sex					string		`form:"master_sex" binding:"omitempty"`
	Master_gradeandmajor		string		`form:"master_gradeandmajor" binding:"omitempty"`
	Master_introduction			string		`form:"master_introduction" binding:"omitempty,max=100"`
	Teammate_introduction		string		`form:"teammate_introduction" binding:"omitempty,max=100"`
	Contact_notnull				int64
}

func(r* Competition)GetError(err error) string {
	m := map[string]string {
		"Title":					"比赛名称",
		"Deadline":					"截止日期",
		"Demanding_num":			"总需求人数",
		"Now_num":					"已有人数",
		"Demanding_introduction":	"需求介绍",
		"Master_sex":				"帖主性别",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}