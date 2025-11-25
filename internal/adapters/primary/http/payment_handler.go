package http

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	paymentService ports.PaymentService
}

func NewPaymentHandler(paymentService ports.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

// CreatePayment godoc
// @Summary Create payment for an order
// @Description Create a payment transaction using Xendit payment gateway
// @Tags Payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.CreatePaymentRequest true "Create Payment Request"
// @Success 201 {object} domain.PaymentResponse "Payment created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /payments [post]
func (h *PaymentHandler) CreatePayment(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var req domain.CreatePaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	payment, err := h.paymentService.CreatePayment(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Payment created successfully",
		"payment": payment,
	})
}

// GetPaymentByID godoc
// @Summary Get payment details
// @Description Get detailed information about a specific payment
// @Tags Payments
// @Produce json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} domain.PaymentTransaction "Payment retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /payments/{id} [get]
func (h *PaymentHandler) GetPaymentByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	paymentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payment ID",
		})
	}

	payment, err := h.paymentService.GetPaymentByID(userID, paymentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"payment": payment,
	})
}

// GetPaymentByOrderID godoc
// @Summary Get payment by order ID
// @Description Get payment information for a specific order
// @Tags Payments
// @Produce json
// @Security BearerAuth
// @Param order_id query string true "Order ID"
// @Success 200 {object} domain.PaymentTransaction "Payment retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /payments/order [get]
func (h *PaymentHandler) GetPaymentByOrderID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	orderIDStr := c.Query("order_id")
	if orderIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "order_id is required",
		})
	}

	orderID, err := uuid.Parse(orderIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	payment, err := h.paymentService.GetPaymentByOrderID(userID, orderID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"payment": payment,
	})
}

// GetUserPayments godoc
// @Summary Get user's payments
// @Description Retrieve list of user's payments with pagination
// @Tags Payments
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} domain.PaymentListResponse "Payments retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /payments [get]
func (h *PaymentHandler) GetUserPayments(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	payments, err := h.paymentService.GetUserPayments(userID, page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(payments)
}

// GetAllPayments godoc
// @Summary Get all payments (Admin only)
// @Description Retrieve all payments with pagination - admin only
// @Tags Payments
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} domain.PaymentListResponse "Payments retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /admin/payments [get]
func (h *PaymentHandler) GetAllPayments(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	payments, err := h.paymentService.GetAllPayments(page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(payments)
}

// XenditCallback godoc
// @Summary Xendit payment callback
// @Description Webhook endpoint for Xendit payment notifications
// @Tags Payments
// @Accept json
// @Produce json
// @Param payload body domain.XenditCallbackPayload true "Xendit Callback Payload"
// @Success 200 {object} map[string]interface{} "Callback processed successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /payments/xendit/callback [post]
func (h *PaymentHandler) XenditCallback(c *fiber.Ctx) error {
	var payload domain.XenditCallbackPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// TODO: Verify callback authenticity using Xendit callback token

	err := h.paymentService.HandleXenditCallback(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Callback processed successfully",
	})
}

// CheckPaymentStatus godoc
// @Summary Check payment status
// @Description Check the current status of a payment from Xendit
// @Tags Payments
// @Produce json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} domain.PaymentTransaction "Payment status retrieved"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /payments/{id}/status [get]
func (h *PaymentHandler) CheckPaymentStatus(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	paymentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payment ID",
		})
	}

	payment, err := h.paymentService.CheckPaymentStatus(userID, paymentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"payment": payment,
	})
}

// CancelPayment godoc
// @Summary Cancel payment
// @Description Cancel a pending payment transaction
// @Tags Payments
// @Produce json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} map[string]interface{} "Payment cancelled successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /payments/{id}/cancel [post]
func (h *PaymentHandler) CancelPayment(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	paymentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payment ID",
		})
	}

	err = h.paymentService.CancelPayment(userID, paymentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Payment cancelled successfully",
	})
}
