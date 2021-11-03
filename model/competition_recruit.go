package model

import (
	"gorm.io/gorm"
	"time"
)

type Competition_recruit struct {
	gorm.Model
	Title						string
	Contact_qq					string
	Contact_wechat				string
	Contact_tel 				string
	Deadline					time.Time
	Competition_introduction	string
	Demanding_num				int64
	Now_num						int64
	Demanding_introduction		string
	Master_name					string
	Master_sex					string
	Master_gradeandmajor		string
	Master_introduction			string
	Teammate_introduction		string
	Contact_notnull				int64//联系方式至少选一项触发器
}

type ComperitionEntails struct {//帖子详情页
	Created_at					time.Time
	Title						string
	Contact_qq					string
	Contact_wechat				string
	Contact_tel 				string
	Deadline 					time.Time
	Competition_introduction	string
	Demanding_num				int64
	Now_num						int64
	Demanding_introduction		string
	Master_name					string
	Master_sex					string
	Master_gradeandmajor		string
	Master_introduction			string
	Teammate_introduction		string
}