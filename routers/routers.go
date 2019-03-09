package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

var RoutineCard = new(controllers.RoutineCard)

func init() {
	Routers = gin.Default()
}

func SetupRouters() {
	go RoutineCard.UpdateDataOnDB()
	card := new(controllers.Card)
	Routers.GET("/b/cards/review", card.AllCardReview)
	Routers.GET("/b/cards/change-due", card.AllCardChangeDueDate)
}
