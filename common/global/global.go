package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	GinEngine *gin.Engine
	DBMysql   *gorm.DB
)
