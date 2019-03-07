package routers

import (
	"../controllers"
	"../urls"
	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

var RoutineCard = new(controllers.RoutineCard)

func init() {
	Routers = gin.Default()
}

func SetupRouters() {
	url_pattern := urls.ReturnURLS()
	go RoutineCard.UpdateDataOnDB()
	card := new(controllers.Card)
	Routers.GET(url_pattern.CARD_REVIEW_PATH, card.AllCardReview)
	Routers.GET(url_pattern.CARD_CHANGE_DUE_PATH, card.AllCardChangeDueDate)
}
