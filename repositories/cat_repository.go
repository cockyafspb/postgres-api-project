package repositories

import (
	"net/http"
	"postgres-api-project/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CatRepository struct {
	DB *gorm.DB
}

func (r *CatRepository) CreateCat(context *fiber.Ctx) error {
	cat := models.Cat{}
	err := context.BodyParser(&cat)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&cat).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create a cat"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "cat has been added"})
	return nil
}

func (r *CatRepository) DeleteCat(context *fiber.Ctx) error {
	catModel := models.Cat{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(catModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete a cat",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book deleted successfully",
	})
	return nil
}

func (r *CatRepository) GetCats(context *fiber.Ctx) error {
	cats := &[]models.Cat{}

	err := r.DB.Find(cats).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get a cats"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "cats fetched successfully",
		"data":    cats,
	})
	return nil
}

func (r *CatRepository) GetCatByID(context *fiber.Ctx) error {

	id := context.Params("id")
	catModel := &models.Cat{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(catModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get a cat"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "cat id fetched successfully",
		"data":    catModel,
	})
	return nil
}

func (r *CatRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_cat", r.CreateCat)
	api.Delete("delete_cat/:id", r.DeleteCat)
	api.Get("/get_cat/:id", r.GetCatByID)
	api.Get("/cats", r.GetCats)
}
