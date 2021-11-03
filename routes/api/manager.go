package api

import (
	"Pinpin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCompetitionDetailsHandler(ctx *gin.Context) {
	res, err := service.GetCompetitionDetailsService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res,
		})
		return
	}
}

