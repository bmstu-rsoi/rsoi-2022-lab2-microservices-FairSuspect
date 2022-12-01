package main

import (
	"privileges/controllers"
	"privileges/objects"
	"privileges/utils"

	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDBConnection(cnf utils.DBConfiguration) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cnf.Host, cnf.User, cnf.Password, cnf.Name, cnf.Port)
	db, e := gorm.Open(cnf.Type, dsn)

	if e != nil {
		utils.Logger.Print("DB Connection failed")
		utils.Logger.Print(e)
		panic("DB Connection failed")
	} else {
		utils.Logger.Print("DB Connection Established")
	}

	db.SingularTable(true)
	db.AutoMigrate(&objects.Privilege{}, &objects.PrivilegeHistory{})

	privilege := &objects.Privilege{
		Id:       1,
		Username: "Test Max",
		Status:   "BRONZE",
		Balance:  0,
	}
	db.FirstOrCreate(privilege)

	return db
}

func main() {
	rand.Seed(time.Now().UnixNano())

	utils.InitConfig()
	utils.InitLogger()
	defer utils.CloseLogger()

	db := initDBConnection(utils.Config.DB)
	defer db.Close()
	r := controllers.InitRouter(db)

	fmt.Printf("App started: http://localhost:%d\n", utils.Config.Port)
	controllers.RunRouter(r, utils.Config.Port)

}
