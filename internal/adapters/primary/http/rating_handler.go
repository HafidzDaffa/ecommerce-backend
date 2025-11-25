package http

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RatingHandler struct {
	ratingService ports.RatingService
}

func NewRatingHandler(ratingService ports.RatingService) *RatingHandler {
	return &RatingHandler{
		ratingService: ratingService,
	}
}

// CreateRating godoc
// @Summary Create product rating
// @Description Create a rating/review for a product from a delivered order
// @Tags Ratings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.CreateRatingRequest true "Create Rating Request"
// @Success 201 {object} map[string]interface{} "Rating created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /ratings [post]
func (h *RatingHandler) CreateRating(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var req domain.CreateRatingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	rating, err := h.ratingService.CreateRating(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Rating created successfully",
		"rating":  rating,
	})
}

// GetProductRatings godoc
// @Summary Get product ratings
// @Description Retrieve all ratings for a specific product with pagination
// @Tags Ratings
// @Produce json
// @Param product_id query string true "Product ID"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{} "Ratings retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /ratings [get]
func (h *RatingHandler) GetProductRatings(c *fiber.Ctx) error {
	productIDStr := c.Query("product_id")
	if productIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product_id is required",
		})
	}

	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	ratings, total, err := h.ratingService.GetProductRatings(productID, page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"ratings": ratings,
		"total":   total,
		"page":    page,
		"per_page": perPage,
	})
}

// GetUserRatings godoc
// @Summary Get user's ratings
// @Description Retrieve all ratings created by the authenticated user
// @Tags Ratings
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Ratings retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /ratings/my [get]
func (h *RatingHandler) GetUserRatings(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	ratings, err := h.ratingService.GetUserRatings(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"ratings": ratings,
	})
}

// GetRatingStats godoc
// @Summary Get product rating statistics
// @Description Retrieve rating statistics for a product (average, count by stars)
// @Tags Ratings
// @Produce json
// @Param product_id query string true "Product ID"
// @Success 200 {object} domain.ProductRatingStats "Rating statistics retrieved"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /ratings/stats [get]
func (h *RatingHandler) GetRatingStats(c *fiber.Ctx) error {
	productIDStr := c.Query("product_id")
	if productIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "product_id is required",
		})
	}

	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	stats, err := h.ratingService.GetRatingStats(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(stats)
}

// UpdateRating godoc
// @Summary Update rating
// @Description Update user's existing rating
// @Tags Ratings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Rating ID"
// @Param request body domain.UpdateRatingRequest true "Update Rating Request"
// @Success 200 {object} map[string]interface{} "Rating updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /ratings/{id} [put]
func (h *RatingHandler) UpdateRating(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	ratingID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid rating ID",
		})
	}

	var req domain.UpdateRatingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	rating, err := h.ratingService.UpdateRating(userID, ratingID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Rating updated successfully",
		"rating":  rating,
	})
}

// DeleteRating godoc
// @Summary Delete rating
// @Description Delete user's rating
// @Tags Ratings
// @Security BearerAuth
// @Param id path string true "Rating ID"
// @Success 200 {object} map[string]interface{} "Rating deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /ratings/{id} [delete]
func (h *RatingHandler) DeleteRating(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	ratingID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid rating ID",
		})
	}

	if err := h.ratingService.DeleteRating(userID, ratingID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Rating deleted successfully",
	})
}
