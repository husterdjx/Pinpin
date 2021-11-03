package global

import (
	"Pinpin/config"
	"gorm.io/gorm"
)
var (
	Config      *config.Config
	DB          *gorm.DB
)