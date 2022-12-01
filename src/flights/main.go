package main

import (
	"flights/controllers"
	"flights/objects"
	"flights/utils"

	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDBConnection(cnf utils.DBConfig) *gorm.DB {
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
	db.AutoMigrate(&objects.Airport{})
	db.AutoMigrate(&objects.Flight{})

	flight := &objects.Flight{
		Id:           1,
		FlightNumber: "AFL031",
		Datetime:     "2021-10-08 20:00",
		FromAirport:  objects.Airport{Id: 2, Name: "Пулково", City: "Санкт-Петербург", Country: "Россия"},
		ToAirport:    objects.Airport{Id: 1, Name: "Шереметьево", City: "Москва", Country: "Россия"},
		Price:        1500,
	}
	db.FirstOrCreate(flight)

	return db
}

func init() {

}

func main() {
	rand.Seed(time.Now().UnixNano())

	utils.InitConfig()
	utils.InitLogger()
	defer utils.CloseLogger()

	db := initDBConnection(utils.Config.DB)
	defer db.Close()
	r := controllers.InitRouter(db)

	fmt.Printf("Server started: http://localhost:%d\n", utils.Config.Port)

	controllers.RunRouter(r, utils.Config.Port)

}
