package controllers

import (
	"github.com/anikets08/go-rest/database"
	"github.com/anikets08/go-rest/models"
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Healthy",
		"data":    nil,
		"code":    200,
	})
}

func AddRestaurant(c *fiber.Ctx) error {
	db := database.DB
	var listRestaurants []models.Restaurant
	restaurant := new(models.Restaurant)
	if err := c.BodyParser(restaurant); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Find(&listRestaurants)
	for _, rest := range listRestaurants {
		if rest.Email == restaurant.Email || rest.Phone == restaurant.Phone {
			return c.Status(500).SendString("Duplicate Email or Phone Number")
		}
	}
	db.Create(&restaurant)
	return c.JSON(&restaurant)
}

func AddCategory(c *fiber.Ctx) error {
	db := database.DB
	var listCategories []models.Category
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Find(&listCategories)
	for _, cat := range listCategories {
		if cat.Name == category.Name {
			return c.Status(500).SendString("Duplicate Category Name")
		}
	}
	db.Create(&category)
	return c.JSON(&category)
}

func AddMenuItem(c *fiber.Ctx) error {
	db := database.DB
	var listMenuItems []models.MenuItem
	menuItem := new(models.MenuItem)
	if err := c.BodyParser(menuItem); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Find(&listMenuItems)
	for _, item := range listMenuItems {
		if item.Name == menuItem.Name {
			return c.Status(500).SendString("Duplicate Menu Item Name")
		}
	}

	db.Create(&menuItem)
	return c.JSON(&menuItem)
}

func GetRestautant(c *fiber.Ctx) error {
	db := database.DB
	var restaurant models.Restaurant
	id := c.Params("id")
	db.First(&restaurant, id)
	if restaurant.Name == "" {
		return c.Status(404).SendString("No Restaurant found with ID")
	}
	return c.JSON(&restaurant)
}

func GetRestaurantWithCategoryWithMenu(c *fiber.Ctx) error {
	db := database.DB
	var restaurant models.Restaurant
	id := c.Params("id")
	db.Preload("Categories").Preload("Categories.MenuItems").First(&restaurant, id)
	if restaurant.Name == "" {
		return c.Status(404).SendString("No Restaurant found with ID")
	}
	return c.JSON(&restaurant)
}
