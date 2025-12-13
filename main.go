package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alaa-aqeel/booking-system/app/domain"
	"github.com/alaa-aqeel/booking-system/app/services"
	database "github.com/alaa-aqeel/booking-system/database/driver"
	"github.com/alaa-aqeel/optional-value"
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
	services.User.Create(domain.CreateUserCommand{
		Username: optional.SetValue("alaa_aqeel-2"),
		Password: optional.SetValue("123456"),
	})

	res, errs := services.User.GetAll(domain.UserQuery{
		IsActive: optional.SetValue(false),
	})
	fmt.Println(errs)
	fmt.Println(res)
}
