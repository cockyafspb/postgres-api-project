package repositories

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"postgres-api-project/models"
)

type OwnerRepository struct {
	DB *gorm.DB
}

func (r *OwnerRepository) CreateOwner(context *fiber.Ctx) error {
	owner := models.Owner{}
	err := context.BodyParser(&owner)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&owner).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create an owner"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "owner has been added"})
	return nil
}

func (r *OwnerRepository) DeleteOwner(context *fiber.Ctx) error {
	ownerModel := models.Owner{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(ownerModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete an owner",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book deleted successfully",
	})
	return nil
}

func (r *OwnerRepository) GetOwners(context *fiber.Ctx) error {
	owners := &[]models.Owner{}

	err := r.DB.Find(owners).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get an owners"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "owners fetched successfully",
		"data":    owners,
	})
	return nil
}

func (r *OwnerRepository) GetOwnerByID(context *fiber.Ctx) error {
	id := context.Params("id")
	ownerModel := &models.Owner{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(ownerModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get an owner"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "owner id fetched successfully",
		"data":    ownerModel,
	})
	return nil
}

func (r *OwnerRepository) GetCatsOfOwner(context *fiber.Ctx) error {
	id := context.Params("id")
	cats := &[]models.Cat{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	err := r.DB.Where("owner_id = ?", id).Find(cats).Error
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

func (r *OwnerRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_owner", r.CreateOwner)
	api.Delete("delete_owner/:id", r.DeleteOwner)
	api.Get("/get_owner/:id", r.GetOwnerByID)
	api.Get("/get_cats_of_owner/:id", r.GetCatsOfOwner)
	api.Get("/owners", r.GetOwners)
}
