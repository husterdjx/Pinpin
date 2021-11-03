package api

import (
	"Pinpin/http_param"
	"Pinpin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Competitionhandler(ctx *gin.Context) {
	params := http_param.Competition{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	if err := service.CompetitionCreateService(ctx,params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}