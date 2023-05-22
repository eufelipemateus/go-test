package main

import (
	"fmt"

	config "github.com/eufelipemateus/test-example/configs"
	"github.com/eufelipemateus/test-example/routes"
	"github.com/gin-gonic/gin"

	"github.com/eufelipemateus/test-example/database"

)

func main() {

	err := config.Load()
	if err != nil {
		panic("Erro on load config.toml")
	}
	database.OpenConnection()

	if config.GetApp().GenerateDB {
		database.GenerateDB()
	}

	fmt.Printf("%s \n Runing in %s ", config.GetApp().Name, config.GetApp().Port)
	
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	
	r.LoadHTMLGlob(config.GetApp().Views)
	
	r.SetTrustedProxies(nil)
	routes.Setup(r)
	r.Run(config.GetApp().Port)
}
