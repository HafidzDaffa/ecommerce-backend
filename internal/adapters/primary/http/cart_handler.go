package http

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartHandler struct {
	cartService ports.CartService
}

func NewCartHandler(cartService ports.CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

// AddToCart godoc
// @Summary Add product to cart
// @Description Add a product to user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.AddToCartRequest true "Add to Cart Request"
// @Success 201 {object} map[string]interface{} "Product added to cart"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /cart [post]
func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var req domain.AddToCartRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	cart, err := h.cartService.AddToCart(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product added to cart",
		"cart":    cart,
	})
}

// GetCart godoc
// @Summary Get user's cart
// @Description Retrieve user's shopping cart with all items and summary
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.CartSummary "Cart retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /cart [get]
func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	cart, err := h.cartService.GetCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(cart)
}

// UpdateCartItem godoc
// @Summary Update cart item
// @Description Update quantity, note, or selection status of a cart item
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Cart Item ID"
// @Param request body domain.UpdateCartRequest true "Update Cart Request"
// @Success 200 {object} map[string]interface{} "Cart item updated"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /cart/{id} [put]
func (h *CartHandler) UpdateCartItem(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	cartID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid cart ID",
		})
	}

	var req domain.UpdateCartRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	cart, err := h.cartService.UpdateCartItem(userID, cartID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cart item updated",
		"cart":    cart,
	})
}

// RemoveFromCart godoc
// @Summary Remove item from cart
// @Description Remove a specific item from user's cart
// @Tags Cart
// @Security BearerAuth
// @Param id path string true "Cart Item ID"
// @Success 200 {object} map[string]interface{} "Item removed from cart"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /cart/{id} [delete]
func (h *CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	cartID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid cart ID",
		})
	}

	if err := h.cartService.RemoveFromCart(userID, cartID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item removed from cart",
	})
}

// ClearCart godoc
// @Summary Clear cart
// @Description Remove all items from user's cart
// @Tags Cart
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Cart cleared"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /cart/clear [delete]
func (h *CartHandler) ClearCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	if err := h.cartService.ClearCart(userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cart cleared",
	})
}
