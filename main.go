package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alaa-aqeel/booking-system/app/services"
	database "github.com/alaa-aqeel/booking-system/database/driver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	db := database.NewDatabase()
	err = db.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	services := services.NewServices(db)

	row, err := services.User.Row(context.Background(), services.User.Query().Columns("id").Limit(1))
	fmt.Println(err)
	var id string
	err = row.Scan(&id)
	fmt.Println(err, id)

	// err = services.Services.Create(domain.CreateServicesCommand{
	// 	Name:        optional.SetValue("test z"),
	// 	Description: optional.SetValue("test z"),
	// 	Price:       optional.SetValue(0.0),
	// 	IsActive:    optional.SetValue(true),
	// 	CreatedBy:   optional.SetValue(id),
	// })
	// fmt.Println(err)
	res, err := services.User.Find(id, services.User.LoadServicesOne)
	fmt.Println(err)
	fmt.Println(res.Services)
}
