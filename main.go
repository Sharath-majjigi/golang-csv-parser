package main

import(
	"log"
	"github.com/gofiber/fiber/v2"
	"sharath/routes"
	"sharath/database"
)

func setupRoutes(app *fiber.App){
	app.Post("/upload-csv",routes.ImportCSV)
	app.Get("/people",routes.GetData)
	app.Get("/peopleByAge/:age",routes.GetPeopleWithAge)
	app.Get("/peopleByName/:fname",routes.GetPeopleByFName)
	app.Delete("/clean/:age",routes.DeletePeople)
}

func main(){
	database.Connect()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":8002"))
}