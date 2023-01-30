package handler

import (
	"github.com/store/database"

	"github.com/gofiber/fiber/v2"
	"github.com/store/model"
)

// Create a Store
func CreateStore(c *fiber.Ctx) error {
	db := database.DB.StoreDb
	store := new(model.Store)

	var store_id int
	db.Raw("select min(ids) from generate_series(6000, 6999) as ids left join stores on ids=stores.store_id where stores.store_id is null").Scan(&store_id)

	store.StoreID = store_id
	err := db.Create(&store).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create store", "data": err})
	}

	// Return the created store
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Store has created", "data": store})
}

// Get All Stores from db
func GetAllStores(c *fiber.Ctx) error {
	db := database.DB.StoreDb
	var stores []model.Store

	// find all stores in the database
	db.Find(&stores)

	// if no store found, return an error
	if len(stores) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stores not found", "data": nil})
	}

	// return stores
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Stores Found", "data": stores})
}

// GetSingleStore from db
func GetSingleStore(c *fiber.Ctx) error {
	db := database.DB.StoreDb

	// get store_id params
	storeID := c.Params("store_id")

	var store model.Store

	// find single store in the database by store_id
	db.Find(&store, "store_id = ?", storeID)

	// if there is no store with given store_id, return an error
	if store.StoreID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Store not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Store Found", "data": store})
}
