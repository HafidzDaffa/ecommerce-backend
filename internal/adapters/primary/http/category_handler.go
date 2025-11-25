package http

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryService ports.CategoryService
}

func NewCategoryHandler(categoryService ports.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new product category with optional image upload
// @Tags Categories
// @Accept multipart/form-data
// @Produce json
// @Param category_name formData string true "Category Name"
// @Param slug formData string false "Category Slug"
// @Param icon formData string false "Category Icon"
// @Param is_active formData boolean false "Is Active"
// @Param image formData file false "Category Image"
// @Success 201 {object} map[string]interface{} "Category created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req domain.CreateCategoryRequest

	req.CategoryName = c.FormValue("category_name")
	if slug := c.FormValue("slug"); slug != "" {
		req.Slug = &slug
	}
	if icon := c.FormValue("icon"); icon != "" {
		req.Icon = &icon
	}
	if isActive := c.FormValue("is_active"); isActive != "" {
		active := isActive == "true"
		req.IsActive = &active
	}

	imageFile, _ := c.FormFile("image")

	category, err := h.categoryService.CreateCategory(&req, imageFile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Category created successfully",
		"category": category,
	})
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Get a single category by its ID
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{} "Category details"
// @Failure 404 {object} map[string]interface{} "Category not found"
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	category, err := h.categoryService.GetCategoryByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"category": category,
	})
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get all categories with optional filter by active status
// @Tags Categories
// @Produce json
// @Param is_active query boolean false "Filter by active status"
// @Success 200 {object} map[string]interface{} "List of categories"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	var isActive *bool
	if isActiveStr := c.Query("is_active"); isActiveStr != "" {
		active := isActiveStr == "true"
		isActive = &active
	}

	categories, err := h.categoryService.GetAllCategories(isActive)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"categories": categories,
		"total":      len(categories),
	})
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update an existing category
// @Tags Categories
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Category ID"
// @Param category_name formData string false "Category Name"
// @Param slug formData string false "Category Slug"
// @Param icon formData string false "Category Icon"
// @Param is_active formData boolean false "Is Active"
// @Param image formData file false "Category Image"
// @Success 200 {object} map[string]interface{} "Category updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	var req domain.UpdateCategoryRequest

	if categoryName := c.FormValue("category_name"); categoryName != "" {
		req.CategoryName = &categoryName
	}
	if slug := c.FormValue("slug"); slug != "" {
		req.Slug = &slug
	}
	if icon := c.FormValue("icon"); icon != "" {
		req.Icon = &icon
	}
	if isActive := c.FormValue("is_active"); isActive != "" {
		active := isActive == "true"
		req.IsActive = &active
	}

	imageFile, _ := c.FormFile("image")

	category, err := h.categoryService.UpdateCategory(id, &req, imageFile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Category updated successfully",
		"category": category,
	})
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete a category by its ID
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{} "Category deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	err = h.categoryService.DeleteCategory(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}
