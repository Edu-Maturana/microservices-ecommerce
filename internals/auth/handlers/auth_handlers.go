package handlers

import (
	"back-usm/internals/auth/core/domain"
	"back-usm/internals/auth/core/services"
	"back-usm/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandlers struct {
	authServices services.AuthServices
}

func NewAuthHandlers(authServices services.AuthServices) *AuthHandlers {
	return &AuthHandlers{
		authServices: authServices,
	}
}

func (h *AuthHandlers) GetAllAdmins(c *fiber.Ctx) error {
	users, err := h.authServices.GetAllAdmins()
	if err != nil {
		return c.Status(404).JSON("Admins not found")
	}

	return c.Status(200).JSON(users)
}

func (h *AuthHandlers) GetOneAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.authServices.GetOneAdmin(id)
	if err != nil {
		return c.Status(404).JSON("Admin not found")
	}

	return c.Status(200).JSON(user)
}

func (h *AuthHandlers) CreateAdmin(c *fiber.Ctx) error {
	var admin domain.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON("Invalid admin")
	}

	if err := utils.ValidateData(admin); err != nil {
		return c.Status(400).JSON("Invalid data")
	}

	admin, err := h.authServices.CreateAdmin(admin)
	if err != nil {
		return c.Status(400).JSON("Error creating admin")
	}

	return c.Status(201).JSON(admin)
}

func (h *AuthHandlers) UpdateAdmin(c *fiber.Ctx) error {
	var admin domain.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON("Invalid admin")
	}

	admin, err := h.authServices.UpdateAdmin(admin)
	if err != nil {
		return c.Status(400).JSON("Error updating admin")
	}

	return c.Status(200).JSON(admin)
}

func (h *AuthHandlers) DeleteAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.authServices.DeleteAdmin(id)
	if err != nil {
		return c.Status(400).JSON("Error deleting admin")
	}

	return c.SendStatus(200)
}

func (h *AuthHandlers) Login(c *fiber.Ctx) error {
	var admin domain.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON("Invalid credentials")
	}

	admin, err := h.authServices.Login(admin)
	if err != nil {
		return c.Status(400).JSON("Invalid credentials")
	}

	return c.JSON("Login successful")
}

func (h *AuthHandlers) ActivateAccount(c *fiber.Ctx) error {
	var admin domain.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON("Invalid credentials")
	}

	admin, err := h.authServices.ActivateAccount(admin)
	if err != nil {
		return c.Status(400).JSON("Invalid credentials")
	}

	return c.Status(200).JSON("Account activated")
}
