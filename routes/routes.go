package routes

import (
	"Pinpin/routes/api"
	"github.com/gin-gonic/gin"
)

func UsePinpinRouter(r *gin.Engine) {
	Pinpinapi := r.Group("/api")
	{
		welcome := Pinpinapi.Group("/welcome")
		{
			welcome.POST("/competition",api.Competitionhandler)
		}
		manage := Pinpinapi.Group("/manage")
		{
			recruit := manage.Group("/recruit")
			{
				recruit.GET("/competition", api.GetCompetitionDetailsHandler)
			}
		}
	}
}
