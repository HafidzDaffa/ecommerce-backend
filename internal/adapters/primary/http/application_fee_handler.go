package http

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ApplicationFeeHandler struct {
	feeService ports.ApplicationFeeService
}

func NewApplicationFeeHandler(feeService ports.ApplicationFeeService) *ApplicationFeeHandler {
	return &ApplicationFeeHandler{
		feeService: feeService,
	}
}

// CreateApplicationFee godoc
// @Summary Create a new application fee
// @Description Create a new application fee configuration (Admin only)
// @Tags Application Fees
// @Accept json
// @Produce json
// @Param fee body domain.CreateApplicationFeeRequest true "Application Fee Request"
// @Success 201 {object} map[string]interface{} "Application fee created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Security BearerAuth
// @Router /application-fees [post]
func (h *ApplicationFeeHandler) CreateApplicationFee(c *fiber.Ctx) error {
	var req domain.CreateApplicationFeeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get user ID from context (set by auth middleware)
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	fee, err := h.feeService.CreateApplicationFee(&req, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Application fee created successfully",
		"fee":     fee,
	})
}

// GetApplicationFeeByID godoc
// @Summary Get application fee by ID
// @Description Get a single application fee by its ID
// @Tags Application Fees
// @Produce json
// @Param id path string true "Application Fee ID (UUID)"
// @Success 200 {object} map[string]interface{} "Application fee details"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Application fee not found"
// @Router /application-fees/{id} [get]
func (h *ApplicationFeeHandler) GetApplicationFeeByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid application fee ID",
		})
	}

	fee, err := h.feeService.GetApplicationFeeByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"fee": fee,
	})
}

// GetAllApplicationFees godoc
// @Summary Get all application fees
// @Description Get all application fees with optional filter by active status and pagination
// @Tags Application Fees
// @Produce json
// @Param is_active query boolean false "Filter by active status"
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{} "List of application fees"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /application-fees [get]
func (h *ApplicationFeeHandler) GetAllApplicationFees(c *fiber.Ctx) error {
	var isActive *bool
	if isActiveStr := c.Query("is_active"); isActiveStr != "" {
		active := isActiveStr == "true"
		isActive = &active
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	result, err := h.feeService.GetAllApplicationFees(isActive, page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// GetActiveByType godoc
// @Summary Get active application fee by type
// @Description Get the most recent active application fee by fee type
// @Tags Application Fees
// @Produce json
// @Param fee_type query string true "Fee Type (PERCENTAGE or FIXED)"
// @Success 200 {object} map[string]interface{} "Application fee details"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Application fee not found"
// @Router /application-fees/active [get]
func (h *ApplicationFeeHandler) GetActiveByType(c *fiber.Ctx) error {
	feeTypeStr := c.Query("fee_type")
	if feeTypeStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "fee_type query parameter is required",
		})
	}

	feeType := domain.FeeType(feeTypeStr)
	if feeType != domain.FeeTypePercentage && feeType != domain.FeeTypeFixed {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid fee type. Must be PERCENTAGE or FIXED",
		})
	}

	fee, err := h.feeService.GetActiveByType(feeType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if fee == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No active application fee found for the specified type",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"fee": fee,
	})
}

// UpdateApplicationFee godoc
// @Summary Update application fee
// @Description Update an existing application fee (Admin only)
// @Tags Application Fees
// @Accept json
// @Produce json
// @Param id path string true "Application Fee ID (UUID)"
// @Param fee body domain.UpdateApplicationFeeRequest true "Update Application Fee Request"
// @Success 200 {object} map[string]interface{} "Application fee updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Application fee not found"
// @Security BearerAuth
// @Router /application-fees/{id} [put]
func (h *ApplicationFeeHandler) UpdateApplicationFee(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid application fee ID",
		})
	}

	var req domain.UpdateApplicationFeeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	fee, err := h.feeService.UpdateApplicationFee(id, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Application fee updated successfully",
		"fee":     fee,
	})
}

// DeleteApplicationFee godoc
// @Summary Delete application fee
// @Description Delete an application fee by its ID (Admin only)
// @Tags Application Fees
// @Produce json
// @Param id path string true "Application Fee ID (UUID)"
// @Success 200 {object} map[string]interface{} "Application fee deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Application fee not found"
// @Security BearerAuth
// @Router /application-fees/{id} [delete]
func (h *ApplicationFeeHandler) DeleteApplicationFee(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid application fee ID",
		})
	}

	err = h.feeService.DeleteApplicationFee(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Application fee deleted successfully",
	})
}

// CalculateFee godoc
// @Summary Calculate application fee
// @Description Calculate the fee amount based on fee ID and base amount
// @Tags Application Fees
// @Accept json
// @Produce json
// @Param request body map[string]interface{} true "Calculate Fee Request (fee_id: UUID, base_amount: float64)"
// @Success 200 {object} map[string]interface{} "Fee calculation result"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /application-fees/calculate [post]
func (h *ApplicationFeeHandler) CalculateFee(c *fiber.Ctx) error {
	var req struct {
		FeeID      string  `json:"fee_id"`
		BaseAmount float64 `json:"base_amount"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	feeID, err := uuid.Parse(req.FeeID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid fee ID",
		})
	}

	if req.BaseAmount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Base amount must be greater than 0",
		})
	}

	feeAmount, err := h.feeService.CalculateFee(feeID, req.BaseAmount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"fee_amount":   feeAmount,
		"base_amount":  req.BaseAmount,
		"total_amount": req.BaseAmount + feeAmount,
	})
}
