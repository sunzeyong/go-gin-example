package v1

import (
	"net/http"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/alarm"
	"github.com/gin-gonic/gin"
)

func GetNewAlarms(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	appG := app.Gin{C: c}

	alarmService := &alarm.AlarmService{}
	article, err := alarmService.GetNewAlarms()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

func GetAlarmTypeStat(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	appG := app.Gin{C: c}

	alarmService := &alarm.AlarmService{}
	article, err := alarmService.GetAlarmTypeStat()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

func GetCompanyStat(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	appG := app.Gin{C: c}

	alarmService := &alarm.AlarmService{}
	article, err := alarmService.GetCompanyStat()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

func GetCaptainStat(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	appG := app.Gin{C: c}

	alarmService := &alarm.AlarmService{}
	article, err := alarmService.GetCaptainStat()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}
