package http

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderHandler struct {
	orderService ports.OrderService
}

func NewOrderHandler(orderService ports.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create new order
// @Description Create a new order from selected cart items
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.CreateOrderRequest true "Create Order Request"
// @Success 201 {object} map[string]interface{} "Order created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var req domain.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	order, err := h.orderService.CreateOrder(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order created successfully",
		"order":   order,
	})
}

// GetOrder godoc
// @Summary Get order details
// @Description Retrieve detailed information about a specific order
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} domain.Order "Order retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	orderID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	order, err := h.orderService.GetOrder(userID, orderID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(order)
}

// GetUserOrders godoc
// @Summary Get user's orders
// @Description Retrieve list of user's orders with pagination
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} domain.OrderListResponse "Orders retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /orders [get]
func (h *OrderHandler) GetUserOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	orders, err := h.orderService.GetUserOrders(userID, page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(orders)
}

// GetAllOrders godoc
// @Summary Get all orders (Admin only)
// @Description Retrieve all orders with pagination - admin only
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} domain.OrderListResponse "Orders retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Router /admin/orders [get]
func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	orders, err := h.orderService.GetAllOrders(page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(orders)
}

// UpdateOrderStatus godoc
// @Summary Update order status (Admin only)
// @Description Update order status and tracking number - admin only
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Param request body domain.UpdateOrderStatusRequest true "Update Order Status Request"
// @Success 200 {object} map[string]interface{} "Order status updated"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Router /admin/orders/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	orderID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	var req domain.UpdateOrderStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.orderService.UpdateOrderStatus(orderID, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order status updated",
	})
}

// CancelOrder godoc
// @Summary Cancel order
// @Description Cancel an order (only if not yet processed)
// @Tags Orders
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} map[string]interface{} "Order cancelled"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /orders/{id}/cancel [post]
func (h *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	orderID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	if err := h.orderService.CancelOrder(userID, orderID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order cancelled",
	})
}

// GetOrderStatuses godoc
// @Summary Get order statuses
// @Description Retrieve all available order statuses
// @Tags Orders
// @Produce json
// @Success 200 {object} map[string]interface{} "Order statuses retrieved"
// @Router /orders/statuses [get]
func (h *OrderHandler) GetOrderStatuses(c *fiber.Ctx) error {
	statuses, err := h.orderService.GetOrderStatuses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"statuses": statuses,
	})
}
