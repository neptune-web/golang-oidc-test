package handler

import (
	"github.com/store/database"

	"github.com/gofiber/fiber/v2"
	"github.com/store/model"
)

// Get All Testdatas from db
func GetSingleTestdatas(c *fiber.Ctx) error {
	db := database.DB.TestdataDb
	var testdatas []model.Testdata

	// find all testdatas in the database
	db.Find(&testdatas)

	// if no testdata found, return an error
	if len(testdatas) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "success", "message": "Testdatas not found", "data": []model.Testdata{}})
	}

	// return testdatas
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Testdatas Found", "data": testdatas})
}

// GetSingleTestdata from db
func GetSingleTestdata(c *fiber.Ctx) error {
	db := database.DB.TestdataDb

	// get testdata_id params
	requestID := c.Params("request_id")

	var testdata model.Testdata

	// find single testdata in the database by testdata_id
	db.Find(&testdata, "request_id = ?", requestID)

	// if there is no testdata with given testdata_id, return an error
	if testdata.RequestID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "success", "message": "Testdata not found", "data": []model.Testdata{}})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Testdata Found", "data": testdata})
}

func CreateTestdata(c *fiber.Ctx) error {
	db := database.DB.TestdataDb
	testdata := new(model.Testdata)

	var request_id int
	db.Raw("select min(ids) from generate_series(6000, 6999) as ids left join testdata on ids=testdata.request_id where testdata.request_id is null").Scan(&request_id)

	testdata.RequestID = request_id
	err := db.Create(&testdata).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create testdata", "data": err})
	}

	// Return the created testdata
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Testdata has created", "data": testdata})
}

// func UpdateUser(c *fiber.Ctx) error {
// 	type updateUser struct {
// 		Requestername string `json:"requestername"`
// 	}

// 	db := database.DB.Db

// 	var user model.Store

// 	// get id params
// 	id := c.Params("id")

// 	// find single user in the database by id
// 	db.Find(&user, "id = ?", id)

// 	// need to be fixed
// 	// if user.ID == uuid.Nil {
// 	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
// 	// }

// 	var updateUserData updateUser
// 	err := c.BodyParser(&updateUserData)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
// 	}

// 	user.Requestername = updateUserData.Requestername

// 	// Save the Changes
// 	db.Save(&user)

// 	// Return the updated user
// 	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})

// }

// func DeleteUserByID(c *fiber.Ctx) error {
// 	db := database.DB.Db
// 	var store model.Store

// 	// get id params
// 	id := c.Params("id")

// 	// find single user in the database by id
// 	db.Find(&store, "id = ?", id)
// 	store.StoreID = 0
// 	db.Save(&store)

// 	err := db.Delete(&store, "id = ?", id).Error
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
// 	}

// 	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
// }

// Login user

// func Login(c *fiber.Ctx) error {
// 	db := database.DB.Db
// 	var user model.User

// 	var input middleware.LoginInput

// 	// binding user input to a struct
// 	if err := c.BodyParser(&input); err != nil {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	// set a variable depending on the condition
// 	var query string
// 	if valid(input.Identity) {
// 		query = "email= ?"
// 	} else {
// 		query = "username= ?"
// 	}

// 	if err := db.Where(query, input.Identity).First(&user).Error; err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"status": "error", "message": "User does not exists",
// 		})

// 	}

// 	identity := input.Identity
// 	pass := input.Password

// 	if !helper.ValidatePassword(pass, user.Password) {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"status": "error", "messvalidage": "Password incorrect",
// 		})
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["identity"] = identity
// 	claims["admin"] = true
// 	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "token": t})

// }
