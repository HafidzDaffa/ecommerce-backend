package http

import (
	"ecommerce-backend/internal/core/domain"
	"ecommerce-backend/internal/core/ports"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productService ports.ProductService
}

func NewProductHandler(productService ports.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with categories
// @Tags Products
// @Accept json
// @Produce json
// @Param request body domain.CreateProductRequest true "Product Request"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{} "Product created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	log.Println("📦 CreateProduct called")
	
	userID := c.Locals("user_id")
	log.Printf("   userID from context: %v (type: %T)", userID, userID)
	
	if userID == nil {
		log.Println("   ❌ userID is nil")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var req domain.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("   ❌ BodyParser error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	log.Printf("   ✅ Request parsed: %+v", req)

	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		log.Printf("   ❌ Type assertion failed. userID type: %T, value: %v", userID, userID)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": fmt.Sprintf("Invalid user ID type: expected uuid.UUID but got %T", userID),
		})
	}
	log.Printf("   ✅ User UUID: %s", userUUID)

	product, err := h.productService.CreateProduct(userUUID, &req)
	if err != nil {
		log.Printf("   ❌ CreateProduct service error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	log.Printf("   ✅ Product created: %s", product.ID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": product,
	})
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get a single product by its ID with categories and galleries
// @Tags Products
// @Produce json
// @Param id path string true "Product ID (UUID)"
// @Success 200 {object} map[string]interface{} "Product details"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"product": product,
	})
}

// GetProductBySlug godoc
// @Summary Get product by slug
// @Description Get a single product by its slug with categories and galleries
// @Tags Products
// @Produce json
// @Param slug path string true "Product Slug"
// @Success 200 {object} map[string]interface{} "Product details"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /products/slug/{slug} [get]
func (h *ProductHandler) GetProductBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	product, err := h.productService.GetProductBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"product": product,
	})
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products with pagination
// @Tags Products
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param is_published query boolean false "Filter by published status"
// @Success 200 {object} map[string]interface{} "List of products"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	var isPublished *bool
	if isPublishedStr := c.Query("is_published"); isPublishedStr != "" {
		published := isPublishedStr == "true"
		isPublished = &published
	}

	products, total, err := h.productService.GetAllProducts(page, limit, isPublished)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	totalPages := (total + limit - 1) / limit

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
		"pagination": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

// GetProductsByCategoryID godoc
// @Summary Get products by category ID
// @Description Get all published products in a specific category with pagination
// @Tags Products
// @Produce json
// @Param category_id path int true "Category ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{} "List of products"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /products/category/{category_id} [get]
func (h *ProductHandler) GetProductsByCategoryID(c *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("category_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	products, total, err := h.productService.GetProductsByCategoryID(categoryID, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	totalPages := (total + limit - 1) / limit

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products": products,
		"pagination": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update an existing product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID (UUID)"
// @Param request body domain.UpdateProductRequest true "Product Update Request"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Product updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	var req domain.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	product, err := h.productService.UpdateProduct(id, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Soft delete a product by its ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID (UUID)"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Product deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	err = h.productService.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

// AddProductGallery godoc
// @Summary Add product gallery image
// @Description Upload and add an image to product gallery
// @Tags Product Galleries
// @Accept multipart/form-data
// @Produce json
// @Param product_id formData string true "Product ID (UUID)"
// @Param display_order formData int false "Display Order" default(0)
// @Param is_thumbnail formData boolean false "Is Thumbnail" default(false)
// @Param image formData file true "Product Image"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{} "Gallery image added successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /products/galleries [post]
func (h *ProductHandler) AddProductGallery(c *fiber.Ctx) error {
	var req domain.CreateProductGalleryRequest

	req.ProductID = c.FormValue("product_id")
	displayOrder, _ := strconv.Atoi(c.FormValue("display_order"))
	req.DisplayOrder = displayOrder

	if isThumbnail := c.FormValue("is_thumbnail"); isThumbnail != "" {
		thumbnail := isThumbnail == "true"
		req.IsThumbnail = &thumbnail
	}

	imageFile, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Image file is required",
		})
	}

	gallery, err := h.productService.AddProductGallery(&req, imageFile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Gallery image added successfully",
		"gallery": gallery,
	})
}

// GetProductGalleries godoc
// @Summary Get product galleries
// @Description Get all gallery images for a specific product
// @Tags Product Galleries
// @Produce json
// @Param product_id path string true "Product ID (UUID)"
// @Success 200 {object} map[string]interface{} "List of gallery images"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /products/{product_id}/galleries [get]
func (h *ProductHandler) GetProductGalleries(c *fiber.Ctx) error {
	productID, err := uuid.Parse(c.Params("product_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	galleries, err := h.productService.GetProductGalleries(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"galleries": galleries,
		"total":     len(galleries),
	})
}

// UpdateProductGallery godoc
// @Summary Update product gallery
// @Description Update gallery image display order or thumbnail status
// @Tags Product Galleries
// @Accept json
// @Produce json
// @Param id path string true "Gallery ID (UUID)"
// @Param request body domain.UpdateProductGalleryRequest true "Gallery Update Request"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Gallery updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /products/galleries/{id} [put]
func (h *ProductHandler) UpdateProductGallery(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid gallery ID",
		})
	}

	var req domain.UpdateProductGalleryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	gallery, err := h.productService.UpdateProductGallery(id, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gallery updated successfully",
		"gallery": gallery,
	})
}

// DeleteProductGallery godoc
// @Summary Delete product gallery
// @Description Delete a gallery image by its ID
// @Tags Product Galleries
// @Produce json
// @Param id path string true "Gallery ID (UUID)"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Gallery deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /products/galleries/{id} [delete]
func (h *ProductHandler) DeleteProductGallery(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid gallery ID",
		})
	}

	err = h.productService.DeleteProductGallery(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gallery deleted successfully",
	})
}
