package routes

import (
	"github.com/eufelipemateus/test-example/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {


	app.GET("/", controllers.GetContact)
	app.POST("/", controllers.SsendMail)

}
