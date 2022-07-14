package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"shop/Config"
	"shop/Models"
	"shop/Repo"
	"shop/Routes"
	"shop/Service"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer func() {
		err := Config.DB.Close()
		if err != nil {
			fmt.Println("Closing DB failed with error: ", err.Error())
		}
	}()

	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Order{})

	Service.NewProdService(Repo.NewProductRepo())
	Service.NewOrderService(Repo.NewOrderRepo())
	r := Routes.SetupRouter()
	// running
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
