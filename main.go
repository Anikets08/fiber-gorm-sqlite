package main

import (
	"github.com/anikets08/go-rest/controllers"
	"github.com/anikets08/go-rest/database"
	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Get("/health", controllers.Health)
	app.Post("/restaurant", controllers.AddRestaurant)
	app.Get("/restaurant/:id", controllers.GetRestautant)
	app.Post("/category", controllers.AddCategory)
	app.Get("/res-category/:id", controllers.GetRestaurantWithCategoryWithMenu)
	app.Post("/menu", controllers.AddMenuItem)
}

func main() {
	app := fiber.New()
	database.ConnectDB()
	Routers(app)
	app.Listen(":3000")
}
