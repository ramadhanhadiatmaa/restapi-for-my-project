package usercontroller

import (
	"apicsmfib/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var user []models.User
	models.DB.Find(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}

func Show(c *fiber.Ctx) error {

	username := c.Params("username")

	var user models.User

	if err := models.DB.Model(&user).Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "User not Found.",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error.",
		})
	}

	return c.JSON(user)
}

func Create(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := models.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User Successfully to Create.",
	})
}
func Update(c *fiber.Ctx) error {

	username := c.Params("username")

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.UpdatedAt = time.Now()

	if models.DB.Where("username = ?", username).Updates(&user).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Update Data.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data Successfully Updated.",
	})
}
func Delete(c *fiber.Ctx) error {

	username := c.Params("username")
	var user models.User

	if models.DB.Where("username = ?", username).Delete(&user).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User " + username + " not Found.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully Delete User " + username + ".",
	})
}