package http

import (
	"strconv"

	"github.com/bergsantana/go-contacts/internal/entity"
	"github.com/bergsantana/go-contacts/internal/usecase"
	"github.com/bergsantana/go-contacts/pkg/sanitize"
	"github.com/gofiber/fiber/v2"
)

type ContactHandler struct {
	usecase *usecase.ContactUsecase
}

func NewContactHandler(app *fiber.App, uc *usecase.ContactUsecase) {
	handler := &ContactHandler{usecase: uc}

	app.Get("/contacts", handler.GetContacts)
	app.Get("/contacts/:id", handler.GetContact)
	app.Post("/contacts", handler.CreateContact)
	app.Put("/contacts/:id", handler.UpdateContact)
	app.Delete("/contacts/:id", handler.DeleteContact)
}

func (h *ContactHandler) GetContacts(c *fiber.Ctx) error {
	contacts, err := h.usecase.GetContacts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(contacts)
}

func (h *ContactHandler) GetContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	contact, err := h.usecase.GetContactByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Contact not found"})
	}
	return c.JSON(contact)
}

func (h *ContactHandler) CreateContact(c *fiber.Ctx) error {
	var contact entity.Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Sanitiazação de campos
	contact.Name = sanitize.StrictHTML(contact.Name)
	contact.Email = sanitize.StrictHTML(contact.Email)
	contact.Phone = sanitize.StrictHTML(contact.Phone)

	if err := h.usecase.CreateContact(&contact); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(contact)
}

func (h *ContactHandler) UpdateContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var contact entity.Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	// Sanitiazação de campos
	contact.Name = sanitize.StrictHTML(contact.Name)
	contact.Email = sanitize.StrictHTML(contact.Email)
	contact.Phone = sanitize.StrictHTML(contact.Phone)

	contact.ID = uint(id)
	if err := h.usecase.UpdateContact(&contact); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(contact)
}

func (h *ContactHandler) DeleteContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.usecase.DeleteContact(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
