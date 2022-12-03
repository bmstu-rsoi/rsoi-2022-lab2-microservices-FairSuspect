package main

import (
	"gateway/controllers"
	"gateway/utils"

	"fmt"
	"math/rand"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	utils.InitConfig()
	utils.InitLogger()
	defer utils.CloseLogger()

	r := controllers.InitRouter()
	fmt.Printf("App started: http://localhost:%d\n", utils.Config.Port)
	controllers.RunRouter(r, utils.Config.Port)

}
